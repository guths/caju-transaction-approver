package factory

import (
	"database/sql"

	"github.com/guths/caju-transaction-approver/internal/transaction/domain"
	"github.com/shopspring/decimal"
)

type BalanceFactory struct {
	db *sql.DB
}

func NewBalanceFactory(db *sql.DB) *BalanceFactory {
	return &BalanceFactory{db: db}
}

func (f *BalanceFactory) CreateBalance(accountId int, categoryId int, amount decimal.Decimal) (*domain.Balance, error) {

	q := ` INSERT INTO balance (account_id, category_id, amount) VALUES (?, ?, ?)`

	result, err := f.db.Exec(q, accountId, categoryId, amount)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	q = `SELECT id, amount, account_id, category_id FROM balance WHERE id = ?`

	var balance domain.Balance

	err = f.db.QueryRow(q, id).Scan(&balance.Id, &balance.Amount, &balance.AccountId, &balance.CategoryId)

	if err != nil {
		return nil, err
	}

	return &balance, nil
}
