package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/guths/caju-transaction-approver/configs"
	account_repo "github.com/guths/caju-transaction-approver/internal/account/infra/repository"
	account_service "github.com/guths/caju-transaction-approver/internal/account/infra/service"
	merchant_repo "github.com/guths/caju-transaction-approver/internal/merchant/infra/repository"
	merchant_service "github.com/guths/caju-transaction-approver/internal/merchant/infra/service"
	http_handler "github.com/guths/caju-transaction-approver/internal/transaction/infra/http"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/repository"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/service"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/usecase"
	"github.com/spf13/cobra"
)

var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Command to start the api",
	Run: func(cmd *cobra.Command, args []string) {
		server := GetServer()
		err := server.Run()

		if err != nil {
			panic(err)
		}

	},
}

func GetServer() *gin.Engine {
	server := gin.New()

	mysqlBalanceRepository := repository.NewBalanceRepository(configs.DB)
	balanceService := service.NewBalanceService(mysqlBalanceRepository)
	merchantRepo := merchant_repo.NewMysqlMerchantRepository(configs.DB)
	merchantService := merchant_service.NewMerchantService(merchantRepo)
	mccRepo := repository.NewMccRepository(configs.DB)
	mccService := service.NewMccService(mccRepo)
	accountRepo := account_repo.NewMysqlAccountRepository(configs.DB)
	accountService := account_service.NewAccountService(accountRepo)
	authorizeTransactionUseCase := usecase.NewAuthorizeTransactionUseCase(balanceService, merchantService, mccService, *accountService)
	transactionHandler := http_handler.NewTransactionHandler(&authorizeTransactionUseCase)

	server.POST("/authorize-transaction", transactionHandler.AuthorizeTransaction)

	return server
}
