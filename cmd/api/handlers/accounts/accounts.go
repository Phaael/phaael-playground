package accounts

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phaael/phaael-playground/cmd/api/internal/transactions"
)

type Handler struct {
	TransactionsService transactions.Service
}

type OperationType int64

func (handler *Handler) GetAccountInfo(c *gin.Context) {
	// Validate accountId
	accountID, err := strconv.ParseInt(c.Param("accountId"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid account_id")
		return
	}

	accountInfo, errAccount := handler.TransactionsService.GetAccountInfo(accountID)
	if errAccount != nil {
		c.JSON(errAccount.Status, errAccount)
		return
	}

	c.JSON(http.StatusOK, &accountInfo)
}

func (handler *Handler) CreateAccount(c *gin.Context) {
	var accountData transactions.NewAccount
	bindErr := c.BindJSON(&accountData)

	if bindErr != nil {
		c.String(http.StatusBadRequest, bindErr.Error())
		return
	}

	switch accountData.DocumentNumber.(type) {
	case string:
		docNumber, err := strconv.ParseInt(accountData.DocumentNumber.(string), 10, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid document_number")
			return
		}
		accountData.DocumentNumber = docNumber
	default:
		strDocNumber := fmt.Sprintf("%v", accountData.DocumentNumber)
		docNumber, err := strconv.ParseInt(strDocNumber, 10, 64)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid document_number")
			return
		}
		accountData.DocumentNumber = docNumber

	}

	accountCreated, errAccount := handler.TransactionsService.CreateAccount(accountData)
	if errAccount != nil {
		c.JSON(errAccount.Status, errAccount)
		return
	}

	c.JSON(http.StatusCreated, accountCreated)
}

func (handler *Handler) CreateTransaction(c *gin.Context) {
	var transactionData transactions.NewTransaction
	bindErr := c.BindJSON(&transactionData)

	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": bindErr.Error(),
		})
		return
	}

	ot := transactions.OperationType(transactionData.OperationTypeId)
	if errOtValid := ot.IsValid(); errOtValid != nil {
		c.JSON(http.StatusBadRequest, errOtValid.Error())
		return
	}

	transactionInfo, errTransaction := handler.TransactionsService.CreateTransaction(transactionData)
	if errTransaction != nil {
		c.JSON(errTransaction.Status, errTransaction)
		return
	}

	c.JSON(http.StatusCreated, transactionInfo)
}
