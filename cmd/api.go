package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/guths/caju-transaction-approver/configs"
	http_handler "github.com/guths/caju-transaction-approver/internal/transaction/infra/http"
	balance_repository "github.com/guths/caju-transaction-approver/internal/transaction/infra/repository"
	balance_service "github.com/guths/caju-transaction-approver/internal/transaction/infra/service"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/usecase"
	"github.com/spf13/cobra"
)

var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Command to start the api",
	Run: func(cmd *cobra.Command, args []string) {
		server := gin.New()

		mysqlBalanceRepository := balance_repository.NewBalanceRepository(configs.DB)
		balanceService := balance_service.NewBalanceService(mysqlBalanceRepository)
		authorizeTransactionUseCase := usecase.NewAuthorizeTransactionUseCase(balanceService)
		transactionHandler := http_handler.NewTransactionHandler(&authorizeTransactionUseCase)

		server.POST("/authorize-transaction", transactionHandler.AuthorizeTransaction)

		server.Run()
	},
}
