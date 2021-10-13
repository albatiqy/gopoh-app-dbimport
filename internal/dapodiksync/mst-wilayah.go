package dapodiksync

import (
	"github.com/albatiqy/gopoh/contract/repository/sqldb"
	"github.com/albatiqy/gopoh-app-dbimport/internal/sqlgen"
)

type MstWilayah struct {
	sqlGen *sqlgen.RegistrasiDapodikMstWilayah
}

func NewMstWilayah(srcSqlserverConn, dstPostgresConn *sqldb.Conn) *MstWilayah {
	return &MstWilayah{
		sqlGen: &sqlgen.RegistrasiDapodikMstWilayah{},
	}
}

func (dbo MstWilayah) CreateTable() error {
	return nil
}