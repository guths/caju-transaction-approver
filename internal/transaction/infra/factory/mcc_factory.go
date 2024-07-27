package factory

import (
	"database/sql"

	"github.com/guths/caju-transaction-approver/internal/transaction/domain"
)

type MccFactory struct {
	db *sql.DB
}

func NewMccFactory(db *sql.DB) MccFactory {
	return MccFactory{
		db: db,
	}
}

func (f *MccFactory) CreateMccWithCategory() error {
	q := `
		INSERT INTO category (name)
		VALUES (?)
	`

	result, err := f.db.Exec(q, "FOOD")

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	var category domain.Category

	q = `
		SELECT id, name
		FROM category
		WHERE id = ?
	`
	err = f.db.QueryRow(q, id).Scan(&category.Id, &category.Name)

	if err != nil {
		return err
	}

	q = `
		INSERT INTO mcc (mcc, category_id)
		VALUES (?, ?)
	`

	_, err = f.db.Exec(q, "5411", category.Id)

	if err != nil {
		return err
	}

	return nil
}
