package transactions_test

import (
	"testing"

	"github.com/phaael/phaael-playground/cmd/api/internal/errors"
	"github.com/phaael/phaael-playground/cmd/api/internal/transactions"
	"github.com/stretchr/testify/assert"
)

type TransactionRepoImpl struct {
}

func TestGetAccountOk(t *testing.T) {

	accountID := int64(123456)

	transactionRepo := TransactionRepoImpl{}

	transactionServ := transactions.ServiceImpl{
		Repo: &transactionRepo,
	}

	resp, err := transactionServ.GetAccountInfo(accountID)

	assert.Nil(t, err, "Error must be nil")
	assert.NotNil(t, resp, "resp must be not nil")

}

func TestGetAccountError(t *testing.T) {

	accountID := int64(1212)

	transactionRepo := TransactionRepoImpl{}

	transactionServ := transactions.ServiceImpl{
		Repo: &transactionRepo,
	}

	resp, err := transactionServ.GetAccountInfo(accountID)

	assert.Nil(t, resp, "resp must be nil")
	assert.NotNil(t, err, "Error must be not nil")
	assert.Contains(t, err.Message, err.Message, "Internal error")

}

func TestGetAccountNotFound(t *testing.T) {

	accountID := int64(99999)

	transactionRepo := TransactionRepoImpl{}

	transactionServ := transactions.ServiceImpl{
		Repo: &transactionRepo,
	}

	resp, err := transactionServ.GetAccountInfo(accountID)

	assert.Nil(t, resp, "resp must be nil")
	assert.NotNil(t, err, "Error must be not nil")
	assert.Contains(t, err.Message, err.Message, "account with id: 99999 not found")

}

func TestCreateTransactionOkWithOperationType1(t *testing.T) {

	transaction := transactions.NewTransaction{AccountID: 1, OperationTypeId: 1, Amount: 10}

	transactionRepo := TransactionRepoImpl{}

	transactionServ := transactions.ServiceImpl{
		Repo: &transactionRepo,
	}

	resp, err := transactionServ.CreateTransaction(transaction)

	assert.Nil(t, err, "Error must be nil")
	assert.NotNil(t, resp, "resp must be not nil")
	assert.Equal(t, float64(-10), resp.Amount)

}

func TestCreateTransactionOkWithOperationType4(t *testing.T) {

	transaction := transactions.NewTransaction{AccountID: 1, OperationTypeId: 4, Amount: 10}

	transactionRepo := TransactionRepoImpl{}

	transactionServ := transactions.ServiceImpl{
		Repo: &transactionRepo,
	}

	resp, err := transactionServ.CreateTransaction(transaction)

	assert.Nil(t, err, "Error must be nil")
	assert.NotNil(t, resp, "resp must be not nil")
	assert.Equal(t, float64(10), resp.Amount)

}

func TestCreateTransactionError(t *testing.T) {

	transaction := transactions.NewTransaction{AccountID: 1212, OperationTypeId: 4, Amount: 10}

	transactionRepo := TransactionRepoImpl{}

	transactionServ := transactions.ServiceImpl{
		Repo: &transactionRepo,
	}

	resp, err := transactionServ.CreateTransaction(transaction)

	assert.Nil(t, resp, "resp must be nil")
	assert.NotNil(t, err, "err must be not nil")
	assert.Contains(t, err.Message, err.Message, "Internal error")

}

func TestCreateTransactionWhenAccountDoesNotExists(t *testing.T) {

	transaction := transactions.NewTransaction{AccountID: 99999, OperationTypeId: 4, Amount: 10}

	transactionRepo := TransactionRepoImpl{}

	transactionServ := transactions.ServiceImpl{
		Repo: &transactionRepo,
	}

	resp, err := transactionServ.CreateTransaction(transaction)

	assert.Nil(t, resp, "resp must be nil")
	assert.NotNil(t, err, "err must be not nil")
	assert.Contains(t, err.Message, err.Message, "account with id: 99999 not found")

}

func TestCreateAccountError(t *testing.T) {

	account := transactions.NewAccount{DocumentNumber: "1212"}

	transactionRepo := TransactionRepoImpl{}

	transactionServ := transactions.ServiceImpl{
		Repo: &transactionRepo,
	}

	resp, err := transactionServ.CreateAccount(account)

	assert.Nil(t, resp, "resp must be nil")
	assert.NotNil(t, err, "err must be not nil")
	assert.Contains(t, err.Message, err.Message, "Internal error")

}

func TestCreateAccountOk(t *testing.T) {

	account := transactions.NewAccount{DocumentNumber: "1234"}

	transactionRepo := TransactionRepoImpl{}

	transactionServ := transactions.ServiceImpl{
		Repo: &transactionRepo,
	}

	resp, err := transactionServ.CreateAccount(account)

	assert.Nil(t, err, "err must be nil")
	assert.NotNil(t, resp, "resp must be not nil")
}

/** mock repository USED in service **/
func (a *TransactionRepoImpl) GetAccountInfo(accountId int64) (*transactions.AccountData, *errors.ApiErrorResponse) {
	if accountId == 99999 {
		err := errors.GetError(404, "account with id: 99999 not found")
		return nil, &err
	}

	if accountId == 1212 {
		err := errors.GetError(500, "Internal error")
		return nil, &err
	}

	accountInfo := transactions.AccountData{AccountID: 1223456, DocumentNumber: int64(1111)}
	return &accountInfo, nil

}

func (a *TransactionRepoImpl) CreateAccount(account transactions.NewAccount) (*transactions.AccountData, *errors.ApiErrorResponse) {
	if account.DocumentNumber == "1212" {
		err := errors.GetError(500, "Internal error")
		return nil, &err
	}

	accountInfo := transactions.AccountData{AccountID: 1223456, DocumentNumber: int64(1111)}
	return &accountInfo, nil
}

func (a *TransactionRepoImpl) CreateTransaction(transaction transactions.NewTransaction) (*transactions.TransactionInfo, *errors.ApiErrorResponse) {
	if transaction.AccountID == 1212 {
		err := errors.GetError(500, "Internal error")
		return nil, &err
	}

	if transaction.AccountID == 99999 {
		err := errors.GetError(404, "account with id: 99999 not found")
		return nil, &err
	}

	tx := transactions.TransactionInfo{TransactionId: 1, Amount: transaction.Amount, OperationTypeId: transaction.OperationTypeId, AccountID: transaction.AccountID}

	return &tx, nil
}
