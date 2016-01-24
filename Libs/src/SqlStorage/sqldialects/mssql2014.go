// SqlDialects project mssql2014.go
// Contains functionality to support dialect of MSSQL Server 2014
package SqlDialects

import (
	. "customErrors"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

const (
	MSSQL2014_DIALECT_REGISTER_NAME = "mssql2014"
	MSSQL2014_DBNULL_SCRPT_STRING   = "NULL"
	MSSQL2014_TIMEPARSE_TEMPLATE    = "20060102 15:04:05.99999999 Z07:00"
)

func init() {
	d := GetMsSql2014Dialect()
	RegisterSupportOfDialect(MSSQL2014_DIALECT_REGISTER_NAME, d)
}

type MsSql2014Dialect struct {
}

// MSSQL2014 dialect factory
func GetMsSql2014Dialect() ISqlDialect {
	return new(MsSql2014Dialect)
}

// converts input i into Sql Script string
func convertSomethingIntoMssql2014SqlScriptString(i interface{}) (SqlScriptString, *Error) {
	if i == nil {
		return MSSQL2014_DBNULL_SCRPT_STRING, nil
	}
	var r SqlScriptString
	switch x := i.(type) {
	case string:
		r = SqlScriptString("N'" + strings.Replace(x, "'", "''", -1) + "'")
		break
	case *string:
		r = SqlScriptString("N'" + strings.Replace(*x, "'", "''", -1) + "'")
		break

	case int:
		r = SqlScriptString(strconv.Itoa(x))
		break
	case *int:
		r = SqlScriptString(strconv.Itoa(*x))
		break
	case int64:
		r = SqlScriptString(strconv.FormatInt(x, 10))
		break
	case *int64:
		r = SqlScriptString(strconv.FormatInt(*x, 10))
		break

	case float32:
		r = SqlScriptString(strconv.FormatFloat(float64(x), 'f', -1, 32))
		break
	case *float32:
		r = SqlScriptString(strconv.FormatFloat(float64(*x), 'f', -1, 32))
		break
	case float64:
		r = SqlScriptString(strconv.FormatFloat(x, 'f', -1, 64))
		break
	case *float64:
		r = SqlScriptString(strconv.FormatFloat(*x, 'f', -1, 64))
		break

	case time.Time:
		r = SqlScriptString("TRY_CAST('" + x.Format(MSSQL2014_TIMEPARSE_TEMPLATE) + "' AS DATETIMEOFFSET)")
		break
	case *time.Time:
		r = SqlScriptString("TRY_CAST('" + x.Format(MSSQL2014_TIMEPARSE_TEMPLATE) + "' AS DATETIMEOFFSET)")
		break

	case bool:
		if x {
			r = "1"
		} else {
			r = "0"
		}
		break
	case *bool:
		if *x {
			r = "1"
		} else {
			r = "0"
		}
		break

	case []byte:
		r = SqlScriptString("0x" + hex.EncodeToString(x))
		break
	case *[]byte:
		r = SqlScriptString("0x" + hex.EncodeToString(*x))
		break

	default:
		return "?", NewError(Nonsupported, "Unsupported type")
	}

	return r, nil

}
