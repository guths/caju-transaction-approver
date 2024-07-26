package service_test

import (
	"testing"

	"github.com/guths/caju-transaction-approver/configs"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/factory"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/repository"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/service"
)

var mccTables = []string{"mcc", "category"}

func TestGetCategoryByMcc(t *testing.T) {
	defer configs.TearDown(mccTables, configs.DB, t)

	mccFactory := factory.NewMccFactory(configs.DB)

	err := mccFactory.CreateMccWithCategory()

	if err != nil {
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
	defer configs.TearDown(mccTables, configs.DB, t)
	repo := repository.NewMccRepository(configs.DB)

	service := service.NewMccService(repo)

	_, err := service.GetCategoryByMcc("5411")

	if err == nil {
		t.Error("category must not exist")
	}
}
