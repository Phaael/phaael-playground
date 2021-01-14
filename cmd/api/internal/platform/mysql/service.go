package mysql

import "database/sql"

//ServiceImpl default implementation of the mysql service
type ServiceImpl struct {
	Repository Repository
}

//Select data into mysql
func (s *ServiceImpl) Select(query string, args []interface{}) (*sql.Rows, error) {
	return s.Repository.Select(query, args)
}

func (s *ServiceImpl) Insert(query string, args []interface{}) (sql.Result, error) {
	return s.Repository.Insert(query, args)
}

func (s *ServiceImpl) Delete(query string, args []interface{}) (sql.Result, error) {
	return s.Repository.Delete(query, args)
}
