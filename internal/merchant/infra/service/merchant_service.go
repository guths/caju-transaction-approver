package service

import (
	"github.com/guths/caju-transaction-approver/internal/merchant/domain"
	transaction_domain "github.com/guths/caju-transaction-approver/internal/transaction/domain"
)

type MerchantService struct {
	repo domain.MerchantRepository
}

func NewMerchantService(repo domain.MerchantRepository) MerchantService {
	return MerchantService{
		repo: repo,
	}
}

func (service *MerchantService) GetCategoryByMerchantName(name string) (*transaction_domain.Category, error) {
	category, err := service.repo.GetCategoryByMerchantName(name)

	if err != nil {
		return nil, err
	}
	return category, nil
}
