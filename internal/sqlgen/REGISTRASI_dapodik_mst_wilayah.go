package sqlgen

import (
	"fmt"
	"github.com/albatiqy/gopoh/pkg/lib/null"
	"github.com/albatiqy/gopoh/pkg/lib/decimal"
	"time"
)

type RegistrasiDapodikMstWilayah struct {

}

func (gen RegistrasiDapodikMstWilayah) Values(kodeWilayah string, nama string, idLevelWilayah int32, mstKodeWilayah null.String, negaraId string, asalWilayah null.String, kodeBps null.String, kodeDagri null.String, kodeKeu null.String, idProv null.String, idKabkota null.String, idKec null.String, aDesa decimal.Decimal, aKelurahan decimal.Decimal, a35 decimal.Decimal, aUrban decimal.Decimal, kategoriDesaId null.Int32, expiredDate null.Time, syncedAt time.Time) string {
	return fmt.Sprintf(`('%s','%s',%d,%s,'%s',%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)`, kodeWilayah, nama, idLevelWilayah, mstKodeWilayah, negaraId, asalWilayah, kodeBps, kodeDagri, kodeKeu, idProv, idKabkota, idKec, aDesa, aKelurahan, a35, aUrban, kategoriDesaId, expiredDate, syncedAt)
}