package MySendBox

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	//"text/template"
	"time"
)

var FuncMap map[reflect.Type]func(interface{}) string

func FuncForString(x interface{}) string {
	return "string"
}

func FuncForInteger(x interface{}) string {
	return "integer"
}

func FuncForFloat(x interface{}) string {
	return "float"
}

func FuncForBool(x interface{}) string {
	return "bit"
}

func FuncForBytes(x interface{}) string {
	return "binary"
}

func ConvertBySwitch(i interface{}) string {
	switch x := i.(type) {
	case string:
		return FuncForString(x) /*"string"*/
		break
	case *string:
		return FuncForString(*x) /*"string"*/
		break
	case int:
		return FuncForInteger(x) /* "integer"*/
		break
	case *int:
		return FuncForInteger(*x) /* "integer"*/
		break
	case float64:
		return FuncForFloat(x) /*"float"*/
		break
	case *float64:
		return FuncForFloat(*x) /*"float"*/
		break
	case float32:
		return FuncForFloat(x) /*"float"*/
		break
	case *float32:
		return FuncForFloat(*x) /*"float"*/
		break
	case time.Time:
		return "time"
		break
	case *time.Time:
		return "time"
		break
	case bool:
		return FuncForBool(x)
		break
	case []byte:
		return FuncForBytes(x)
		break
	default:
		return fmt.Sprintf("DEFAULT: %T %#v", x, x)
	}
	return ""
}

func ConvertIntToStringBySprintf(i int) string {
	return fmt.Sprintf("%v", i)
}

func ConvertIntToStringByStrConv(i int) string {
	return strconv.Itoa(i)
}

func ConvertFloatToStringBySprintf(i float64) string {
	return fmt.Sprintf("%f64", i)
}

func ConvertFloatToStringByStrConv(i float64) string {
	return strconv.FormatFloat(i, 'f', -1, 32)
}

func ConvertBytesIntoSqlString(v []byte) string {
	return "0x" + hex.EncodeToString(v)
}

func ConvertInt64ToStringBySprintf(i int64) string {
	return fmt.Sprintf("%i64", i)
}
func ConvertInt64ToStringByStrConv(i int64) string {
	return strconv.FormatInt(i, 10)
}

func ConvertTimeToStringBySprintf(i time.Time) string {
	return fmt.Sprintf("%v", i)
}
func ConvertTimeToStringByFormat(i time.Time) string {
	return i.Format("2006-01-02T15:04:05Z07:00")
}
func ConvertTimeToStringByStringer(i time.Time) string {
	return i.String()
}

const REPLACE_TEMPLATE1 = "INSERT INTO {{.TableName}} ({{.ColumnLins}}) VALUES({{.ValuesList}})"
const REPLACE_TEMPLATE2 = "INSERT INTO %s (%s) VALUES(%s)"

func GenerateStringsReplace(tableName, columnList, valuesList string) string {
	r := strings.Replace(REPLACE_TEMPLATE1, "{{.TableName}}", tableName, -1)
	r = strings.Replace(r, "{{.ColumnLins}}", columnList, -1)
	r = strings.Replace(r, "{{.ValuesList}}", valuesList, -1)
	return r
}

func GenerateFmtSprintf(tableName, columnList, valuesList string) string {
	return fmt.Sprintf(REPLACE_TEMPLATE2, tableName, columnList, valuesList)
}

func GenerateCustom(tableName, columnList, valuesList string) string {
	return "INSERT INTO " + tableName + " (" + columnList + ") VALUES(" + valuesList + ")"
}
