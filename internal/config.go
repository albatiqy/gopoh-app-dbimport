package internal

import (
	"os"
	"path/filepath"
	"strings"
	"fmt"

	"github.com/albatiqy/gopoh/pkg/lib/fs"
	"github.com/albatiqy/gopoh/contract/repository/sqldb"
)

type Config struct {
	AppFSRoot      string
	HTTPListenBind string
	HTTPBasePath   string
	dbSettings     map[string]*sqldb.DBSetting
}

func (cfg *Config) GetDBSetting(envKey string) *sqldb.DBSetting {
	setting, ok := cfg.dbSettings[envKey]
	if !ok {
		driverName := os.Getenv("DB_" + envKey + "_DRIVER")
		if driverName == "" {
			return nil
		}

		setting = &sqldb.DBSetting{
			DSN:       os.Getenv("DB_" + envKey + "_DSN"),
			DriverName: driverName,
		}

		cfg.dbSettings[envKey] = setting
	}
	return setting
}

/*
func (cfg Config) readEnvKeyString(envKey string, cfgField *string) {
	if val, ok := os.LookupEnv(envKey); ok {
		*cfgField = val
	}
}
*/

func NewConfig(envDir string) (*Config, error) {
	cfg := Config{
		AppFSRoot:  "./_APPFS_",
		dbSettings: map[string]*sqldb.DBSetting{},
	}
	if val, ok := os.LookupEnv("APP_FS_ROOT"); ok {
		cfg.AppFSRoot = val
	}
	if strings.HasPrefix(cfg.AppFSRoot, ".") {
		cfg.AppFSRoot = filepath.Join(envDir, cfg.AppFSRoot)
	}
	fileInfo := fs.FileInfo(cfg.AppFSRoot)
	if fileInfo == nil {
		return nil, fmt.Errorf("FS ROOT tidak ditemukan")
	}
	if !fileInfo.IsDir() {
		return nil, fmt.Errorf("FS ROOT bukan direktori")
	}
	// cfg.readEnvKeyString("HTTP_LISTEN_BIND", &cfg.HTTPListenBind)
	return &cfg, nil
}