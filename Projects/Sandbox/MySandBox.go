package MySendBox

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	//"text/template"
	"time"
	"unsafe"
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

func CustomJoinStrings(ss ...string) (r string) {
	for i, s := range ss {
		if i != 0 {
			r += ","
		}
		r += s
	}
	return r
}

// Represents subset of structure fields
type FieldSubset struct {
	Name        string
	Type        reflect.Type
	fieldsIds   []int
	fieldsNames []string
}

// Creates new FieldsSubset based on the sample pointer and the samples fields pointers
// It returns InvalidOperation Error if any of fields does not belong to the sample
func NewFieldsSubsetByReflection(name string, sample interface{}, sampleFields ...interface{}) (FieldSubset, error) {
	typ := reflect.TypeOf(sample).Elem()
	ret := FieldSubset{Name: name, Type: typ}
	c := len(sampleFields)
	// if no fields provided return empty FieldSubset
	if c == 0 {
		return ret, nil
	}
	fids := make([]int, c, c)
	sampleFirstPtr := reflect.ValueOf(sample).Pointer()
	sampleLeastPtr := sampleFirstPtr + typ.Size()

	for i := 0; i < c; i++ {
		sfPtr := reflect.ValueOf(sampleFields[i]).Pointer()
		if sfPtr < sampleFirstPtr || sampleLeastPtr <= sfPtr {
			return ret, fmt.Errorf("InvalidOperation")
		}
		for fi := 0; fi < typ.NumField(); fi++ {
			if sfPtr == sampleFirstPtr+typ.Field(fi).Offset {
				fids[i] = fi
				break
			}
		}
	}
	ret.fieldsIds = fids
	return ret, nil
}

func NewFieldsSubsetByUnsafe(name string, sample interface{}, sampleFields ...interface{}) (FieldSubset, error) {
	typ := reflect.TypeOf(sample)
	ret := FieldSubset{Name: name, Type: typ}
	c := len(sampleFields)
	// if no fields provided return empty FieldSubset
	if c == 0 {
		return ret, nil
	}
	fids := make([]int, c, c)
	sampleFirstPtr := unsafe.Pointer(&sample)
	sampleLeastPtr := uintptr(sampleFirstPtr) + unsafe.Sizeof(sample)

	for i := 0; i < c; i++ {
		sfPtr := uintptr(unsafe.Pointer(&(sampleFields[i]))) // reflect.ValueOf(sampleFields[i]).Pointer()
		if sfPtr < uintptr(sampleFirstPtr) || sampleLeastPtr <= sfPtr {
			return ret, fmt.Errorf("InvalidOperation %v %v %v | &sample = %v", uintptr(sampleFirstPtr), sfPtr, sampleLeastPtr, &sample)
		}
		for fi := 0; fi < typ.NumField(); fi++ {
			if sfPtr == uintptr(sampleFirstPtr)+typ.Field(fi).Offset {
				fids[i] = fi
				break
			}
		}
	}
	ret.fieldsIds = fids
	return ret, nil
}

func NewFieldsSubset(name string, sample interface{}, sampleFields ...interface{}) (FieldSubset, error) {

	return NewFieldsSubsetByReflection(name, sample, sampleFields...)
	//return NewFieldsSubsetByUnsafe(name, sample, sampleFields...)
}
