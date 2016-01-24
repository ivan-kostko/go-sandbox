package MySendBox

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
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
