package sqlgen

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/albatiqy/gopoh/pkg/lib/null"
	"github.com/albatiqy/gopoh/pkg/lib/decimal"
	"time"
	// "github.com/albatiqy/gopoh-app-dbimport/pkg/gendriver"
)

type RegistrasiDapodikMstWilayahModel struct {
	KodeWilayah string `json:"kode_wilayah"`
	Nama string `json:"nama"`
	IdLevelWilayah int32 `json:"id_level_wilayah"`
	MstKodeWilayah null.String `json:"mst_kode_wilayah"`
	NegaraId string `json:"negara_id"`
	AsalWilayah null.String `json:"asal_wilayah"`
	KodeBps null.String `json:"kode_bps"`
	KodeDagri null.String `json:"kode_dagri"`
	KodeKeu null.String `json:"kode_keu"`
	IdProv null.String `json:"id_prov"`
	IdKabkota null.String `json:"id_kabkota"`
	IdKec null.String `json:"id_kec"`
	ADesa decimal.Decimal `json:"a_desa"`
	AKelurahan decimal.Decimal `json:"a_kelurahan"`
	A35 decimal.Decimal `json:"a_35"`
	AUrban decimal.Decimal `json:"a_urban"`
	KategoriDesaId null.Int32 `json:"kategori_desa_id"`
	ExpiredDate null.Time `json:"expired_date"`
	SyncedAt time.Time `json:"synced_at"`
}

/*
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
			&record.ExpiredDate, // warning from UTC result
			&record.SyncedAt, // warning from UTC result
*/

/*
INSERT INTO dapodik_mst_wilayah (kode_wilayah,nama,id_level_wilayah,mst_kode_wilayah,negara_id,asal_wilayah,kode_bps,kode_dagri,kode_keu,id_prov,id_kabkota,id_kec,a_desa,a_kelurahan,a_35,a_urban,kategori_desa_id,expired_date,synced_at) VALUES %s

SELECT kode_wilayah,nama,id_level_wilayah,mst_kode_wilayah,negara_id,asal_wilayah,kode_bps,kode_dagri,kode_keu,id_prov,id_kabkota,id_kec,a_desa,a_kelurahan,a_35,a_urban,kategori_desa_id,expired_date,synced_at FROM dapodik_mst_wilayah
*/


// 	record.ExpiredDate = null.NewTime(record.ExpiredDate.Time.Local(), record.ExpiredDate.Valid) // convert to local
// 	record.SyncedAt = record.SyncedAt.Local() // convert to local


type RegistrasiDapodikMstWilayah struct {
}

func (gen RegistrasiDapodikMstWilayah) Values(kodeWilayah string, nama string, idLevelWilayah int32, mstKodeWilayah null.String, negaraId string, asalWilayah null.String, kodeBps null.String, kodeDagri null.String, kodeKeu null.String, idProv null.String, idKabkota null.String, idKec null.String, aDesa decimal.Decimal, aKelurahan decimal.Decimal, a35 decimal.Decimal, aUrban decimal.Decimal, kategoriDesaId null.Int32, expiredDate null.Time, syncedAt time.Time) string {

	quoteString := func(str string) string {
		return strings.ReplaceAll(str, "'", "''")
	}
	quote := func(val interface{}) string {
		switch val := val.(type) {
		case time.Time:
			return "'" + val.Format("2006-01-02 15:04:05") + "'"
		case null.String:
			if val.Valid {
				return "'" + quoteString(val.String) + "'"
			}
			return "NULL"
		case null.Time:
			if val.Valid {
				return "'" + val.Time.Format("2006-01-02 15:04:05") + "'"
			}
			return "NULL"
		case null.Bool:
			if val.Valid {
				if val.Bool {
					return "TRUE"
				}
				return "FALSE"
			}
			return "NULL"
		case null.Int32:
			if val.Valid {
				return strconv.FormatInt(int64(val.Int32), 10)
			}
			return "NULL"
		case null.Int64:
			if val.Valid {
				return strconv.FormatInt(val.Int64, 10)
			}
			return "NULL"
		case null.Float64:
			if val.Valid {
				return strconv.FormatFloat(val.Float64, 'f', 10, 64) // precission?
			}
			return "NULL"
		case decimal.Decimal:
			return val.String()
		case decimal.NullDecimal:
			if val.Valid {
				return val.Decimal.String()
			}
			return "NULL"
		}
		return "NULL"
	}
	
	return fmt.Sprintf(`('%s','%s',%d,%s,'%s',%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)`, quoteString(kodeWilayah), quoteString(nama), idLevelWilayah, quote(mstKodeWilayah), quoteString(negaraId), quote(asalWilayah), quote(kodeBps), quote(kodeDagri), quote(kodeKeu), quote(idProv), quote(idKabkota), quote(idKec), quote(aDesa), quote(aKelurahan), quote(a35), quote(aUrban), quote(kategoriDesaId), quote(expiredDate), quote(syncedAt))
}

func NewRegistrasiDapodikMstWilayah() *RegistrasiDapodikMstWilayah {
	return &RegistrasiDapodikMstWilayah{}
}