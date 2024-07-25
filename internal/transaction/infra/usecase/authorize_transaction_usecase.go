package usecase

import (
	"fmt"
	"strconv"

	merchant_service "github.com/guths/caju-transaction-approver/internal/merchant/infra/service"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/repository"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/service"
	"github.com/shopspring/decimal"
)

type AuthorizeTransactionUseCase struct {
	balanceService  service.BalanceService
	merchantService merchant_service.MerchantService
	mccService      service.MccService
}

type InputTransactionDTO struct {
	Account     string  `json:"account"`
	TotalAmount float64 `json:"total_amount"`
	Mcc         string  `json:"mcc"`
	Merchant    string  `json:"merchant"`
}

func NewAuthorizeTransactionUseCase(balanceService service.BalanceService) AuthorizeTransactionUseCase {
	return AuthorizeTransactionUseCase{
		balanceService: balanceService,
	}
}

func (uc *AuthorizeTransactionUseCase) Execute(inputAuthorizeTransactionDTO InputTransactionDTO) error {
	amount := decimal.NewFromFloat(inputAuthorizeTransactionDTO.TotalAmount)

	accountId, err := strconv.Atoi(inputAuthorizeTransactionDTO.Account)

	if err != nil {
		return err
	}

	fallbackCategory, err := uc.mccService.GetFallbackCategory()

	if err != nil {
		return err
	}

	category, err := uc.mccService.GetCategoryByMcc(inputAuthorizeTransactionDTO.Mcc)

	if err == nil {
		transaction, err := uc.balanceService.DebitAmount(accountId, category.Id, amount)

		if err == nil {
			//retornar transacao aqui //SUCESSO
			return nil
		}

		if err != repository.ErrInsufficientFunds && transaction == nil {
			return err
		}

		transaction, err = uc.balanceService.DebitAmount(accountId, fallbackCategory.Id, amount)

		if err == nil {
			//retornar transacao aqui //SUCESSO
			return nil
		}

		if err == repository.ErrInsufficientFunds && transaction == nil {
			return err
		}

		return err
	}

	if err != repository.ErrCategoryNotFound {
		return err
	}

	category, err = uc.merchantService.GetCategoryByMerchantName(inputAuthorizeTransactionDTO.Merchant)

	if err != nil {
		//retornar erro genérico de merchant não encontrado
		return err
	}

	transaction, err := uc.balanceService.DebitAmount(accountId, category.Id, amount)

	if err == nil {
		fmt.Println(transaction)
		//sucesso
		//retornar transacao
		return nil
	}

	if err != repository.ErrInsufficientFunds {
		//erro generico
		return err
	}

	transaction, err = uc.balanceService.DebitAmount(accountId, fallbackCategory.Id, amount)

	if err != nil {
		//retornar qualquer erro aqui
		return err
	}

	//retornar transacao certinha aqui
	return nil
}
