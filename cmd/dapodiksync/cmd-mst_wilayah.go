package main

import (
	"strings"

	"github.com/mitchellh/cli"

	"github.com/albatiqy/gopoh/contract/repository/sqldb"
	_ "github.com/albatiqy/gopoh/contract/repository/sqldb/sqlserver"
	_ "github.com/albatiqy/gopoh/contract/repository/sqldb/postgres"

	"github.com/albatiqy/gopoh-app-dbimport/internal/dapodiksync"
)

type cmdMstWilayah struct {
	Ui cli.Ui
}

func (cmd *cmdMstWilayah) Help() string {
	helpText := `
Usage: dapodiksync mst_wilayah operation
	Dispatches a custom event across the Serf cluster.
Operation:
	createtbl
`
	return strings.TrimSpace(helpText)
}

func (cmd *cmdMstWilayah) Run(args []string) int {

	if len(args) < 1 {
		cmd.Ui.Error("An operation must be specified.")
		cmd.Ui.Error("")
		cmd.Ui.Error(cmd.Help())
		return 1
	}

	postgresCon := cmd.getPostgresDB("REGISTRASI")
	if postgresCon == nil {
		return 1
	}

	sqlserverCon := cmd.getSqlserverDB("DEFAULT")
	if sqlserverCon == nil {
		return 1
	}

	mstWilayah := dapodiksync.NewMstWilayah(sqlserverCon, postgresCon)

	switch args[0] {
	case "createtbl":
		if err := mstWilayah.CreateTable(); err != nil {
			cmd.Ui.Error(err.Error())
			return 1
		}
	case "import":
		mstWilayah.Import(cmd.Ui)
	default:
		cmd.Ui.Error("Invalid operation.")
		cmd.Ui.Error("")
		cmd.Ui.Error(cmd.Help())
	}

	return 0
}

func (cmd *cmdMstWilayah) Synopsis() string {
	return "Send a custom event through the Serf cluster"
}

func (cmd *cmdMstWilayah) getPostgresDB(envKey string) *sqldb.Conn {
	dbSetting := cfg.GetDBSetting(envKey)
	if dbSetting.DriverName != "postgres" {
		cmd.Ui.Warn(`sqldb: driver tidak cocok`)
		return nil
	}
	dbConn := sqldb.NewConn(dbSetting)
	if dbConn == nil {
		cmd.Ui.Warn("tidak dapat terkoneksi dengan database")
		return nil
	}

	return dbConn
}

func (cmd *cmdMstWilayah) getSqlserverDB(envKey string) *sqldb.Conn {
	dbSetting := cfg.GetDBSetting(envKey)
	if dbSetting.DriverName != "sqlserver" {
		cmd.Ui.Warn(`sqldb: driver tidak cocok`)
		return nil
	}
	dbConn := sqldb.NewConn(dbSetting)
	if dbConn == nil {
		cmd.Ui.Warn("tidak dapat terkoneksi dengan database")
		return nil
	}

	return dbConn
}

func init() {
	cmds["mst_wilayah"] = func() (cli.Command, error) {
		return &cmdMstWilayah{Ui: ui}, nil
	}
}
