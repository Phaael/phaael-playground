package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/phaael/phaael-playground/cmd/api/internal/platform/config"
)

var (
	DB      *sql.DB
	DbError error
)

//RepositoryImpl default implementation of the mysql repository
type RepositoryImpl struct {
}

func dbOpenConnection(cnnString string) {

	DB, DbError = sql.Open("mysql", cnnString)

	if DbError != nil {
		log.Panicf("Cannot connect to DB: '%s(MISSING)'", DbError)
	}

	if DbError = DB.Ping(); DbError != nil {
		log.Panicf("Cannot ping the DB connection. %v(MISSING)", DbError)
	}

	DB.SetMaxIdleConns(50)

}

func getDatabasePath(DbUser, DbPass, DbHost, DbName string) string {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", DbUser, DbPass, DbHost, DbName)
	return connectionString
}

func (repo *RepositoryImpl) Init() {
	cnn := getDatabasePath(config.DbUserName, config.PassDB, config.DBHost, config.DataBaseName)
	dbOpenConnection(cnn)
}

// need to receive the params query = ".... with prepared statement", vals := []interface{}{1, 2}
func (repo *RepositoryImpl) Select(query string, args []interface{}) (*sql.Rows, error) {
	rows, err := DB.Query(query, args...)
	return rows, err
}

// need to receive the params query = ".... with prepared statement", vals := []interface{}{1, 2}
func (repo *RepositoryImpl) Insert(query string, args []interface{}) (sql.Result, error) {
	result, err := DB.Exec(query, args...)
	return result, err
}
