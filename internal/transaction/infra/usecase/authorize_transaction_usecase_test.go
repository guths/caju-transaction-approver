package usecase_test

import (
	"fmt"
	"testing"

	"github.com/guths/caju-transaction-approver/configs"
	account_factory "github.com/guths/caju-transaction-approver/internal/account/infra/factory"
	account_repo "github.com/guths/caju-transaction-approver/internal/account/infra/repository"
	account_service "github.com/guths/caju-transaction-approver/internal/account/infra/service"
	merchant_factory "github.com/guths/caju-transaction-approver/internal/merchant/infra/factory"
	merchant_repo "github.com/guths/caju-transaction-approver/internal/merchant/infra/repository"
	merchant_service "github.com/guths/caju-transaction-approver/internal/merchant/infra/service"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/factory"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/repository"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/service"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/usecase"
	"github.com/shopspring/decimal"
)

var (
	authorizeTransactionUseCase usecase.AuthorizeTransactionUseCase
	balanceService              service.BalanceService
)

func init() {
	balanceRepo := repository.NewBalanceRepository(configs.DB)
	balanceService = service.NewBalanceService(balanceRepo)
	merchantRepo := merchant_repo.NewMysqlMerchantRepository(configs.DB)
	merchantService := merchant_service.NewMerchantService(merchantRepo)
	mccRepo := repository.NewMccRepository(configs.DB)
	mccService := service.NewMccService(mccRepo)
	accountRepo := account_repo.NewMysqlAccountRepository(configs.DB)
	accountService := account_service.NewAccountService(accountRepo)
	authorizeTransactionUseCase = usecase.NewAuthorizeTransactionUseCase(balanceService, merchantService, mccService, *accountService)
}

func TestAuthorizeTransactionWithInvalidAccount(t *testing.T) {
	defer configs.TearDown([]string{"account", "balance", "category", "merchant", "mcc", "transaction"}, configs.DB, t)
	input := usecase.InputTransactionDTO{
		Account:     "invalid_account",
		TotalAmount: 100,
		Mcc:         "mcc",
		Merchant:    "merchant",
	}

	response := authorizeTransactionUseCase.Execute(input)

	if response.Code != "07" {
		t.Errorf("Expected code 07, got %s", response.Code)
	}

	if response.Message != "account not found" {
		t.Errorf("Expected message 'account not found', got %s", response.Message)
	}
}

func TestAuthorizeTransactionWithInvalidFallbackCategory(t *testing.T) {
	defer configs.TearDown([]string{"account", "balance", "category", "merchant", "mcc", "transaction"}, configs.DB, t)

	accountFactory := account_factory.NewAccountFactory(configs.DB)
	account, err := accountFactory.CreateAccount()

	if err != nil {
		t.Errorf("Error creating account: %v", err)
	}

	input := usecase.InputTransactionDTO{
		Account:     account.AccountIdentifier,
		TotalAmount: 100,
		Mcc:         "mcc",
		Merchant:    "merchant",
	}

	response := authorizeTransactionUseCase.Execute(input)

	if response.Code != "07" {
		t.Errorf("Expected code 07, got %s", response.Code)
	}

	if response.Message != "category not found" {
		t.Errorf("Expected message 'category not found', got %s", response.Message)
	}
}

func TestAuthorizeTransactionWithCorrectMcc(t *testing.T) {
	defer configs.TearDown([]string{"account", "balance", "category", "merchant", "mcc", "transaction"}, configs.DB, t)
	accountFactory := account_factory.NewAccountFactory(configs.DB)
	categoryFactory := factory.NewCategoryFactory(configs.DB)
	mccFactory := factory.NewMccFactory(configs.DB)
	balanceFactory := factory.NewBalanceFactory(configs.DB)

	account, err := accountFactory.CreateAccount()

	if err != nil {
		t.Errorf("Error creating account: %v", err)
	}

	_, err = categoryFactory.CreateCategory(true)

	if err != nil {
		t.Errorf("Error creating category: %v", err)
	}

	err = mccFactory.CreateMccWithCategory()

	if err != nil {
		t.Errorf("Error creating mcc: %v", err)
	}

	if err != nil {
		t.Errorf("Error creating category: %v", err)
	}

	balance, err := balanceFactory.CreateBalance(account.Id, 2, decimal.New(100, 0))

	fmt.Printf("%v\n", balance)
	if err != nil {
		t.Errorf("Error creating balance: %v", err)
	}

	input := usecase.InputTransactionDTO{
		Account:     account.AccountIdentifier,
		TotalAmount: 100.00,
		Mcc:         "5411",
		Merchant:    "merchant",
	}

	response := authorizeTransactionUseCase.Execute(input)

	if response.Code != "00" {
		t.Errorf("Expected code 00, got %s", response.Code)
	}

	currentBalance, err := balanceService.GetBalance(account.Id, 2)

	if err != nil {
		t.Errorf("Error getting balance: %v", err)
	}

	if currentBalance.Amount.String() != "0" {
		t.Errorf("Expected balance 0, got %s", currentBalance.Amount.String())
	}
}

func TestAuthorizeTransactionWithCorrectMccAndBalanceNotFound(t *testing.T) {
	defer configs.TearDown([]string{"account", "balance", "category", "merchant", "mcc", "transaction"}, configs.DB, t)
	accountFactory := account_factory.NewAccountFactory(configs.DB)
	categoryFactory := factory.NewCategoryFactory(configs.DB)
	mccFactory := factory.NewMccFactory(configs.DB)
	balanceFactory := factory.NewBalanceFactory(configs.DB)

	account, err := accountFactory.CreateAccount()

	if err != nil {
		t.Errorf("Error creating account: %v", err)
	}

	_, err = categoryFactory.CreateCategory(true)

	if err != nil {
		t.Errorf("Error creating category: %v", err)
	}

	err = mccFactory.CreateMccWithCategory()

	if err != nil {
		t.Errorf("Error creating mcc: %v", err)
	}

	if err != nil {
		t.Errorf("Error creating category: %v", err)
	}

	_, err = balanceFactory.CreateBalance(account.Id, 1, decimal.New(100, 0))

	if err != nil {
		t.Errorf("Error creating balance: %v", err)
	}

	input := usecase.InputTransactionDTO{
		Account:     account.AccountIdentifier,
		TotalAmount: 100.00,
		Mcc:         "5411",
		Merchant:    "merchant",
	}

	response := authorizeTransactionUseCase.Execute(input)

	if response.Code != "07" {
		t.Errorf("Expected code 07, got %s", response.Code)
	}

	if response.Message != "balance not found in db" {
		t.Errorf("Expected message 'balance not found', got %s", response.Message)
	}
}

func TestAuthorizeTransactionWithCorrectMccAndDebitFromFallbackBalance(t *testing.T) {
	defer configs.TearDown([]string{"account", "balance", "category", "merchant", "mcc", "transaction"}, configs.DB, t)
	accountFactory := account_factory.NewAccountFactory(configs.DB)
	categoryFactory := factory.NewCategoryFactory(configs.DB)
	mccFactory := factory.NewMccFactory(configs.DB)
	balanceFactory := factory.NewBalanceFactory(configs.DB)
	merchantFactory := merchant_factory.NewMerchantFactory(configs.DB)

	account, err := accountFactory.CreateAccount()

	if err != nil {
		t.Errorf("Error creating account: %v", err)
	}

	c, err := categoryFactory.CreateCategory(true)

	if err != nil {
		t.Errorf("Error creating category: %v", err)
	}

	merchant, err := merchantFactory.CreateMerchant("merchant", c.Id)

	if err != nil {
		t.Errorf("Error creating merchant: %v", err)
	}

	err = mccFactory.CreateMccWithCategory()

	if err != nil {
		t.Errorf("Error creating mcc: %v", err)
	}

	if err != nil {
		t.Errorf("Error creating category: %v", err)
	}

	_, err = balanceFactory.CreateBalance(account.Id, 1, decimal.New(100, 0))

	if err != nil {
		t.Errorf("Error creating balance: %v", err)
	}

	input := usecase.InputTransactionDTO{
		Account:     account.AccountIdentifier,
		TotalAmount: 100.00,
		Mcc:         "5412",
		Merchant:    merchant.Name,
	}

	response := authorizeTransactionUseCase.Execute(input)

	if response.Code != "00" {
		t.Errorf("Expected code 00, got %s", response.Code)
	}

	if response.Message != "approved" {
		t.Errorf("Expected message 'approved', got %s", response.Message)
	}

	currentFallBackBalance, err := balanceService.GetBalance(account.Id, 1)

	if err != nil {
		t.Errorf("Error getting balance: %v", err)
	}

	if currentFallBackBalance.Amount.String() != "0" {
		t.Errorf("Expected balance 0, got %s", currentFallBackBalance.Amount.String())
	}
}
