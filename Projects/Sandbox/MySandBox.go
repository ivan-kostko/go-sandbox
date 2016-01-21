package MySendBox

import (
	"fmt"
	"reflect"
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
	case int:
		return FuncForInteger(x) /* "integer"*/
		break
	case float64:
		return FuncForFloat(x) /*"float"*/
		break
	case bool:
		return FuncForBool(x) /*"bit"*/
		break
	case []byte:
		return FuncForBytes(x) /*"binary"*/
		break
	default:
		return fmt.Sprintf("DEFAULT: %T %#v", x, x)
	}
	return ""
}
