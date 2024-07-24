package repository

import (
	"database/sql"

	"github.com/guths/caju-transaction-approver/internal/transaction/domain"
)

type mysqlBalanceRepository struct {
	db *sql.DB
}

func NewBalanceRepository(db *sql.DB) domain.BalanceRepository {
	return &mysqlBalanceRepository{
		db: db,
	}
}
