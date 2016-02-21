// SqlStorage project SqlDatabase.go
package SqlStorage

import (
	"Logger"
	. "customErrors"

	"database/sql"
	"github.com/jmoiron/sqlx" // opensource Sql Extentions lib.

	_ "github.com/alexbrainman/odbc" // mssql or freetds. Registred as odbc
	_ "github.com/lib/pq"            // postgreSql. Registred as postrgres
)

const (
	ERR_CONNOTESTABLISHED = "The connection to database is not established yet"
	ERR_PINGFUNCNOTSET    = "The Ping() function is not set"
	ERR_QUERYFUNCNOTSET   = "The Query() function is not set"
)

var getNewSqlConnection func(driverName, dataSourceName string) (*SqlDatabase, *Error)

func init() {
	getNewSqlConnection = defaultGetNewSqlConnection
}

// Sql database functionality wrapper
type SqlDatabase struct {
	conn              *sqlx.DB
	driverName        string
	dataSourceName    string
	query             func(sdb *sqlx.DB, query string, args ...interface{}) (*sql.Rows, *Error)
	ping              func(sdb *sqlx.DB) *Error
	scanRowsIntoSlice func(rows *sql.Rows, sl []interface{}) *Error
	log               Logger.ILogger
}

func GetNewSqlDatabase(deriverName, dataSourceName string) (*SqlDatabase, *Error) {
	return getNewSqlConnection(deriverName, dataSourceName)
}

func (sd *SqlDatabase) Ping() *Error {
	if sd.ping == nil {
		return NewError(AccessViolation, ERR_PINGFUNCNOTSET)
	}
	return sd.ping(sd.conn)
}

func (sd *SqlDatabase) Query(query string, args ...interface{}) (*sql.Rows, *Error) {
	if sd.query == nil {
		return nil, NewError(AccessViolation, ERR_QUERYFUNCNOTSET)
	}
	return sd.query(sd.conn, query, args...)
}

func (sd *SqlDatabase) Close() *Error {
	if sd.conn == nil {
		return NewError(AccessViolation, ERR_CONNOTESTABLISHED)
	}
	err := sd.conn.Close()
	if err != nil {
		return NewError(InvalidOperation, err.Error())
	}
	return nil
}

func defaultGetNewSqlConnection(driverName, dataSourceName string) (*SqlDatabase, *Error) {
	conn, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		return nil, NewError(InvalidOperation, err.Error())
	}
	sd := SqlDatabase{
		conn:           conn,
		driverName:     driverName,
		dataSourceName: dataSourceName,
		query:          query,
		ping:           ping,
	}
	return &sd, nil
}

func query(sdb *sqlx.DB, query string, args ...interface{}) (*sql.Rows, *Error) {
	rows, err := sdb.Query(query, args...)
	if err != nil {
		return rows, NewError(InvalidOperation, err.Error())
	}
	return rows, nil
}

func ping(sdb *sqlx.DB) *Error {
	err := sdb.Ping()
	if err != nil {
		return NewError(InvalidOperation, err.Error())
	}
	return nil
}
