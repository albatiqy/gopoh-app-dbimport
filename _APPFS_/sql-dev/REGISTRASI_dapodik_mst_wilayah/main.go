package main

import (
	"os"

	"github.com/albatiqy/gopoh/contract/log"
	// "github.com/albatiqy/gopoh/contract/sqldb"

	_ "github.com/albatiqy/gopoh-app-dbimport/pkg/gendriver/postgres"
)

var (
	sqlGen = genDev{
		DBEnvKey: "REGISTRASI",
		DBDriver: "postgres",
		TableName: "dapodik_mst_wilayah",
		FieldDefs: fieldDefs,
		KeyAttr: "kode_wilayah",
		KeyAuto: false,
		KeyCanUpdate: false,
		SoftDelete: false,
		ObjectName: "REGISTRASI_dapodik_mst_wilayah", // lower case singular
		ObjectPackage: "sqlgen",
		OverridesLabel: map[string]string{
				// "kode_wilayah":  "Kode Wilayah",
				// "nama":  "Nama",
				// "id_level_wilayah":  "Id Level Wilayah",
				// "mst_kode_wilayah":  "Mst Kode Wilayah",
				// "negara_id":  "Negara Id",
				// "asal_wilayah":  "Asal Wilayah",
				// "kode_bps":  "Kode Bps",
				// "kode_dagri":  "Kode Dagri",
				// "kode_keu":  "Kode Keu",
				// "id_prov":  "Id Prov",
				// "id_kabkota":  "Id Kabkota",
				// "id_kec":  "Id Kec",
				// "a_desa":  "A Desa",
				// "a_kelurahan":  "A Kelurahan",
				// "a_35":  "A 35",
				// "a_urban":  "A Urban",
				// "kategori_desa_id":  "Kategori Desa Id",
				// "expired_date":  "Expired Date",
				// "synced_at":  "Synced At",
			},
		OverridesStructField: map[string]string{
				// "kode_wilayah":  "KodeWilayah",
				// "nama":  "Nama",
				// "id_level_wilayah":  "IdLevelWilayah",
				// "mst_kode_wilayah":  "MstKodeWilayah",
				// "negara_id":  "NegaraId",
				// "asal_wilayah":  "AsalWilayah",
				// "kode_bps":  "KodeBps",
				// "kode_dagri":  "KodeDagri",
				// "kode_keu":  "KodeKeu",
				// "id_prov":  "IdProv",
				// "id_kabkota":  "IdKabkota",
				// "id_kec":  "IdKec",
				// "a_desa":  "ADesa",
				// "a_kelurahan":  "AKelurahan",
				// "a_35":  "A35",
				// "a_urban":  "AUrban",
				// "kategori_desa_id":  "KategoriDesaId",
				// "expired_date":  "ExpiredDate",
				// "synced_at":  "SyncedAt",
			},
		OverridesType: map[string]interface{}{
				// "kode_wilayah":  (*string)(nil),
				// "nama":  (*string)(nil),
				// "id_level_wilayah":  (*int32)(nil),
				// "mst_kode_wilayah":  (*null.String)(nil),
				// "negara_id":  (*string)(nil),
				// "asal_wilayah":  (*null.String)(nil),
				// "kode_bps":  (*null.String)(nil),
				// "kode_dagri":  (*null.String)(nil),
				// "kode_keu":  (*null.String)(nil),
				// "id_prov":  (*null.String)(nil),
				// "id_kabkota":  (*null.String)(nil),
				// "id_kec":  (*null.String)(nil),
				// "a_desa":  (*decimal.Decimal)(nil),
				// "a_kelurahan":  (*decimal.Decimal)(nil),
				// "a_35":  (*decimal.Decimal)(nil),
				// "a_urban":  (*decimal.Decimal)(nil),
				// "kategori_desa_id":  (*null.Int32)(nil),
				// "expired_date":  (*null.Time)(nil),
				// "synced_at":  (*time.Time)(nil),
			},
		OverridesJSON: map[string]string{
				// "kode_wilayah":  "kode_wilayah",
				// "nama":  "nama",
				// "id_level_wilayah":  "id_level_wilayah",
				// "mst_kode_wilayah":  "mst_kode_wilayah",
				// "negara_id":  "negara_id",
				// "asal_wilayah":  "asal_wilayah",
				// "kode_bps":  "kode_bps",
				// "kode_dagri":  "kode_dagri",
				// "kode_keu":  "kode_keu",
				// "id_prov":  "id_prov",
				// "id_kabkota":  "id_kabkota",
				// "id_kec":  "id_kec",
				// "a_desa":  "a_desa",
				// "a_kelurahan":  "a_kelurahan",
				// "a_35":  "a_35",
				// "a_urban":  "a_urban",
				// "kategori_desa_id":  "kategori_desa_id",
				// "expired_date":  "expired_date",
				// "synced_at":  "synced_at",
			},
	}
)

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	sqlGen.Generate(workingDir)
}