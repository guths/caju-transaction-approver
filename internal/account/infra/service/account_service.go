package service

import "github.com/guths/caju-transaction-approver/internal/account/domain"

type AccountService struct {
	repo domain.AccountRepository
}

func NewAccountService(repo domain.AccountRepository) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) GetAccountByIdentifier(accountIdentifier string) (*domain.Account, error) {
	account, err := s.repo.GetAccountByIdentifier(accountIdentifier)
	if err != nil {
		return nil, err
	}

	return account, nil
}
