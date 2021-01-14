package transactions

import "github.com/phaael/phaael-playground/cmd/api/internal/errors"

type Service interface {
	GetAccountInfo(accountId int64) (*AccountData, *errors.ApiErrorResponse)
	CreateAccount(account NewAccount) (*AccountData, *errors.ApiErrorResponse)
	CreateTransaction(transaction NewTransaction) (*TransactionInfo, *errors.ApiErrorResponse)
}

type Repository interface {
	GetAccountInfo(accountId int64) (*AccountData, *errors.ApiErrorResponse)
	CreateAccount(account NewAccount) (*AccountData, *errors.ApiErrorResponse)
	CreateTransaction(transaction NewTransaction) (*TransactionInfo, *errors.ApiErrorResponse)
}
