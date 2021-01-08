package accounts

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phaael/phaael-playground/cmd/api/internal/transactions"
)

type Handler struct {
	TransactionsService transactions.Service
}

func (handler *Handler) GetAccountInfo(c *gin.Context) {
	// Validate accountId
	accountID, err := strconv.ParseInt(c.Param("accountId"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid account_id")
		return
	}

	accountInfo, err := handler.TransactionsService.GetAccountInfo(accountID)

	c.JSON(http.StatusOK, accountInfo)
}

func (handler *Handler) CreateAccount(c *gin.Context) {
	var accountData transactions.NewAccount
	bindErr := c.BindJSON(&accountData)

	if bindErr != nil {
		c.String(http.StatusBadRequest, bindErr.Error())
		return
	}

	accountCreated, err := handler.TransactionsService.CreateAccount(accountData)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, accountCreated)
}

func (handler *Handler) CreateTransaction(c *gin.Context) {
	var transactionData transactions.NewTransaction
	bindErr := c.BindJSON(&transactionData)

	if bindErr != nil {
		c.String(http.StatusBadRequest, bindErr.Error())
		return
	}

	err := handler.TransactionsService.CreateTransaction(transactionData)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}
