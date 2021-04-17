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
	accountInfo, errAccountInvalid := s.Repo.GetAccountInfo(transaction.AccountID)
	if errAccountInvalid != nil {
		return nil, errAccountInvalid
	}

	var newLimit float64
	switch transaction.OperationTypeId {
	case 1, 2, 3:
		if transaction.Amount > accountInfo.AvailableCreditLimit {
			err := errors.GetError(403, "transaction not allowed")
			return nil, &err
		}

		transaction.Amount = transaction.Amount * -1
		newLimit = accountInfo.AvailableCreditLimit + transaction.Amount
	default:
		newLimit = accountInfo.AvailableCreditLimit + transaction.Amount
	}

	txInfo, errCreated := s.Repo.CreateTransaction(transaction)
	if errCreated != nil {
		return nil, errCreated
	}

	errUpdate := s.UpdateAccountInfo(newLimit, transaction.AccountID)
	if errUpdate != nil {
		return nil, errUpdate
	}

	return txInfo, nil
}

func (s ServiceImpl) UpdateAccountInfo(newLimit float64, accountId int64) (err *errors.ApiErrorResponse) {
	err = s.Repo.UpdateAccountInfo(newLimit, accountId)

	if err != nil {
		return err
	}

	return nil
}
