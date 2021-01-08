package transactions

type Service interface {
	GetAccountInfo(accountId int64) (AccountData, error)
}

type Repository interface {
	RetrieveAccountInfo(accountId int64) (AccountData, error)
}
