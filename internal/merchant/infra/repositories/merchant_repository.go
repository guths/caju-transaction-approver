package repositories

import (
	"database/sql"

	"github.com/guths/caju-transaction-approver/internal/merchant/domain"
	transaction_domain "github.com/guths/caju-transaction-approver/internal/transaction/domain"
)

type mysqlMerchantRepository struct {
	db *sql.DB
}

func NewMysqlMerchantRepository(db *sql.DB) domain.MerchantRepository {
	return &mysqlMerchantRepository{
		db: db,
	}
}

func (repo *mysqlMerchantRepository) GetCategoryByMerchantName(name string) (transaction_domain.Category, error) {
	return transaction_domain.Category{}, nil
}
