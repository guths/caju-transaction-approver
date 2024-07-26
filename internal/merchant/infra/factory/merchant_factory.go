package factory

import (
	"database/sql"

	"github.com/guths/caju-transaction-approver/internal/merchant/domain"
)

type MerchantFactory struct {
	db *sql.DB
}

func NewMerchantFactory(db *sql.DB) MerchantFactory {
	return MerchantFactory{
		db: db,
	}
}

func (factory *MerchantFactory) CreateMerchant(name string, categoryId int) (*domain.Merchant, error) {
	q := `
		INSERT INTO merchant (name, category_id)
		VALUES (?, ?)
	`

	result, err := factory.db.Exec(q, name, categoryId)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	q = `
		SELECT id, name
		FROM merchant
		WHERE id = ?
	`
	var merchant domain.Merchant

	err = factory.db.QueryRow(q, id).Scan(
		&merchant.Id,
		&merchant.Name,
	)

	if err != nil {
		return nil, err
	}

	return &merchant, nil
}
