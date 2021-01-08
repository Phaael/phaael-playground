package mysql

import (
	"database/sql"
)

type Service interface {
	Select(query string, args []interface{}) (*sql.Rows, error)
	Insert(query string, args []interface{}) (sql.Result, error)
}

type Repository interface {
	Init()
	Select(query string, args []interface{}) (*sql.Rows, error)
	Insert(query string, args []interface{}) (sql.Result, error)
}
