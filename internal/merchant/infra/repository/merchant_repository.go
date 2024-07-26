package repository

import (
	"database/sql"
	"fmt"

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

var (
	ErrCategoryNotFound = fmt.Errorf("category not found")
)

func (repo *mysqlMerchantRepository) GetCategoryByMerchantName(name string) (*transaction_domain.Category, error) {
	var category transaction_domain.Category

	q := `
		SELECT category.id, category.name
		FROM merchant
		INNER JOIN category
		ON merchant.category_id = category.id
		WHERE merchant.name = ?
	`

	err := repo.db.QueryRow(q, name).Scan(
		&category.Id,
		&category.Name,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrCategoryNotFound
		}

		return nil, err
	}

	return &category, nil
}
