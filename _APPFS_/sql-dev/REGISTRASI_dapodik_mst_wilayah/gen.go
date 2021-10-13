package main

import (
	_ "embed"
	"reflect"
	// "fmt"
	"path/filepath"
	// "reflect"
	"strings"
	"time"
	"unicode"

	// "github.com/albatiqy/gopoh-app-dbimport/pkg/gendriver"

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
		fieldsLen          = len(obj.FieldDefs)
		imports            []string
		useImport          = map[string]string{}
		valuesPlaceholders = make([]string, fieldsLen)
		valuesArgs         = make([]string, fieldsLen)
		valuesVars         = make([]string, fieldsLen)
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
		case *string:
			valuesVars[field.Ordinal] = varName
		case *null.String:
			
		}

		switch field.Type.(type) {
		case *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64:
			valuesPlaceholders[field.Ordinal] = `%d`
		case *float32, *float64:
			valuesPlaceholders[field.Ordinal] = `%f`
		case *string:
			valuesPlaceholders[field.Ordinal] = `'%s'`
		default:
			valuesPlaceholders[field.Ordinal] = `%s` // need quote
		}

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

	// genDriver := gendriver.Get(obj.DBDriver)

	tplData := map[string]string{
		"imports":            strings.Join(imports, "\n"),
		"genStructName":      genStructName,
		"objectPackage":      obj.ObjectPackage,
		"newJsonKeyAttr":     newJsonKeyAttr,
		"valuesArgs":         strings.Join(valuesArgs, ", "),
		"valuesPlaceholders": strings.Join(valuesPlaceholders, ","),
		"valuesVars":         strings.Join(valuesVars, ", "),
	}

	util.WriteTplFile(fnameOut, txtObjectGen, tplData)

}
