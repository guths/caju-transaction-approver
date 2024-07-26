package factory

import "database/sql"

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

	_, err := f.db.Exec(q, "FOOD")

	if err != nil {
		return err
	}

	q = `
		INSERT INTO mcc (mcc, category_id)
		VALUES (?, ?)
	`

	_, err = f.db.Exec(q, "5411", 1)

	if err != nil {
		return err
	}

	return nil
}
