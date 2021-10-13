package dapodiksync

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mitchellh/cli"

	"github.com/albatiqy/gopoh-app-dbimport/internal/sqlgen"
	"github.com/albatiqy/gopoh/contract/repository/sqldb"

	"github.com/albatiqy/gopoh/pkg/lib/decimal"
	"github.com/albatiqy/gopoh/pkg/lib/null"
)

type QRefMstWilayah struct {
	KodeWilayah    string              `json:"kode_wilayah"`
	Nama           string              `json:"nama"`
	IdLevelWilayah int32               `json:"id_level_wilayah"`
	MstKodeWilayah null.String         `json:"mst_kode_wilayah"`
	NegaraId       string              `json:"negara_id"`
	AsalWilayah    null.String         `json:"asal_wilayah"`
	KodeBps        null.String         `json:"kode_bps"`
	KodeDagri      null.String         `json:"kode_dagri"`
	KodeKeu        null.String         `json:"kode_keu"`
	IdProv         null.String         `json:"id_prov"`
	IdKabkota      null.String         `json:"id_kabkota"`
	IdKec          null.String         `json:"id_kec"`
	ADesa          decimal.Decimal     `json:"a_desa"`
	AKelurahan     decimal.Decimal     `json:"a_kelurahan"`
	A35            decimal.Decimal     `json:"a_35"`
	AUrban         decimal.Decimal     `json:"a_urban"`
	KategoriDesaId decimal.NullDecimal `json:"kategori_desa_id"`
	CreateDate     time.Time           `json:"create_date"`
	LastUpdate     time.Time           `json:"last_update"`
	ExpiredDate    null.Time           `json:"expired_date"`
	LastSync       time.Time           `json:"last_sync"`
}

type MstWilayah struct {
	sqlGen           *sqlgen.RegistrasiDapodikMstWilayah
	srcSqlserverConn *sqldb.Conn
	dstPostgresConn  *sqldb.Conn
}

func NewMstWilayah(srcSqlserverConn, dstPostgresConn *sqldb.Conn) *MstWilayah {
	return &MstWilayah{
		sqlGen:           sqlgen.NewRegistrasiDapodikMstWilayah(),
		srcSqlserverConn: srcSqlserverConn,
		dstPostgresConn:  dstPostgresConn,
	}
}

func (dbo MstWilayah) CreateTable() error {
	return nil
}

func (dbo MstWilayah) Dump(ui cli.Ui) {
	strInsert := "INSERT INTO dapodik_mst_wilayah (kode_wilayah,nama,id_level_wilayah,mst_kode_wilayah,negara_id,asal_wilayah,kode_bps,kode_dagri,kode_keu,id_prov,id_kabkota,id_kec,a_desa,a_kelurahan,a_35,a_urban,kategori_desa_id,expired_date,synced_at) VALUES %s"

	rows, err := dbo.srcSqlserverConn.DB.Query(
		"SELECT a.kode_wilayah,a.nama,a.id_level_wilayah,a.mst_kode_wilayah,a.negara_id,a.asal_wilayah,a.kode_bps,a.kode_dagri,a.kode_keu,a.id_prov,a.id_kabkota,a.id_kec,a.a_desa,a.a_kelurahan,a.a_35,a.a_urban,a.kategori_desa_id,a.create_date,a.last_update,a.expired_date,a.last_sync FROM ref.mst_wilayah a",
	)
	if err != nil {
		ui.Error(err.Error())
		return
	}
	defer rows.Close()

	var (
		bbResult strings.Builder
		copyCnt  int
		i        = int(1)
	)

	ui.Output("\033[2J")
	// fmt.Printf("\033[%d;%dH", line, col) gotoxy

	for rows.Next() {
		record := QRefMstWilayah{}
		if err := rows.Scan(
			&record.KodeWilayah,
			&record.Nama,
			&record.IdLevelWilayah,
			&record.MstKodeWilayah,
			&record.NegaraId,
			&record.AsalWilayah,
			&record.KodeBps,
			&record.KodeDagri,
			&record.KodeKeu,
			&record.IdProv,
			&record.IdKabkota,
			&record.IdKec,
			&record.ADesa,
			&record.AKelurahan,
			&record.A35,
			&record.AUrban,
			&record.KategoriDesaId,
			&record.CreateDate,  // warning from UTC result
			&record.LastUpdate,  // warning from UTC result
			&record.ExpiredDate, // warning from UTC result
			&record.LastSync,    // warning from UTC result
		); err != nil {
			ui.Error(err.Error())
			return
		}

		record.LastUpdate = record.LastUpdate.Local()                                                // convert to local
		record.CreateDate = record.CreateDate.Local()                                                // convert to local
		record.ExpiredDate = null.NewTime(record.ExpiredDate.Time.Local(), record.ExpiredDate.Valid) // convert to local
		record.LastSync = record.LastSync.Local()                                                    // convert to local

		kategoriDesaId := null.NewInt32(int32(record.KategoriDesaId.Decimal.IntPart()), record.KategoriDesaId.Valid)
		
		bbResult.WriteString(
			dbo.sqlGen.Values(
				record.KodeWilayah,
				record.Nama,
				record.IdLevelWilayah,
				record.MstKodeWilayah,
				record.NegaraId,
				record.AsalWilayah,
				record.KodeBps,
				record.KodeDagri,
				record.KodeKeu,
				record.IdProv,
				record.IdKabkota,
				record.IdKec,
				record.ADesa,
				record.AKelurahan,
				record.A35,
				record.AUrban,
				kategoriDesaId,
				record.ExpiredDate,
				time.Now(),
			) + ",\n",
		)
		if i%50 == 0 {
			insertValues := strings.TrimRight(bbResult.String(), ",\n")
			//ui.Output(insertValues)
			_, err := dbo.dstPostgresConn.DB.Exec(fmt.Sprintf(strInsert, insertValues))
			if err != nil {
				ui.Error(err.Error())
				return
			}
			/*
				else {
					ui.Output("OK")
				}
			*/
			copyCnt += 50
			ui.Output("\033[u\033[K")
			ui.Output(strconv.Itoa(copyCnt) + " records tercopy")
			//ui.Output("---")
			bbResult.Reset()
		}
		i++
	}
	if bbResult.Len() > 0 {
		insertValues := strings.TrimRight(bbResult.String(), ",\n")
		//ui.Output(insertValues)
		_, err := dbo.dstPostgresConn.DB.Exec(fmt.Sprintf(strInsert, insertValues))
		if err != nil {
			ui.Error(err.Error())
			return
		} else {
			ui.Output("\033[u\033[K")
			ui.Output("SELESAI, " + strconv.Itoa(i-1) + " records tercopy")
		}
		bbResult.Reset()
	} else {
		ui.Output("\033[u\033[K")
		ui.Output("SELESAI, " + strconv.Itoa(i-1) + " records tercopy")
	}
}
