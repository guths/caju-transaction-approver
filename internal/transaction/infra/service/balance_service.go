package service

import (
	"github.com/guths/caju-transaction-approver/internal/transaction/domain"
	"github.com/shopspring/decimal"
)

type BalanceService struct {
	repo domain.BalanceRepository
}

func NewBalanceService(repo domain.BalanceRepository) BalanceService {
	return BalanceService{
		repo: repo,
	}
}

func (s *BalanceService) DebitAmount(accountId int, categoryId int, amount decimal.Decimal) (*domain.Transaction, error) {
	transaction, err := s.repo.DebitAmount(accountId, categoryId, amount)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}
