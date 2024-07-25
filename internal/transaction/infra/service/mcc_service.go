package service

import (
	"fmt"

	"github.com/guths/caju-transaction-approver/internal/transaction/domain"
)

var (
	ErrCategoryNotFound = fmt.Errorf("category not found")
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
