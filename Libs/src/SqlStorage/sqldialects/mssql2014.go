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
	return &SqlDialect{
		convertIntoSqlScriptString:           convertSomethingIntoMssql2014SqlScriptString,
		buildInsertSqlScriptString:           buildMSSQL2014InsertSqlScriptString,
		buildSelectSqlScriptString:           buildMSSQL2014SelectSqlScriptString,
		buildWhereSqlScriptString:            buildMSSQL2014WhereSqlScriptString,
		buildSelectAllColumnsSqlScriptString: buildMSSQL2014SelectAllColumnsSqlScriptString,
		buildColumnsListSqlScriptString:      buildMSSQL2014ColumnsListSqlScriptString,
	}

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
		if x == (*string)(nil) {
			r = MSSQL2014_DBNULL_SCRPT_STRING
		} else {
			r = SqlScriptString("N'" + strings.Replace(*x, "'", "''", -1) + "'")
		}
		break

	case int:
		r = SqlScriptString(strconv.Itoa(x))
		break
	case *int:
		if x == (*int)(nil) {
			r = MSSQL2014_DBNULL_SCRPT_STRING
		} else {
			r = SqlScriptString(strconv.Itoa(*x))
		}
		break
	case int64:
		r = SqlScriptString(strconv.FormatInt(x, 10))
		break
	case *int64:
		if x == (*int64)(nil) {
			r = MSSQL2014_DBNULL_SCRPT_STRING
		} else {
			r = SqlScriptString(strconv.FormatInt(*x, 10))
		}
		break

	case float32:
		r = SqlScriptString(strconv.FormatFloat(float64(x), 'f', -1, 32))
		break
	case *float32:
		if x == (*float32)(nil) {
			r = MSSQL2014_DBNULL_SCRPT_STRING
		} else {
			r = SqlScriptString(strconv.FormatFloat(float64(*x), 'f', -1, 32))
		}
		break
	case float64:
		r = SqlScriptString(strconv.FormatFloat(x, 'f', -1, 64))
		break
	case *float64:
		if x == (*float64)(nil) {
			r = MSSQL2014_DBNULL_SCRPT_STRING
		} else {
			r = SqlScriptString(strconv.FormatFloat(*x, 'f', -1, 64))
		}
		break

	case time.Time:
		r = SqlScriptString("TRY_CAST('" + x.Format(MSSQL2014_TIMEPARSE_TEMPLATE) + "' AS DATETIMEOFFSET)")
		break
	case *time.Time:
		if x == (*time.Time)(nil) {
			r = MSSQL2014_DBNULL_SCRPT_STRING
		} else {
			r = SqlScriptString("TRY_CAST('" + x.Format(MSSQL2014_TIMEPARSE_TEMPLATE) + "' AS DATETIMEOFFSET)")
		}
		break

	case bool:
		if x {
			r = "1"
		} else {
			r = "0"
		}
		break
	case *bool:
		if x == (*bool)(nil) {
			r = MSSQL2014_DBNULL_SCRPT_STRING
		} else {
			if *x {
				r = "1"
			} else {
				r = "0"
			}
		}
		break

	case []byte:
		r = SqlScriptString("0x" + hex.EncodeToString(x))
		break
	case *[]byte:
		if x == (*[]byte)(nil) {
			r = MSSQL2014_DBNULL_SCRPT_STRING
		} else {
			r = SqlScriptString("0x" + hex.EncodeToString(*x))
		}
		break

	default:
		return "?", NewError(Nonsupported, "Unsupported type")
	}

	return r, nil

}

// Builds Insert Sql Script(statement)
func buildMSSQL2014InsertSqlScriptString(tableName, columnList, valuesList SqlScriptString) SqlScriptString {
	return "INSERT INTO " + tableName + "(" + columnList + ") VALUES(" + valuesList + ")"
}

// if limit <0 selects all
func buildMSSQL2014SelectSqlScriptString(tableName, columnList, whereStmt SqlScriptString, limit int) (query SqlScriptString) {
	limitStr := SqlScriptString("")
	if limit >= 0 {
		limitStr = SqlScriptString("TOP (" + strconv.Itoa(limit) + ") ")
	}
	query = "SELECT " + limitStr + columnList + " FROM " + tableName
	if whereStmt != "" {
		query += " WHERE " + whereStmt
	}
	return query
}

func buildMSSQL2014WhereSqlScriptString(columnList, valuesList []SqlScriptString) SqlScriptString {
	var where SqlScriptString
	if len(columnList) != len(valuesList) {
		return where
	}
	for i, c := range columnList {
		if i != 0 {
			where += " AND "
		}
		if valuesList[i] == SqlScriptString(MSSQL2014_DBNULL_SCRPT_STRING) {
			where += c + " IS " + SqlScriptString(MSSQL2014_DBNULL_SCRPT_STRING)
		} else {
			where += c + " = " + valuesList[i]
		}
	}
	return where
}

func buildMSSQL2014SelectAllColumnsSqlScriptString(tablename string) SqlScriptString {
	return SqlScriptString("SELECT TOP(0) * FROM " + tablename)
}

// buildMSSQL2014ColumnsListSqlScriptString
func buildMSSQL2014ColumnsListSqlScriptString(cols []string) (ret SqlScriptString) {
	for i, c := range cols {
		if i != 0 {
			ret += ","
		}
		ret += SqlScriptString("[" + c + "]")
	}
	return
}
