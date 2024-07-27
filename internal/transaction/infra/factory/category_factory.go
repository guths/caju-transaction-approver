package factory

import (
	"database/sql"

	"github.com/guths/caju-transaction-approver/internal/transaction/domain"
)

type CategoryFactory struct {
	db *sql.DB
}

func NewCategoryFactory(db *sql.DB) CategoryFactory {
	return CategoryFactory{
		db: db,
	}
}

func (f *CategoryFactory) CreateCategory(isFallback bool) (*domain.Category, error) {
	q := ` INSERT INTO category (name, is_fallback) VALUES (?, ?)`

	result, err := f.db.Exec(q, "FOOD", isFallback)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	q = `SELECT id, name FROM category WHERE id = ?`

	var category domain.Category

	err = f.db.QueryRow(q, id).Scan(&category.Id, &category.Name)

	if err != nil {
		return nil, err
	}

	return &category, nil
}
