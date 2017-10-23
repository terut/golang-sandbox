package infrastructure

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/terut/golang-sandbox/cleanarchitecture/interface/database"
)

type SQLHandler struct {
	Conn *sql.DB
}

func NewSQLHandler() database.SQLHandler {
	conn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/cleanarchitecture")
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func (handler *SQLHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SQLResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return nil, err
	}
	res.Result = result
	return res, nil
}

func (handler *SQLHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	var row SQLRow
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return row, err
	}
	row.Rows = rows
	return row, nil
}

type SQLResult struct {
	Result sql.Result
}

func (r SQLResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SQLResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SQLRow struct {
	Rows *sql.Rows
}

func (r SQLRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SQLRow) Next() bool {
	return r.Rows.Next()
}

func (r SQLRow) Close() error {
	return r.Rows.Close()
}
