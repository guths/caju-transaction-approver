package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guths/caju-transaction-approver/internal/transaction/domain"
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
	var inputTransactionDTO usecase.InputTransactionDTO

	if err := c.ShouldBindBodyWithJSON(&inputTransactionDTO); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": domain.GetGenericResponseError("invalid request format")})
	}

	err := h.authorizeTransaction.Execute(inputTransactionDTO)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": domain.GetGenericResponseError("invalid request format")})
	}

	c.JSON(http.StatusOK, gin.H{
		"success": domain.GetApprovedResponse(),
	})
}
