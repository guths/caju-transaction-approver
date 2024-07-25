package service

import "github.com/guths/caju-transaction-approver/internal/transaction/domain"

type BalanceService struct {
	repo domain.BalanceRepository
}

func NewBalanceService(repo domain.BalanceRepository) BalanceService {
	return BalanceService{
		repo: repo,
	}
}

func (s *BalanceService) GetAmountByAccountId(accountId int, categoryId int) {

}
