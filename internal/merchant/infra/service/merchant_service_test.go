package service_test

import (
	"testing"

	"github.com/guths/caju-transaction-approver/configs"
	"github.com/guths/caju-transaction-approver/internal/merchant/infra/factory"
	"github.com/guths/caju-transaction-approver/internal/merchant/infra/repository"
	"github.com/guths/caju-transaction-approver/internal/merchant/infra/service"
	category_factory "github.com/guths/caju-transaction-approver/internal/transaction/infra/factory"
)

func TestCategoryByMerchantName(t *testing.T) {
	defer configs.TearDown([]string{"merchant", "category"}, configs.DB, t)
	factory := factory.NewMerchantFactory(configs.DB)
	categoryFactory := category_factory.NewCategoryFactory(configs.DB)

	category, err := categoryFactory.CreateCategory()

	if err != nil {
		t.Fatalf("error executing factory")
	}

	m, err := factory.CreateMerchant("Test Merchant", category.Id)

	if err != nil {
		t.Fatalf("error executing factory")
	}

	repo := repository.NewMysqlMerchantRepository(configs.DB)
	service := service.NewMerchantService(repo)

	merchant, err := service.GetCategoryByMerchantName(m.Name)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if merchant.Id != 1 {
		t.Fatalf("invalid category id")
	}

}

func TestCategoryNotFoundByMerchantName(t *testing.T) {
	defer configs.TearDown([]string{"merchant", "category"}, configs.DB, t)
	factory := factory.NewMerchantFactory(configs.DB)
	categoryFactory := category_factory.NewCategoryFactory(configs.DB)

	c, err := categoryFactory.CreateCategory()

	if err != nil {
		t.Fatalf("error executing factory")
	}

	_, err = factory.CreateMerchant("Test Merchant", c.Id)

	if err != nil {
		t.Fatalf("error executing factory")
	}

	repo := repository.NewMysqlMerchantRepository(configs.DB)
	service := service.NewMerchantService(repo)

	_, err = service.GetCategoryByMerchantName("xxx")

	if err == nil {
		t.Fatalf("error expected")
	}
}
