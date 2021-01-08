package transactions

import (
	"github.com/phaael/phaael-playground/cmd/api/internal/platform/mysql"
)

type AccountData struct {
	AccountID      int64  `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

type NewAccount struct {
	DocumentNumber string `json:"document_number" binding:"required"`
}

type NewTransaction struct {
	AccountID       int64   `json:"account_id" binding:"required"`
	OperationTypeId string  `json:"operation_type_id" binding:"required"`
	Ammount         float64 `json:"ammount" binding:"required"`
}

type RepositoryImpl struct {
	MysqlService mysql.Service
}

func (repo *RepositoryImpl) GetAccountInfo(accountId int64) (accountInfo AccountData, err error) {

	//  Execute query
	// rows, err := repo.MysqlService.Select(query, params)

	// just test
	accountInfo = AccountData{AccountID: 1223456, DocumentNumber: "XPTO"}

	return accountInfo, err
}

func (repo *RepositoryImpl) CreateAccount(Account NewAccount) (accountInfo AccountData, err error) {

	//  Execute query
	// rows, err := repo.MysqlService.Insert(query, params)

	// just test
	accountInfo = AccountData{AccountID: 1223456, DocumentNumber: "XPTO"}

	return accountInfo, err
}

func (repo *RepositoryImpl) CreateTransaction(Account NewTransaction) (err error) {

	//  Execute query
	// rows, err := repo.MysqlService.Insert(query, params)

	// just test

	return nil
}
