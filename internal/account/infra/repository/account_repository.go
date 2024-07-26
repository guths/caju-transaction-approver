package repository

import (
	"database/sql"
	"fmt"

	"github.com/guths/caju-transaction-approver/internal/account/domain"
)

type mysqlAccountRepository struct {
	db *sql.DB
}

var (
	ErrorAccountNotFound = fmt.Errorf("account not found")
)

func NewMysqlAccountRepository(db *sql.DB) *mysqlAccountRepository {
	return &mysqlAccountRepository{db: db}
}

func (r *mysqlAccountRepository) GetAccountByIdentifier(accountIdentifier string) (*domain.Account, error) {
	q := `
		SELECT id, account_identifier, name 
		FROM account 
		WHERE account_identifier = ?`

	var account domain.Account

	err := r.db.QueryRow(q, accountIdentifier).Scan(&account.Id, &account.AccountIdentifier, &account.Name)

	switch {
	case err == sql.ErrNoRows:
		return nil, ErrorAccountNotFound
	case err != nil:
		return nil, err
	}

	return &account, nil
}
