package usecase

import "github.com/guths/caju-transaction-approver/internal/transaction/infra/service"

type AuthorizeTransactionUseCase struct {
	balanceService service.BalanceService
}

func NewAuthorizeTransactionUseCase(balanceService service.BalanceService) AuthorizeTransactionUseCase {
	return AuthorizeTransactionUseCase{
		balanceService: balanceService,
	}
}
