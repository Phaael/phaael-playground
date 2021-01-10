package transactions

import (
	"fmt"
	"log"

	errStandard "errors"

	"github.com/phaael/phaael-playground/cmd/api/internal/errors"
	"github.com/phaael/phaael-playground/cmd/api/internal/platform/mysql"
)

const SELECT_ACCOUNT_INFO = "SELECT * from accounts where id = ?"
const INSERT_ACCOUNT = "INSERT INTO accounts (`document_number`) VALUES (?)"
const INSERT_TRANSACTION = "INSERT INTO transactions (`account_id`, `operation_type_id`, `amount`, `event_date`) VALUES ( ?, ?, ?, NOW())"

type AccountData struct {
	AccountID      int64  `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

type NewAccount struct {
	DocumentNumber string `json:"document_number" binding:"required"`
}

type NewTransaction struct {
	AccountID       int64   `json:"account_id" binding:"required"`
	OperationTypeId int64   `json:"operation_type_id" binding:"required"`
	Amount          float64 `json:"amount" binding:"required"`
}

type OperationType int64

type RepositoryImpl struct {
	MysqlService mysql.Service
}

type Teste struct {
	Id          int64   `json:"id"`
	Description string  `json:"descriprion"`
	Ammount     float64 `json:"ammount" binding:"required"`
}

func (repo *RepositoryImpl) GetAccountInfo(accountId int64) (accountInfo *AccountData, err *errors.ApiErrorResponse) {

	//  Execute query
	params := []interface{}{}
	params = append(params, accountId)
	rows, dbError := repo.MysqlService.Select(SELECT_ACCOUNT_INFO, params)

	if dbError != nil {
		log.Println(dbError.Error())
		err := errors.GetError(500, dbError.Error())
		return nil, &err
	}

	defer rows.Close()

	if !rows.Next() {
		message := fmt.Sprintf("account with id: %d not found", accountId)
		log.Println(message)
		err := errors.GetError(404, message)
		return nil, &err
	}

	for rows.Next() {
		rows.Scan(
			&accountInfo.AccountID,
			&accountInfo.DocumentNumber,
		)
	}

	return accountInfo, nil

}

func (repo *RepositoryImpl) CreateAccount(account NewAccount) (accountInfo *AccountData, err *errors.ApiErrorResponse) {

	params := []interface{}{}
	params = append(params, account.DocumentNumber)
	result, dbError := repo.MysqlService.Insert(INSERT_ACCOUNT, params)

	if dbError != nil {
		log.Println(dbError.Error())
		err := errors.GetError(500, dbError.Error())
		return nil, &err
	}

	accNumber, _ := result.LastInsertId()
	accountInfo = &AccountData{accNumber, account.DocumentNumber}
	return accountInfo, err
}

func (repo *RepositoryImpl) CreateTransaction(account NewTransaction) (err *errors.ApiErrorResponse) {

	params := []interface{}{}
	params = append(params, account.AccountID)
	params = append(params, account.OperationTypeId)
	params = append(params, account.Amount)

	_, dbError := repo.MysqlService.Insert(INSERT_TRANSACTION, params)

	if dbError != nil {
		log.Println(dbError.Error())
		err := errors.GetError(500, dbError.Error())
		return &err
	}

	return nil
}

func (ot OperationType) IsValid() error {
	switch ot {
	case 1, 2, 3, 4:
		return nil
	}
	return errStandard.New("Invalid operation_type_id")
}
