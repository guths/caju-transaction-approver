package repository

import (
	"database/sql"

	"github.com/guths/caju-transaction-approver/internal/transaction/domain"
)

type mysqlMccRepository struct {
	db *sql.DB
}

func NewMccRepository(db *sql.DB) domain.MccRepository {
	return &mysqlMccRepository{
		db: db,
	}
}

func (repo *mysqlMccRepository) GetCategoryByMcc(mmc string) domain.Category {
	return domain.Category{}
}
