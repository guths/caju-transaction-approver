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

func (s *BalanceService) GetAmountByAccountId(accountId int, categoryId int) (decimal.Decimal, error) {
	balance, err := s.repo.GetBalanceByAccountId(categoryId, accountId)

	if err != nil {
		return decimal.Zero, err
	}

	return balance.Amount, nil
}

func (s *BalanceService) GetFallbackBalanceAmount(accountId int) (decimal.Decimal, error) {
	balance, err := s.repo.GetFallbackBalance(accountId)

	if err != nil {
		return decimal.Zero, err
	}

	return balance.Amount, nil
}

func (s *BalanceService) IsBalanceSufficient(amountToDebit decimal.Decimal, accountAmount decimal.Decimal) bool {
	return accountAmount.GreaterThan(amountToDebit)
}
