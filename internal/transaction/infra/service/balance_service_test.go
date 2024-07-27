package service_test

import (
	"testing"

	"github.com/guths/caju-transaction-approver/configs"
	account_factory "github.com/guths/caju-transaction-approver/internal/account/infra/factory"
	"github.com/guths/caju-transaction-approver/internal/transaction/domain"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/factory"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/repository"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/service"
	"github.com/shopspring/decimal"
)

var balanceTables = []string{"balance", "account", "category"}

func TestDebitAccountAmountWithFunds(t *testing.T) {
	defer configs.TearDown(balanceTables, configs.DB, t)
	accountFactory := account_factory.NewAccountFactory(configs.DB)
	categoryFactory := factory.NewCategoryFactory(configs.DB)
	balanceFactory := factory.NewBalanceFactory(configs.DB)

	account, err := accountFactory.CreateAccount()
	if err != nil {
		t.Error(err.Error())
	}

	category, err := categoryFactory.CreateCategory(false)

	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = balanceFactory.CreateBalance(account.Id, category.Id, decimal.New(100, 0))

	if err != nil {
		t.Fatalf(err.Error())
	}

	repo := repository.NewBalanceRepository(configs.DB)
	service := service.NewBalanceService(repo)

	amount := decimal.New(50, 0)

	transaction, err := service.DebitAmount(account.Id, category.Id, amount)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if transaction.Amount.Cmp(amount) != 0 {
		t.Fatalf("expected %v, got %v", amount, transaction.Amount)
	}

	q := `
		SELECT amount
		FROM balance
		WHERE account_id = ?
		AND category_id = ?
	`

	var balance domain.Balance

	err = configs.DB.QueryRow(q, account.Id, category.Id).Scan(&balance.Amount)

	if err != nil {
		t.Fatalf(err.Error())
	}

	rest := balance.Amount.Cmp(decimal.New(50, 0))

	if rest != 0 {
		t.Fatalf("expected 0, got %v", rest)
	}
}

func TestDebitAccountWithInsufficientFunds(t *testing.T) {
	defer configs.TearDown(balanceTables, configs.DB, t)
	accountFactory := account_factory.NewAccountFactory(configs.DB)
	categoryFactory := factory.NewCategoryFactory(configs.DB)
	balanceFactory := factory.NewBalanceFactory(configs.DB)

	account, err := accountFactory.CreateAccount()

	if err != nil {
		t.Error(err.Error())
	}

	category, err := categoryFactory.CreateCategory(false)

	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = balanceFactory.CreateBalance(account.Id, category.Id, decimal.New(100, 0))

	if err != nil {
		t.Fatalf(err.Error())
	}

	repo := repository.NewBalanceRepository(configs.DB)
	service := service.NewBalanceService(repo)

	amount := decimal.New(150, 0)

	_, err = service.DebitAmount(account.Id, category.Id, amount)

	if err != repository.ErrInsufficientFunds {
		t.Fatalf("expected %v, got %v", repository.ErrInsufficientFunds, err)
	}
}

func TestDebitFromNonExistBalance(t *testing.T) {
	defer configs.TearDown(balanceTables, configs.DB, t)
	accountFactory := account_factory.NewAccountFactory(configs.DB)
	categoryFactory := factory.NewCategoryFactory(configs.DB)

	account, err := accountFactory.CreateAccount()

	if err != nil {
		t.Error(err.Error())
	}

	category, err := categoryFactory.CreateCategory(false)

	if err != nil {
		t.Fatalf(err.Error())
	}

	repo := repository.NewBalanceRepository(configs.DB)
	service := service.NewBalanceService(repo)

	amount := decimal.New(150, 0)

	_, err = service.DebitAmount(account.Id, category.Id, amount)

	if err != repository.ErrBalanceNotFound {
		t.Fatalf("expected %v, got %v", repository.ErrBalanceNotFound, err)
	}
}
