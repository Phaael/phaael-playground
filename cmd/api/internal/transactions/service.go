package transactions

import (
	"github.com/phaael/phaael-playground/cmd/api/internal/errors"
)

const defaultLimit = 20

type ServiceImpl struct {
	Repo Repository
}

func (s ServiceImpl) GetAccountInfo(accountId int64) (accountInfo *AccountData, err *errors.ApiErrorResponse) {
	accountInfo, err = s.Repo.GetAccountInfo(accountId)

	if err != nil {
		return accountInfo, err
	}

	return accountInfo, nil
}

func (s ServiceImpl) CreateAccount(account NewAccount) (accountInfo *AccountData, err *errors.ApiErrorResponse) {
	accountInfo, err = s.Repo.CreateAccount(account)

	if err != nil {
		return accountInfo, err
	}

	return accountInfo, nil
}

func (s ServiceImpl) CreateTransaction(transaction NewTransaction) (txInfo *TransactionInfo, err *errors.ApiErrorResponse) {
	_, errAccountInvalid := s.Repo.GetAccountInfo(transaction.AccountID)
	if errAccountInvalid != nil {
		return nil, errAccountInvalid
	}

	switch transaction.OperationTypeId {
	case 1, 2, 3:
		transaction.Amount = transaction.Amount - transaction.Amount*2
	}

	txInfo, errCreated := s.Repo.CreateTransaction(transaction)

	if errCreated != nil {
		return nil, errCreated
	}

	return txInfo, nil
}
