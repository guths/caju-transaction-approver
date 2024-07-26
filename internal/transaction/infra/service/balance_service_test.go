package service_test

import (
	"testing"

	"github.com/guths/caju-transaction-approver/configs"
	account_factory "github.com/guths/caju-transaction-approver/internal/account/infra/factory"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/factory"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/repository"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/service"
	"github.com/shopspring/decimal"
)

var balanceTables = []string{"balance", "account", "category"}

func TestDebitAccountAmountWithFunds(t *testing.T) {
	defer TearDown(balanceTables, configs.DB, t)
	accountFactory := account_factory.NewAccountFactory(configs.DB)
	categoryFactory := factory.NewCategoryFactory(configs.DB)
	balanceFactory := factory.NewBalanceFactory(configs.DB)

	account, err := accountFactory.CreateAccount()

	if err != nil {
		t.Error(err.Error())
	}

	category, err := categoryFactory.CreateCategory()

	if err != nil {
		t.Error(err.Error())
	}

	_, err = balanceFactory.CreateBalance(account.Id, category.Id, decimal.New(100, 0))

	if err != nil {
		t.Error(err.Error())
	}

	repo := repository.NewBalanceRepository(configs.DB)
	service := service.NewBalanceService(repo)

	amount := decimal.New(50, 0)

	transaction, err := service.DebitAmount(account.Id, category.Id, amount)

	if err != nil {
		t.Error(err.Error())
	}

	if transaction.Amount != amount {
		t.Errorf("expected %v, got %v", amount, transaction.Amount)
	}

	q := `
		SELECT amount
		FROM balance
		WHERE account_id = $1
		AND category_id = $2
	`

	var balance decimal.Decimal

	err = configs.DB.QueryRow(q, account.Id, category.Id).Scan(&balance)

	if err != nil {
		t.Error(err.Error())
	}

	rest := balance.Cmp(decimal.New(50, 0))

	if rest != 0 {
		t.Errorf("expected 0, got %v", rest)
	}
}
