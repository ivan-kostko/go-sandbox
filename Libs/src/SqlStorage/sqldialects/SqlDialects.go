// SqlDialects project SqlDialects.go
package SqlDialects

import (
	. "customErrors"
)

const (
	ERR_NILDIALECT_REGISTER     = "Won't register nil as dialect"
	ERR_NOONEDIALECT_REGISTERED = "There are noone registered dialect"
	ERR_DIALECTNOTFOUND         = "Requested dialect is not found in list of registered"
)

// Represents Singleton supported dialects regestry
// NB: is lazy initialized (or in other words - initialized on demand)
var supportedSqlDialects map[string]ISqlDialect

// Sql Script / query string type
type SqlScriptString string

// Sql Dialect shared interface
type ISqlDialect interface {
	BuildSelectAllColumnsSqlScriptString(tableName string) SqlScriptString
}

// Implements ISqlDialect query generation functionality
type SqlDialect struct {
	convertIntoSqlScriptString           func(interface{}) (SqlScriptString, *Error)
	buildInsertSqlScriptString           func(tableName, columnList, valuesList SqlScriptString) SqlScriptString
	buildSelectSqlScriptString           func(tableName, columnList, whereStmt SqlScriptString, limit int) SqlScriptString
	buildSelectAllColumnsSqlScriptString func(tableName string) SqlScriptString
}

// Builds query to get empty resultset but with all columns
// Usually is used for schema validation on mapping
func (sd *SqlDialect) BuildSelectAllColumnsSqlScriptString(tableName string) SqlScriptString {
	return sd.buildSelectAllColumnsSqlScriptString(tableName)
}

// Registers dialect as supported
func RegisterSupportOfDialect(alias string, dialect ISqlDialect) *Error {
	if dialect == nil {
		return NewError(InvalidArgument, ERR_NILDIALECT_REGISTER)
	}

	if supportedSqlDialects == nil {
		supportedSqlDialects = make(map[string]ISqlDialect)
	}

	supportedSqlDialects[alias] = dialect

	return nil
}

// Gets registered/supported dialect by alias. Returns error if alias is not found as registered
func GetDialectByAlias(alias string) (ISqlDialect, *Error) {
	if supportedSqlDialects == nil {
		return nil, NewError(Nonsupported, ERR_NOONEDIALECT_REGISTERED)
	}
	d, ok := supportedSqlDialects[alias]
	if !ok {
		return nil, NewError(Nonsupported, ERR_DIALECTNOTFOUND)
	}
	return d, nil
}
