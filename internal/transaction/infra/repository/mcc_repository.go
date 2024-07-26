package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/guths/caju-transaction-approver/internal/transaction/domain"
)

var (
	ErrCategoryNotFound = fmt.Errorf("category not found")
)

type mysqlMccRepository struct {
	db *sql.DB
}

func NewMccRepository(db *sql.DB) domain.MccRepository {
	return &mysqlMccRepository{
		db: db,
	}
}

func (repo *mysqlMccRepository) GetFallbackCategory() (*domain.Category, error) {
	q := `
		SELECT id, name
		FROM category
		WHERE category.is_fallback = ?
	`
	var category domain.Category

	err := repo.db.QueryRow(q, true).Scan(
		&category.Id,
		&category.Name,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrCategoryNotFound
		}

		return nil, err
	}

	return &category, nil
}

func (repo *mysqlMccRepository) GetCategoryByMcc(mcc string) (domain.Category, error) {
	q := `
		SELECT category.id, category.name
		FROM mcc
		INNER JOIN category 
		ON mcc.category_id = category.id
		WHERE mcc.mcc = ?
	`
	args := []interface{}{mcc}

	var category domain.Category

	err := repo.db.QueryRow(q, args...).Scan(
		&category.Id,
		&category.Name,
	)

	fmt.Printf("CATEGORY %v", category)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return domain.Category{}, err
		}
	}

	return category, nil
}
