package sqlgen

import (
	"fmt"
	"github.com/albatiqy/gopoh-app-dbimport/pkg/gendriver"
	"github.com/albatiqy/gopoh/pkg/lib/decimal"
	"github.com/albatiqy/gopoh/pkg/lib/null"
	"time"
)

type RegistrasiDapodikMstWilayahModel struct {
	KodeWilayah    string          `json:"kode_wilayah"`
	Nama           string          `json:"nama"`
	IdLevelWilayah int32           `json:"id_level_wilayah"`
	MstKodeWilayah null.String     `json:"mst_kode_wilayah"`
	NegaraId       string          `json:"negara_id"`
	AsalWilayah    null.String     `json:"asal_wilayah"`
	KodeBps        null.String     `json:"kode_bps"`
	KodeDagri      null.String     `json:"kode_dagri"`
	KodeKeu        null.String     `json:"kode_keu"`
	IdProv         null.String     `json:"id_prov"`
	IdKabkota      null.String     `json:"id_kabkota"`
	IdKec          null.String     `json:"id_kec"`
	ADesa          decimal.Decimal `json:"a_desa"`
	AKelurahan     decimal.Decimal `json:"a_kelurahan"`
	A35            decimal.Decimal `json:"a_35"`
	AUrban         decimal.Decimal `json:"a_urban"`
	KategoriDesaId null.Int32      `json:"kategori_desa_id"`
	ExpiredDate    null.Time       `json:"expired_date"`
	SyncedAt       time.Time       `json:"synced_at"`
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
	record.KategoriDesaId,
	record.ExpiredDate,
	record.SyncedAt,
*/

// 	record.ExpiredDate = null.NewTime(record.ExpiredDate.Time.Local(), record.ExpiredDate.Valid) // convert to local
// 	record.SyncedAt = record.SyncedAt.Local() // convert to local

type RegistrasiDapodikMstWilayah struct {
	cols               []string
	tableName          string
	genDriver          gendriver.Engine
	insertPlaceHolders string
}

func (gen RegistrasiDapodikMstWilayah) Values(kodeWilayah string, nama string, idLevelWilayah int32, mstKodeWilayah null.String, negaraId string, asalWilayah null.String, kodeBps null.String, kodeDagri null.String, kodeKeu null.String, idProv null.String, idKabkota null.String, idKec null.String, aDesa decimal.Decimal, aKelurahan decimal.Decimal, a35 decimal.Decimal, aUrban decimal.Decimal, kategoriDesaId null.Int32, expiredDate null.Time, syncedAt time.Time) string {
	return fmt.Sprintf("("+gen.insertPlaceHolders+")", gen.genDriver.QuoteString(kodeWilayah), gen.genDriver.QuoteString(nama), idLevelWilayah, gen.genDriver.Quote(mstKodeWilayah), gen.genDriver.QuoteString(negaraId), gen.genDriver.Quote(asalWilayah), gen.genDriver.Quote(kodeBps), gen.genDriver.Quote(kodeDagri), gen.genDriver.Quote(kodeKeu), gen.genDriver.Quote(idProv), gen.genDriver.Quote(idKabkota), gen.genDriver.Quote(idKec), gen.genDriver.Quote(aDesa), gen.genDriver.Quote(aKelurahan), gen.genDriver.Quote(a35), gen.genDriver.Quote(aUrban), gen.genDriver.Quote(kategoriDesaId), gen.genDriver.Quote(expiredDate), gen.genDriver.Quote(syncedAt))
}

func NewRegistrasiDapodikMstWilayah(genDriver gendriver.Engine) *RegistrasiDapodikMstWilayah {
	insertPlaceHolders := genDriver.InsertPlaceholders([]interface{}{
		(*string)(nil),
		(*string)(nil),
		(*int32)(nil),
		(*null.String)(nil),
		(*string)(nil),
		(*null.String)(nil),
		(*null.String)(nil),
		(*null.String)(nil),
		(*null.String)(nil),
		(*null.String)(nil),
		(*null.String)(nil),
		(*null.String)(nil),
		(*decimal.Decimal)(nil),
		(*decimal.Decimal)(nil),
		(*decimal.Decimal)(nil),
		(*decimal.Decimal)(nil),
		(*null.Int32)(nil),
		(*null.Time)(nil),
		(*time.Time)(nil),
	})
	return &RegistrasiDapodikMstWilayah{
		cols: []string{
			"kode_wilayah",
			"nama",
			"id_level_wilayah",
			"mst_kode_wilayah",
			"negara_id",
			"asal_wilayah",
			"kode_bps",
			"kode_dagri",
			"kode_keu",
			"id_prov",
			"id_kabkota",
			"id_kec",
			"a_desa",
			"a_kelurahan",
			"a_35",
			"a_urban",
			"kategori_desa_id",
			"expired_date",
			"synced_at",
		},
		tableName:          "dapodik_mst_wilayah",
		genDriver:          genDriver,
		insertPlaceHolders: insertPlaceHolders,
	}
}
