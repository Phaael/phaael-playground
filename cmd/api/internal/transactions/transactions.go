package transactions

type Service interface {
	GetAccountInfo(accountId int64) (AccountData, error)
	CreateAccount(account NewAccount) (AccountData, error)
	CreateTransaction(transaction NewTransaction) error
}

type Repository interface {
	GetAccountInfo(accountId int64) (AccountData, error)
	CreateAccount(account NewAccount) (AccountData, error)
	CreateTransaction(transaction NewTransaction) error
}
