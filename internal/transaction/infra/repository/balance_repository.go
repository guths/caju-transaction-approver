package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/guths/caju-transaction-approver/internal/transaction/domain"
	"github.com/shopspring/decimal"
)

var (
	ErrBalanceNotFound     = fmt.Errorf("balance not found in db")
	ErrInsufficientFunds   = fmt.Errorf("insufficient funds")
	ErrUpdatingBalance     = fmt.Errorf("balance can not be updated")
	ErrCreatingTransaction = fmt.Errorf("transaction can not be created")
	ErrTransactionNotFound = fmt.Errorf("transaction not found")
)

type mysqlBalanceRepository struct {
	db *sql.DB
}

func NewBalanceRepository(db *sql.DB) domain.BalanceRepository {
	return &mysqlBalanceRepository{
		db: db,
	}
}

func (repo *mysqlBalanceRepository) DebitAmount(accountId int, categoryId int, amount decimal.Decimal) (*domain.Transaction, error) {
	tx, err := repo.db.Begin()

	if err != nil {
		return nil, err
	}

	defer func() {
		err = tx.Rollback()
		if err != nil {
			fmt.Println("Error rolling back transaction", err)
		}
	}()

	var currentBalance domain.Balance

	q := `
		SELECT id, amount
		FROM balance
		WHERE account_id = ?
		AND category_id = ?
	`

	err = tx.QueryRow(q, accountId, categoryId).Scan(
		&currentBalance.Id,
		&currentBalance.Amount,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrBalanceNotFound
		}
	}

	if ok := isTransactionAllowed(currentBalance.Amount, amount); !ok {
		return nil, ErrInsufficientFunds
	}

	newBalance := currentBalance.Amount.Sub(amount)

	q = `
		UPDATE balance
		SET amount = ?
		WHERE account_id = ?
		AND category_id = ?
	`

	_, err = tx.Exec(q, newBalance, accountId, categoryId)

	if err != nil {
		return nil, ErrUpdatingBalance
	}

	q = `
		INSERT INTO transaction (account_id, balance_id, type, amount)
		VALUES (?, ?, ?, ?)
	`

	args := []interface{}{
		accountId,
		currentBalance.Id,
		domain.Debit,
		amount,
	}

	result, err := tx.Exec(q, args...)

	if err != nil {
		_ = tx.Rollback()
		return nil, ErrCreatingTransaction
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	q = `
		SELECT id, amount, type
		FROM transaction
		WHERE id = ?
	`

	var transaction domain.Transaction

	err = repo.db.QueryRow(q, lastInsertId).Scan(
		&transaction.Id,
		&transaction.Amount,
		&transaction.Type,
	)

	if err != nil {
		return nil, ErrTransactionNotFound
	}

	return &transaction, nil
}

func isTransactionAllowed(currentBalance decimal.Decimal, amount decimal.Decimal) bool {
	return currentBalance.GreaterThan(amount) || currentBalance.Equal(amount)
}

func (repo *mysqlBalanceRepository) GetBalance(accountId int, categoryId int) (*domain.Balance, error) {
	q := `
		SELECT id, amount
		FROM balance
		WHERE account_id = ?
		AND category_id = ?
	`

	args := []interface{}{
		accountId,
		categoryId,
	}

	var balance domain.Balance

	err := repo.db.QueryRow(q, args...).Scan(
		&balance.Id,
		&balance.Amount,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrBalanceNotFound
		}
	}

	return &balance, nil
}
