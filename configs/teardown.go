package configs

import (
	"database/sql"
	"fmt"
	"testing"
)

func TearDown(tables []string, db *sql.DB, t *testing.T) {
	_, err := db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	if err != nil {
		t.Fatalf("Failed to disable foreign key checks: %v", err)
	}

	for _, table := range tables {
		query := fmt.Sprintf("TRUNCATE TABLE %s", table)
		_, err := db.Exec(query)
		if err != nil {
			t.Fatal(err)
		}
	}

	_, err = db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	if err != nil {
		t.Fatalf("Failed to enable foreign key checks: %v", err)
	}
}
