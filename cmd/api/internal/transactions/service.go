package transactions

const defaultLimit = 20

type ServiceImpl struct {
	Repo Repository
}

func (s ServiceImpl) GetAccountInfo(accountId int64) (accountInfo AccountData, err error) {
	accountInfo, err = s.Repo.RetrieveAccountInfo(accountId)

	if err != nil {
		return accountInfo, err
	}

	return accountInfo, nil
}
