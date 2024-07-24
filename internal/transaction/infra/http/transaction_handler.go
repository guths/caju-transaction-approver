package http

import (
	"github.com/gin-gonic/gin"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/usecase"
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

}
