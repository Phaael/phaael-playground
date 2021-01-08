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
