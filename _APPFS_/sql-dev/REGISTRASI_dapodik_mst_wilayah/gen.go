package main

import (
	_ "embed"
	"fmt"
	"path/filepath"
	"reflect"
	"strings"
	"time"
	"unicode"

	"github.com/albatiqy/gopoh/contract/gen/driver"
	"github.com/albatiqy/gopoh/contract/gen/util"
	"github.com/albatiqy/gopoh/contract/log"
	"github.com/albatiqy/gopoh/pkg/lib/decimal"
	"github.com/albatiqy/gopoh/pkg/lib/fs"
	"github.com/albatiqy/gopoh/pkg/lib/null"
)

type genDev struct {
	DBEnvKey             string
	DBDriver             string
	TableName            string
	ObjectName           string
	ObjectPackage        string
	FieldDefs            map[string]driver.FieldDef
	KeyAttr              string
	KeyAuto              bool
	KeyCanUpdate         bool
	SoftDelete           bool
	OverridesStructField map[string]string
	OverridesType        map[string]interface{}
	OverridesJSON        map[string]string
	OverridesLabel       map[string]string
}

var (
	//go:embed _embed/object-gen.txt
	txtObjectGen string
)

func (obj genDev) Generate(pathPrjDir string) {
	modName := util.GetModName(pathPrjDir)
	if modName == "" {
		log.Fatal("direktori project tidak valid")
	}

	pathSaveRoot := filepath.Join(pathPrjDir, "internal", obj.ObjectPackage)
	if success, err := fs.MkDirIfNotExists(pathSaveRoot); !success {
		log.Fatal(err)
	}

	nsName := obj.DBEnvKey + "_" + strings.Replace(obj.TableName, ".", "_", 1)

	fnameOut := filepath.Join(pathSaveRoot, nsName+".go")

	genStructName := strings.ReplaceAll(obj.ObjectName, "_", " ")
	genStructName = strings.ToLower(genStructName)
	genStructName = strings.Title(genStructName)
	genStructName = strings.ReplaceAll(genStructName, " ", "")

	var (
		fieldsLen       = len(obj.FieldDefs)
		imports         []string
		useImport       = map[string]string{}
		valuesArgs      = make([]string, fieldsLen)
		valuesVars      = make([]string, fieldsLen)
		strFieldsModel  = make([]string, fieldsLen)
		fieldScansModel = make([]string, fieldsLen)
		fieldModel      = make([]string, fieldsLen)
		tableCols       = make([]string, fieldsLen)
		fieldTypes      = make([]string, fieldsLen)
		qTimeLocalModel []string
	)
	newJsonKeyAttr := obj.KeyAttr
	for attr, field := range obj.FieldDefs {

		switch field.Type.(type) {
		case *time.Time:
			useImport["time"] = ""
		case *null.String:
			useImport["null"] = ""
		case *decimal.Decimal, *decimal.NullDecimal:
			useImport["decimal"] = ""
		}

		if obj.OverridesType != nil {
			if fieldType, ok := obj.OverridesType[attr]; ok {
				field.Type = fieldType
			}
		}

		if obj.OverridesJSON != nil {
			if fieldJSON, ok := obj.OverridesJSON[attr]; ok {
				field.JSON = fieldJSON
			}
		}

		if attr == obj.KeyAttr {
			newJsonKeyAttr = field.JSON
		}

		if obj.OverridesLabel != nil {
			if fieldLabel, ok := obj.OverridesLabel[attr]; ok {
				field.Label = fieldLabel
			}
		}

		field.StructField = strings.ReplaceAll(field.Label, " ", "")
		if obj.OverridesStructField != nil {
			if structField, ok := obj.OverridesStructField[attr]; ok {
				field.StructField = structField
			}
		}

		r := []rune(field.StructField)
		r[0] = unicode.ToLower(r[0])
		varName := string(r)
		fieldTypeStr := reflect.TypeOf(field.Type).Elem().String()
		valuesArgs[field.Ordinal] = varName + " " + fieldTypeStr
		switch field.Type.(type) {
		/*
			case *string, *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *float32, *float64:
				valuesVars[field.Ordinal] = varName
		*/
		case *string:
			valuesVars[field.Ordinal] = "gen.genDriver.QuoteString(" + varName + ")"
		case *time.Time, *null.String, *null.Time, *null.Bool, *null.Int32, *null.Int64, *null.Float64, *decimal.Decimal, *decimal.NullDecimal:
			valuesVars[field.Ordinal] = "gen.genDriver.Quote(" + varName + ")"
		default:
			valuesVars[field.Ordinal] = varName
		}

		switch field.Type.(type) {
		case *time.Time:
			fieldScansModel[field.Ordinal] = "\t\t\t&record." + field.StructField + ", // warning from UTC result"
			qTimeLocalModel = append(qTimeLocalModel, "// \trecord."+field.StructField+" = record."+field.StructField+".Local() // convert to local")
		case *null.Time:
			fieldScansModel[field.Ordinal] = "\t\t\t&record." + field.StructField + ", // warning from UTC result"
			qTimeLocalModel = append(qTimeLocalModel, "// \trecord."+field.StructField+" = null.NewTime(record."+field.StructField+".Time.Local(), record."+field.StructField+".Valid) // convert to local")
		default:
			fieldScansModel[field.Ordinal] = "\t\t\t&record." + field.StructField + ","
		}
		strFieldsModel[field.Ordinal] = fmt.Sprintf("\t%[1]s %[2]s `json:\"%[3]s\"`", field.StructField, fieldTypeStr, field.JSON)

		tableCols[field.Ordinal] = "\t\t\t\"" + field.Col + "\","
		fieldTypes[field.Ordinal] = "\t\t\t(*" + fieldTypeStr + ")(nil),"
		fieldModel[field.Ordinal] = "\t\t\trecord." + field.StructField + ","

		obj.FieldDefs[attr] = field
	}

	for impk, impv := range useImport {
		if impv != "" {
			imports = append(imports, "\t\""+impv+`"`)
		} else {
			if impl, ok := util.ImportsMap[impk]; ok {
				imports = append(imports, "\t\""+impl+`"`)
			}
		}
	}

	tplData := map[string]string{
		"imports":            strings.Join(imports, "\n"),
		"genStructName":      genStructName,
		"dbDriver":           obj.DBDriver,
		"objectPackage":      obj.ObjectPackage,
		"tableName":          obj.TableName,
		"newJsonKeyAttr":     newJsonKeyAttr, // unused
		"valuesArgs":         strings.Join(valuesArgs, ", "),
		"tableCols":          strings.Join(tableCols, "\n"),
		"fieldTypes":         strings.Join(fieldTypes, "\n"),
		"valuesVars":         strings.Join(valuesVars, ", "),
		"fieldScansModel":    strings.Join(fieldScansModel, "\n"),
		"fieldModel":         strings.Join(fieldModel, "\n"),
		"qTimeLocalModel":    strings.Join(qTimeLocalModel, "\n"),
		"fieldsModel":        strings.Join(strFieldsModel, "\n"),
	}

	util.WriteTplFile(fnameOut, txtObjectGen, tplData)
}
