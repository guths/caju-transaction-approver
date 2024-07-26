package factory

import (
	"database/sql"

	"github.com/guths/caju-transaction-approver/internal/account/domain"
)

type AccountFactory struct {
	db *sql.DB
}

func NewAccountFactory(db *sql.DB) *AccountFactory {
	return &AccountFactory{db: db}
}

func (f *AccountFactory) CreateAccount() (*domain.Account, error) {
	q := ` INSERT INTO account (account_identifier, name) VALUES (?, ?)`

	result, err := f.db.Exec(q, "123", "John Doe")

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	q = `SELECT id, account_identifier, name FROM account WHERE id = ?`

	var account domain.Account

	err = f.db.QueryRow(q, id).Scan(&account.Id, &account.AccountIdentifier, &account.Name)

	if err != nil {
		return nil, err
	}

	return &account, nil
}
