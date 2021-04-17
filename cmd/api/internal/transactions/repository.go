package transactions

import (
	"fmt"
	"log"

	errStandard "errors"

	"github.com/phaael/phaael-playground/cmd/api/internal/errors"
	"github.com/phaael/phaael-playground/cmd/api/internal/platform/mysql"
)

const SELECT_ACCOUNT_INFO = "SELECT id, document_number, available_credit_limit from accounts where id = ?"
const INSERT_ACCOUNT = "INSERT INTO accounts (`document_number`, `available_credit_limit`) VALUES (?,?)"
const INSERT_TRANSACTION = "INSERT INTO transactions (`account_id`, `operation_type_id`, `amount`, `event_date`) VALUES ( ?, ?, ?, NOW())"
const UPDATE_ACCOUNT = "UPDATE accounts SET available_credit_limit = ? where id = ?"

const AVALIABLE_CREDIT_LIMIT_DEFAULT = 1000.00

type AccountData struct {
	AccountID            int64   `json:"account_id"`
	DocumentNumber       int64   `json:"document_number"`
	AvailableCreditLimit float64 `json:"available_credit_limit"`
}

type NewAccount struct {
	DocumentNumber       interface{} `json:"document_number" binding:"required"`
	AvailableCreditLimit float64     `json:"available_credit_limit"`
}

type NewTransaction struct {
	AccountID       int64   `json:"account_id" binding:"required"`
	OperationTypeId int64   `json:"operation_type_id" binding:"required"`
	Amount          float64 `json:"amount" binding:"required"`
}

type TransactionInfo struct {
	TransactionId   int64   `json:"transaction_id"`
	AccountID       int64   `json:"account_id"`
	OperationTypeId int64   `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

type OperationType int64

type RepositoryImpl struct {
	MysqlService mysql.Service
}

func (repo *RepositoryImpl) GetAccountInfo(accountId int64) (*AccountData, *errors.ApiErrorResponse) {

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

	var accountInfo = AccountData{}
	// for rows.Next() {
	rows.Scan(
		&accountInfo.AccountID,
		&accountInfo.DocumentNumber,
		&accountInfo.AvailableCreditLimit,
	)
	// }

	return &accountInfo, nil

}

func (repo *RepositoryImpl) CreateAccount(account NewAccount) (accountInfo *AccountData, err *errors.ApiErrorResponse) {

	params := []interface{}{}
	params = append(params, account.DocumentNumber)
	params = append(params, account.AvailableCreditLimit)

	result, dbError := repo.MysqlService.Insert(INSERT_ACCOUNT, params)

	if dbError != nil {
		log.Println(dbError.Error())
		err := errors.GetError(500, dbError.Error())
		return nil, &err
	}

	accNumber, _ := result.LastInsertId()
	accountInfo = &AccountData{accNumber, account.DocumentNumber.(int64), account.AvailableCreditLimit}
	return accountInfo, err
}

func (repo *RepositoryImpl) CreateTransaction(account NewTransaction) (transaction *TransactionInfo, err *errors.ApiErrorResponse) {

	params := []interface{}{}
	params = append(params, account.AccountID)
	params = append(params, account.OperationTypeId)
	params = append(params, account.Amount)

	result, dbError := repo.MysqlService.Insert(INSERT_TRANSACTION, params)

	if dbError != nil {
		log.Println(dbError.Error())
		err := errors.GetError(500, dbError.Error())
		return nil, &err
	}

	txNumber, _ := result.LastInsertId()
	transactionInfo := &TransactionInfo{AccountID: account.AccountID, OperationTypeId: account.OperationTypeId, Amount: account.Amount, TransactionId: txNumber}

	return transactionInfo, nil
}

func (ot OperationType) IsValid() error {
	switch ot {
	case 1, 2, 3, 4:
		return nil
	}
	return errStandard.New("Invalid operation_type_id")
}

func (repo *RepositoryImpl) UpdateAccountInfo(newLimit float64, accountId int64) (err *errors.ApiErrorResponse) {

	params := []interface{}{}
	params = append(params, newLimit)
	params = append(params, accountId)

	_, dbError := repo.MysqlService.Update(UPDATE_ACCOUNT, params)

	if dbError != nil {
		log.Println(dbError.Error())
		err := errors.GetError(500, dbError.Error())
		return &err
	}

	return nil
}
