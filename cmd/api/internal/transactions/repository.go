package transactions

import (
	"github.com/phaael/phaael-playground/cmd/api/internal/platform/mysql"
)

type AccountData struct {
	AccountID      int64  `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

type RepositoryImpl struct {
	MysqlService mysql.Service
}

func (repo *RepositoryImpl) RetrieveAccountInfo(accountId int64) (accountInfo AccountData, err error) {

	// Execute query
	// rows, err := repo.MysqlService.Select(query, params)

	// just test
	accountInfo = AccountData{AccountID: 1223456, DocumentNumber: "XPTO"}

	return accountInfo, err
}
