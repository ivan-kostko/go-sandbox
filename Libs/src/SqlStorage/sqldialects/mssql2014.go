// SqlDialects project mssql2014.go
// Contains functionality to support dialect of MSSQL Server 2014
package SqlDialects

const MSSQL2014_DIALECT_REGISTER_NAME = "mssql2014"

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
