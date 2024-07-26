package service_test

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/guths/caju-transaction-approver/configs"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/factory"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/repository"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/service"
)

var mccTables = []string{"mcc", "category"}

func TearDown(tables []string, db *sql.DB, t *testing.T) {
	_, err := db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	if err != nil {
		log.Fatalf("Failed to disable foreign key checks: %v", err)
	}

	for _, table := range tables {
		query := fmt.Sprintf("DELETE FROM %s", table)
		_, err := db.Exec(query)
		if err != nil {
			t.Logf("Failed to delete data from %s: %v", table, err)
		} else {
			t.Logf("Data deleted from %s successfully", table)
		}
	}

	_, err = db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	if err != nil {
		log.Fatalf("Failed to enable foreign key checks: %v", err)
	}
}

func TestGetCategoryByMcc(t *testing.T) {
	defer TearDown(mccTables, configs.DB, t)

	mccFactory := factory.NewMccFactory(configs.DB)

	err := mccFactory.CreateMccWithCategory()

	if err != nil {
		fmt.Println(err.Error())
		t.Error("error executing factory")
	}

	repo := repository.NewMccRepository(configs.DB)

	service := service.NewMccService(repo)

	category, err := service.GetCategoryByMcc("5411")

	if err != nil {
		t.Error(err.Error())
	}

	if category.Id != 1 {
		t.Error("invalid category id")
	}

	if category.Name != "FOOD" {
		t.Error("invalid category food")
	}
}

func TestGetCategoryNotFoundByMcc(t *testing.T) {
	defer TearDown(mccTables, configs.DB, t)
	repo := repository.NewMccRepository(configs.DB)

	service := service.NewMccService(repo)

	_, err := service.GetCategoryByMcc("5411")

	if err == nil {
		t.Error("category must not exist")
	}
}
