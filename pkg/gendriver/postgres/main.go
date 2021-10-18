package postgres

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
	"strconv"

	"github.com/albatiqy/gopoh-app-dbimport/pkg/gendriver"
	"github.com/albatiqy/gopoh/contract/log"
	"github.com/albatiqy/gopoh/pkg/lib/null"
	"github.com/albatiqy/gopoh/pkg/lib/decimal"
)

type rawField struct {
	TableCatalog           string
	TableSchema            string
	TableName              string
	ColumnName             string
	OrdinalPosition        int
	ColumnDefault          null.String
	IsNullable             string
	DataType               string
	CharacterMaximumLength null.Int32
	CharacterOctetLength   null.Int32
	NumericPrecision       null.Int32
	NumericPrecisionRadix  null.Int32
	NumericScale           null.Int32
	DatetimePrecision      null.Int32
	CharacterSetCatalog    null.String
	CharacterSetSchema     null.String
	CharacterSetName       null.String
	CollationCatalog       null.String
	CollationSchema        null.String
	CollationName          null.String
	DomainCatalog          null.String
	DomainSchema           null.String
	DomainName             null.String
	Ordinal                uint16
}

type Engine struct {
}

func (d Engine) ReadSchema(schemaName string, db *sql.DB) (*gendriver.SchemaData, error) {
	var tblSchema string
	if schema := strings.Split(schemaName, "."); len(schema) != 2 {
		return nil, fmt.Errorf("view/tabel harus mengandung schema ex: schema.nama_tabel")
	} else {
		tblSchema = schema[0]
		schemaName = schema[1]
	}

	fields := make(map[string]rawField)

	// harus order by ordinal

	rows, err := db.Query(fmt.Sprintf(`
	SELECT TABLE_CATALOG,TABLE_SCHEMA,TABLE_NAME,COLUMN_NAME,ORDINAL_POSITION,COLUMN_DEFAULT,IS_NULLABLE,DATA_TYPE,CHARACTER_MAXIMUM_LENGTH,CHARACTER_OCTET_LENGTH,NUMERIC_PRECISION,NUMERIC_PRECISION_RADIX,NUMERIC_SCALE,DATETIME_PRECISION,CHARACTER_SET_CATALOG,CHARACTER_SET_SCHEMA,CHARACTER_SET_NAME,COLLATION_CATALOG,COLLATION_SCHEMA,COLLATION_NAME,DOMAIN_CATALOG,DOMAIN_SCHEMA,DOMAIN_NAME
	FROM INFORMATION_SCHEMA.columns
	WHERE TABLE_NAME = '%s' AND TABLE_SCHEMA='%s'
	ORDER BY ORDINAL_POSITION
	`,
		schemaName, tblSchema))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ordinal := uint16(0)
	for rows.Next() {
		field := rawField{}
		err := rows.Scan(&field.TableCatalog, &field.TableSchema, &field.TableName, &field.ColumnName, &field.OrdinalPosition, &field.ColumnDefault, &field.IsNullable, &field.DataType, &field.CharacterMaximumLength, &field.CharacterOctetLength, &field.NumericPrecision, &field.NumericPrecisionRadix, &field.NumericScale, &field.DatetimePrecision, &field.CharacterSetCatalog, &field.CharacterSetSchema, &field.CharacterSetName, &field.CollationCatalog, &field.CollationSchema, &field.CollationName, &field.DomainCatalog, &field.DomainSchema, &field.DomainName)
		if err != nil {
			return nil, err
		}
		field.Ordinal = ordinal
		ordinal++
		fields[field.ColumnName] = field
	}

	colsData := make([]gendriver.ColData, len(fields))

	for _, field := range fields {
		nullable := (field.IsNullable == "YES")
		colsData[field.Ordinal].Name = field.ColumnName
		colsData[field.Ordinal].Nullable = nullable

		switch field.DataType {
		case "char":
			colsData[field.Ordinal].Type = gendriver.TChar{}
		case "varchar":
			colsData[field.Ordinal].Type = gendriver.TVarChar{}
		case "text":
			colsData[field.Ordinal].Type = gendriver.TText{}
		case "nchar":
			colsData[field.Ordinal].Type = gendriver.TCharUnicode{}
		case "nvarchar":
			colsData[field.Ordinal].Type = gendriver.TVarCharUnicode{}
		case "ntext":
			colsData[field.Ordinal].Type = gendriver.TTextUnicode{}
		case "date":
			colsData[field.Ordinal].Type = gendriver.TDate{}
		case "time":
			colsData[field.Ordinal].Type = gendriver.TTime{}
		case "datetime2":
			colsData[field.Ordinal].Type = gendriver.TTimestamp{}
		case "datetimeoffset":
			colsData[field.Ordinal].Type = gendriver.TTimestampTz{}
		case "int":
			colsData[field.Ordinal].Type = gendriver.TInt{}
		case "smallint":
			colsData[field.Ordinal].Type = gendriver.TSmallInt{}
		case "tinyint":
			colsData[field.Ordinal].Type = gendriver.TTinyInt{}
		case "decimal", "numeric":
			colsData[field.Ordinal].Type = gendriver.TDecimal{}
		default:
			return nil, fmt.Errorf("type " + field.DataType + " tidak terdefinisi")
		}
	}
	return &gendriver.SchemaData{
		ColsData: colsData,
	}, nil
}

func (d Engine) QuoteString(val string) string {
	return strings.ReplaceAll(val, "'", "''")
}

func (d Engine) Quote(val interface{}) string {
	switch val := val.(type) {
	case time.Time:
		return "'" + val.Format("2006-01-02 15:04:05") + "'"
	case null.String:
		if val.Valid {
			return "'" + d.QuoteString(val.String) + "'"
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

func (d Engine) CreateTable(tableName string, schemaData *gendriver.SchemaData) string {
	return ""
}

func (d Engine) InsertValuesTpl(tableName string, cols []string) string {
	return "INSERT INTO " + tableName + " (" + strings.Join(cols, ",") + ") VALUES %s"
}

func (d Engine) SelectTpl(tableName string, cols []string) string {
	return "SELECT " + strings.Join(cols, ",") + " FROM " + tableName
}

func (d Engine) InsertPlaceholders(fieldType []interface{}) string {
	result := make([]string, len(fieldType))
	for i, t := range fieldType {
		switch t.(type) {
		case *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64:
			result[i] = `%d`
		case *float32, *float64:
			result[i] = `%f`
		case *string:
			result[i] = `'%s'`
		default:
			result[i] = `%s` // need quote
		}
	}
	return strings.Join(result, ",")
}

func init() {
	gendriver.LoadedEngines["postgres"] = Engine{}
	log.Debugf("postgres gen: %s", "initialized")
}
