package service_test

import (
	"testing"

	"github.com/guths/caju-transaction-approver/configs"
	"github.com/guths/caju-transaction-approver/internal/account/infra/factory"
	"github.com/guths/caju-transaction-approver/internal/account/infra/repository"
	"github.com/guths/caju-transaction-approver/internal/account/infra/service"
)

func TestGetAccountByIdentifierName(t *testing.T) {
	defer configs.TearDown([]string{"account"}, configs.DB, t)

	factory := factory.NewAccountFactory(configs.DB)

	_, err := factory.CreateAccount()

	if err != nil {
		t.Fatalf("error executing factory")
	}

	repo := repository.NewMysqlAccountRepository(configs.DB)

	service := service.NewAccountService(repo)

	account, err := service.GetAccountByIdentifier("123")

	if err != nil {
		t.Fatalf(err.Error())
	}

	if account.Id != 1 {
		t.Fatalf("invalid account id")
	}
}
