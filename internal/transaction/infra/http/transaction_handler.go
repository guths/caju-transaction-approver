package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guths/caju-transaction-approver/internal/transaction/domain"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/usecase"
	"github.com/guths/caju-transaction-approver/internal/validator"
)

type TransactionHandler struct {
	authorizeTransaction *usecase.AuthorizeTransactionUseCase
}

func NewTransactionHandler(
	authorizeTransaction *usecase.AuthorizeTransactionUseCase,
) *TransactionHandler {
	return &TransactionHandler{
		authorizeTransaction: authorizeTransaction,
	}
}

func (h *TransactionHandler) AuthorizeTransaction(c *gin.Context) {
	var inputTransactionDTO usecase.InputTransactionDTO

	if err := c.ShouldBindBodyWithJSON(&inputTransactionDTO); err != nil {
		c.JSON(http.StatusOK, domain.GetGenericResponseError("invalid request format"))
		return
	}

	v := validator.New()

	if usecase.ValidateInputTransaction(v, inputTransactionDTO); !v.Valid() {
		fmt.Println("errors", v.Errors)
		c.JSON(http.StatusOK, domain.GetGenericResponseError("invalid request format"))
		return
	}

	output := h.authorizeTransaction.Execute(inputTransactionDTO)

	c.JSON(http.StatusOK, output)
}
