package repository

import (
	"database/sql"
	"errors"

	"github.com/guths/caju-transaction-approver/internal/transaction/domain"
)

type mysqlMccRepository struct {
	db *sql.DB
}

func NewMccRepository(db *sql.DB) domain.MccRepository {
	return &mysqlMccRepository{
		db: db,
	}
}

func (repo *mysqlMccRepository) GetCategoryByMcc(mmc string) (domain.Category, error) {
	q := `
		SELECT category.id, category.name
		FROM mcc
		INNER JOIN category 
		ON mcc.category_id = category.id
		WHERE mcc.mcc = $1
	`
	args := []interface{}{
		mmc,
	}

	var category domain.Category

	err := repo.db.QueryRow(q, args...).Scan(
		&category.Id,
		&category.Name,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return domain.Category{}, err
		}
	}

	return category, nil
}
