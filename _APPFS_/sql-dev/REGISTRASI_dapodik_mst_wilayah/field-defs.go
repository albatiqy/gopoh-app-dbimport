// hasil generate, jangan diedit
package main

import (
	"github.com/albatiqy/gopoh/contract/gen/driver"
	"github.com/albatiqy/gopoh/pkg/lib/decimal"
	"github.com/albatiqy/gopoh/pkg/lib/null"
	"time"
)

var (
	fieldDefs = map[string]driver.FieldDef{
		"kode_wilayah": {Col: "kode_wilayah", Type: (*string)(nil), JSON: "kode_wilayah", Label: "Kode Wilayah", Ordinal: 0, DBRequired: true},
		"nama": {Col: "nama", Type: (*string)(nil), JSON: "nama", Label: "Nama", Ordinal: 1, DBRequired: true},
		"id_level_wilayah": {Col: "id_level_wilayah", Type: (*int32)(nil), JSON: "id_level_wilayah", Label: "Id Level Wilayah", Ordinal: 2, DBRequired: true},
		"mst_kode_wilayah": {Col: "mst_kode_wilayah", Type: (*null.String)(nil), JSON: "mst_kode_wilayah", Label: "Mst Kode Wilayah", Ordinal: 3},
		"negara_id": {Col: "negara_id", Type: (*string)(nil), JSON: "negara_id", Label: "Negara Id", Ordinal: 4, DBRequired: true},
		"asal_wilayah": {Col: "asal_wilayah", Type: (*null.String)(nil), JSON: "asal_wilayah", Label: "Asal Wilayah", Ordinal: 5},
		"kode_bps": {Col: "kode_bps", Type: (*null.String)(nil), JSON: "kode_bps", Label: "Kode Bps", Ordinal: 6},
		"kode_dagri": {Col: "kode_dagri", Type: (*null.String)(nil), JSON: "kode_dagri", Label: "Kode Dagri", Ordinal: 7},
		"kode_keu": {Col: "kode_keu", Type: (*null.String)(nil), JSON: "kode_keu", Label: "Kode Keu", Ordinal: 8},
		"id_prov": {Col: "id_prov", Type: (*null.String)(nil), JSON: "id_prov", Label: "Id Prov", Ordinal: 9},
		"id_kabkota": {Col: "id_kabkota", Type: (*null.String)(nil), JSON: "id_kabkota", Label: "Id Kabkota", Ordinal: 10},
		"id_kec": {Col: "id_kec", Type: (*null.String)(nil), JSON: "id_kec", Label: "Id Kec", Ordinal: 11},
		"a_desa": {Col: "a_desa", Type: (*decimal.Decimal)(nil), JSON: "a_desa", Label: "A Desa", Ordinal: 12, DBRequired: true},
		"a_kelurahan": {Col: "a_kelurahan", Type: (*decimal.Decimal)(nil), JSON: "a_kelurahan", Label: "A Kelurahan", Ordinal: 13, DBRequired: true},
		"a_35": {Col: "a_35", Type: (*decimal.Decimal)(nil), JSON: "a_35", Label: "A 35", Ordinal: 14, DBRequired: true},
		"a_urban": {Col: "a_urban", Type: (*decimal.Decimal)(nil), JSON: "a_urban", Label: "A Urban", Ordinal: 15, DBRequired: true},
		"kategori_desa_id": {Col: "kategori_desa_id", Type: (*null.Int32)(nil), JSON: "kategori_desa_id", Label: "Kategori Desa Id", Ordinal: 16},
		"expired_date": {Col: "expired_date", Type: (*null.Time)(nil), JSON: "expired_date", Label: "Expired Date", Ordinal: 17},
		"synced_at": {Col: "synced_at", Type: (*time.Time)(nil), JSON: "synced_at", Label: "Synced At", Ordinal: 18, DBRequired: true},
	}
)