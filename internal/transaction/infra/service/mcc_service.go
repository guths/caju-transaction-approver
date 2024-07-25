package service

import (
	"github.com/guths/caju-transaction-approver/internal/transaction/domain"
)

type MccService struct {
	mccRepo domain.MccRepository
}

func NewMccService(mccRepo domain.MccRepository) MccService {
	return MccService{
		mccRepo: mccRepo,
	}
}

func (s *MccService) GetCategoryByMcc(mcc string) (domain.Category, error) {
	category, err := s.mccRepo.GetCategoryByMcc(mcc)

	if err != nil {
		return domain.Category{}, err
	}

	return category, nil
}

func (s *MccService) GetFallbackCategory() (*domain.Category, error) {
	category, err := s.mccRepo.GetFallbackCategory()

	if err != nil {
		return nil, err
	}

	return category, nil
}
