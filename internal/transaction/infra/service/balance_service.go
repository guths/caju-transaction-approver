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

func (s *BalanceService) GetAmountByAccountId(accountId int, categoryId int) (decimal.Decimal, error) {
	balance, err := s.repo.GetBalanceByAccountId(categoryId, accountId)

	if err != nil {
		return decimal.Zero, err
	}

	return balance.Amount, nil
}

func (s *BalanceService) IsBalanceSufficient(amountToDebit decimal.Decimal, accountAmount decimal.Decimal) bool {
	return accountAmount.GreaterThan(amountToDebit)
}
