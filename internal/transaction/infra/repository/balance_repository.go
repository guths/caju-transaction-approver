package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/guths/caju-transaction-approver/internal/transaction/domain"
)

var (
	ErrBalanceNotFound = fmt.Errorf("balance not found in db")
)

type mysqlBalanceRepository struct {
	db *sql.DB
}

func NewBalanceRepository(db *sql.DB) domain.BalanceRepository {
	return &mysqlBalanceRepository{
		db: db,
	}
}

func (repo *mysqlBalanceRepository) GetBalanceByAccountId(categoryId int, accountId int) (domain.Balance, error) {
	tx, err := repo.db.Begin()

	if err != nil {
		return domain.Balance{}, err
	}

	defer tx.Rollback()

	var balance domain.Balance

	q := `
		SELECT id, amount
		FROM balance
		WHERE account_id = $1
		AND category_id = $2
	`

	args := []interface{}{
		accountId,
		categoryId,
	}

	err = tx.QueryRow(q, args...).Scan(
		&balance.Id,
		&balance.Amount,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return domain.Balance{}, ErrBalanceNotFound
		}
	}

	if err := tx.Commit(); err != nil {
		return domain.Balance{}, err
	}

	return balance, nil
}
