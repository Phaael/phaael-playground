package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/phaael/phaael-playground/cmd/api/handlers/accounts"
	"github.com/phaael/phaael-playground/cmd/api/internal/transactions"
)

func MapURL(router *gin.Engine) {

	// Add health check
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	addAccountsRoute(router)
}

func addAccountsRoute(router *gin.Engine) {
	// mysqlImpl := mysql.RepositoryImpl{}
	// mysqlImpl.Init()

	// mysqlServ := mysql.ServiceImpl{
	// 	Repository: &mysqlImpl,
	// }

	transactionsImpl := transactions.RepositoryImpl{
		//MysqlService: &mysqlServ,
	}

	transactionsServ := transactions.ServiceImpl{
		Repo: &transactionsImpl,
	}

	accountsHandler := accounts.Handler{
		TransactionsService: &transactionsServ,
	}

	router.GET("/accounts/:accountId", accountsHandler.GetAccountInfo)
	router.POST("/accounts", accountsHandler.CreateAccount)
	router.POST("/transactions", accountsHandler.CreateTransaction)

}
