package transactions

const defaultLimit = 20

type ServiceImpl struct {
	Repo Repository
}

func (s ServiceImpl) GetAccountInfo(accountId int64) (accountInfo AccountData, err error) {
	accountInfo, err = s.Repo.GetAccountInfo(accountId)

	if err != nil {
		return accountInfo, err
	}

	return accountInfo, nil
}

func (s ServiceImpl) CreateAccount(account NewAccount) (accountInfo AccountData, err error) {
	accountInfo, err = s.Repo.CreateAccount(account)

	if err != nil {
		return accountInfo, err
	}

	return accountInfo, nil
}

func (s ServiceImpl) CreateTransaction(transaction NewTransaction) (err error) {
	err = s.Repo.CreateTransaction(transaction)

	if err != nil {
		return err
	}

	return nil
}
