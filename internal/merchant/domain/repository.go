package domain

import "github.com/guths/caju-transaction-approver/internal/transaction/domain"

type MerchantRepository interface {
	GetCategoryByMerchantName(name string) (domain.Category, error)
}
