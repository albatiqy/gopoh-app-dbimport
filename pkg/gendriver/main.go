package gendriver

import (
	"database/sql"
)

var LoadedEngines = make(map[string]Engine)

type Engine interface {
	ReadSchema(schemaName string, db *sql.DB) (*SchemaData, error)
	QuoteString(val string) string
	Quote(val interface{}) string
	CreateTable(tableName string, schemaData *SchemaData) string
	InsertValuesTpl(tableName string, cols []string) string
	SelectTpl(tableName string, cols []string) string
	InsertPlaceholders(fieldType []interface{}) string
}

type GoTypeCompatible interface {
	GoType() interface{}
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
	Type             interface{}
	Nullable         bool
}

type SchemaData struct {
	ColsData []ColData
}

type TChar struct {
	compatibleGoType interface{}
	MaximumLength int32
}

func (t TChar) GoType() interface{} {
	return t.compatibleGoType
}

func NewTChar(len int32) TChar {
	return TChar{
		compatibleGoType: (*string)(nil),
		MaximumLength: len,
	}
}

type TText struct {
}

type TVarChar struct {
	compatibleGoType interface{}
	MaximumLength int32
}

func (t TVarChar) GoType() interface{} {
	return t.compatibleGoType
}

func NewTVarChar(len int32) TChar {
	return TChar{
		compatibleGoType: (*string)(nil),
		MaximumLength: len,
	}
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