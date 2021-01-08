package mysql

import (
	"database/sql"
)

type MockImpl struct {
}

func (repo *MockImpl) Init() {
}

func (repo *MockImpl) Select(query string, args []interface{}) (*sql.Rows, error) {
	return nil, nil
}

func (repo *MockImpl) Insert(query string, args []interface{}) (sql.Result, error) {
	return nil, nil
}
