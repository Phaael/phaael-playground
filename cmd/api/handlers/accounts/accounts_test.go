package accounts_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/phaael/phaael-playground/cmd/api/handlers/accounts"
	"github.com/phaael/phaael-playground/cmd/api/internal/errors"
	"github.com/phaael/phaael-playground/cmd/api/internal/transactions"
	"github.com/stretchr/testify/assert"
)

func TestGetAccountWithAccountIdInvalid(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := accounts.Handler{}

	router.GET("/accounts/:accountId", handler.GetAccountInfo)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/accounts/invalid", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusBadRequest, response.Code, "Response status must be 400")
}

func TestGetAccountWithAccountIdOK(t *testing.T) {
	transactionsServ := TransactionsServiceMock{}

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := accounts.Handler{
		TransactionsService: &transactionsServ,
	}

	router.GET("/accounts/:accountId", handler.GetAccountInfo)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/accounts/12345", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code, "Response status must be 200")
}

func TestGetAccountWithAccountIdOKButErronOnCallServ(t *testing.T) {
	transactionsServ := TransactionsServiceMock{}

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := accounts.Handler{
		TransactionsService: &transactionsServ,
	}

	router.GET("/accounts/:accountId", handler.GetAccountInfo)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/accounts/1212", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusInternalServerError, response.Code, "Response status must be 500")
}

func TestCreateAccountWithBodyInvalid(t *testing.T) {

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := accounts.Handler{}

	router.POST("/accounts", handler.CreateAccount)

	response := httptest.NewRecorder()
	request, _ := http.NewRequestWithContext(context.Background(), "POST",
		"/accounts",
		strings.NewReader(`{
			"invalid": false
		}`))
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusBadRequest, response.Code, "Response status must be 400")
}

func TestCreateAccountWithErrorOnServ(t *testing.T) {
	transactionsServ := TransactionsServiceMock{}

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := accounts.Handler{
		TransactionsService: &transactionsServ,
	}

	router.POST("/accounts", handler.CreateAccount)

	response := httptest.NewRecorder()
	request, _ := http.NewRequestWithContext(context.Background(), "POST",
		"/accounts",
		strings.NewReader(`{
			"document_number": "1212"
		}`))
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusInternalServerError, response.Code, "Response status must be 500")
}

func TestCreateAccountOK(t *testing.T) {
	transactionsServ := TransactionsServiceMock{}

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := accounts.Handler{
		TransactionsService: &transactionsServ,
	}

	router.POST("/accounts", handler.CreateAccount)

	response := httptest.NewRecorder()
	request, _ := http.NewRequestWithContext(context.Background(), "POST",
		"/accounts",
		strings.NewReader(`{
			"document_number": "1111"
		}`))
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusCreated, response.Code, "Response status must be 201")
}

func TestCreateTransactionWithBodyInvalid(t *testing.T) {

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := accounts.Handler{}

	router.POST("/transactions", handler.CreateTransaction)

	response := httptest.NewRecorder()
	request, _ := http.NewRequestWithContext(context.Background(), "POST",
		"/transactions",
		strings.NewReader(`{
			"invalid": false
		}`))
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusBadRequest, response.Code, "Response status must be 400")
}

func TestCreateTransactionWithTransactionTypeInvalid(t *testing.T) {

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := accounts.Handler{}

	router.POST("/transactions", handler.CreateTransaction)

	response := httptest.NewRecorder()
	request, _ := http.NewRequestWithContext(context.Background(), "POST",
		"/transactions",
		strings.NewReader(`{
			"account_id": 1,
			"operation_type_id": 5,
			"amount": 123.45
		}`))
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusBadRequest, response.Code, "Response status must be 400")
	assert.Contains(t, response.Body.String(), "Invalid operation_type_id", "Response must be Invalid operation_type_id")

}

func TestCreateTransactionWithTransactionWithAccountDoesNotExists(t *testing.T) {
	transactionsServ := TransactionsServiceMock{}

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := accounts.Handler{
		TransactionsService: &transactionsServ,
	}

	router.POST("/transactions", handler.CreateTransaction)

	response := httptest.NewRecorder()
	request, _ := http.NewRequestWithContext(context.Background(), "POST",
		"/transactions",
		strings.NewReader(`{
			"account_id": 99999,
			"operation_type_id": 4,
			"amount": 123.45
		}`))
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusNotFound, response.Code, "Response status must be 404")
	assert.Contains(t, response.Body.String(), "account with id: 99999 not found", "Response must be account with id: 99999 not found")

}

func TestCreateTransactionWithTransactionWithErronOnServer(t *testing.T) {
	transactionsServ := TransactionsServiceMock{}

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := accounts.Handler{
		TransactionsService: &transactionsServ,
	}

	router.POST("/transactions", handler.CreateTransaction)

	response := httptest.NewRecorder()
	request, _ := http.NewRequestWithContext(context.Background(), "POST",
		"/transactions",
		strings.NewReader(`{
			"account_id": 1212,
			"operation_type_id": 4,
			"amount": 123.45
		}`))
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusInternalServerError, response.Code, "Response status must be 500")
	assert.Contains(t, response.Body.String(), "Internal error", "Response must be Internal error")

}

func TestCreateTransactionOk(t *testing.T) {
	transactionsServ := TransactionsServiceMock{}

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	handler := accounts.Handler{
		TransactionsService: &transactionsServ,
	}

	router.POST("/transactions", handler.CreateTransaction)

	response := httptest.NewRecorder()
	request, _ := http.NewRequestWithContext(context.Background(), "POST",
		"/transactions",
		strings.NewReader(`{
			"account_id": 1111,
			"operation_type_id": 4,
			"amount": 123.45
		}`))
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusCreated, response.Code, "Response status must be 201")
	assert.Contains(t, response.Body.String(), "created", "Response must be created")

}

type TransactionsServiceMock struct{}

func (a *TransactionsServiceMock) GetAccountInfo(accountId int64) (*transactions.AccountData, *errors.ApiErrorResponse) {
	if accountId == 1212 {
		err := errors.GetError(500, "Internal error")
		return nil, &err
	}

	accountInfo := transactions.AccountData{AccountID: 1223456, DocumentNumber: "XPTO"}
	return &accountInfo, nil
}

func (a *TransactionsServiceMock) CreateAccount(account transactions.NewAccount) (*transactions.AccountData, *errors.ApiErrorResponse) {
	if account.DocumentNumber == "1212" {
		err := errors.GetError(500, "Internal error")
		return nil, &err
	}

	accountInfo := transactions.AccountData{AccountID: 1223456, DocumentNumber: "1212"}
	return &accountInfo, nil
}

func (a *TransactionsServiceMock) CreateTransaction(transaction transactions.NewTransaction) *errors.ApiErrorResponse {
	if transaction.AccountID == 1212 {
		err := errors.GetError(500, "Internal error")
		return &err
	}

	if transaction.AccountID == 99999 {
		err := errors.GetError(404, "account with id: 99999 not found")
		return &err
	}

	return nil
}
