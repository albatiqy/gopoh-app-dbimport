package gendriver

import (
	"database/sql"
)

var LoadedEngines = make(map[string]Engine)

type Engine interface {
	ReadSchema(schemaName string, db *sql.DB) (*SchemaData, error)
	QuoteFuncTpl() string
	InsertValuesTpl(tableName string, cols []string) string
	SelectTpl(tableName string, cols []string) string
}

func Get(driverName string) Engine {
	engine, ok := LoadedEngines[driverName]
	if ok {
		return engine //.(Driver)
	}
	return nil
}

type ColData struct {
	Name             string
	CompatibleGoType interface{}
	Type             interface{}
	Nullable         bool
}

type SchemaData struct {
	ColsData []ColData
}

type TChar struct {
}

type TText struct {
}

type TVarChar struct {
}

type TCharUnicode struct {
}

type TVarCharUnicode struct {
}

type TTextUnicode struct {
}

type TDate struct {
}

type TTime struct {
}

type TTimestamp struct {
}

type TTimestampTz struct {
}

type TInt struct {
}

type TBigInt struct {
}

type TSmallInt struct {
}

type TTinyInt struct {
}

type TDecimal struct {
}