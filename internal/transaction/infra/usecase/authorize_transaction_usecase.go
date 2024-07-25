package usecase

import (
	"strconv"

	merchant_service "github.com/guths/caju-transaction-approver/internal/merchant/infra/service"
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

	category, err := uc.mccService.GetCategoryByMcc(inputAuthorizeTransactionDTO.Mcc)

	if err == nil {
		//olhar saldo
		accountAmount, err := uc.balanceService.GetAmountByAccountId(accountId, category.Id)

		if err != nil {
			return err
		}

		//se tiver desconta
		if ok := uc.balanceService.IsBalanceSufficient(amount, accountAmount); ok {
			//descontar saldo aki

			return nil
		}

		//olhar saldo cash
		fallbackAmount, err := uc.balanceService.GetFallbackBalanceAmount(accountId)

		if err != nil {
			return err
		}

		if ok := uc.balanceService.IsBalanceSufficient(amount, fallbackAmount); !ok {
			//retornar saldo insuficiente

			return nil
		}

		//descontar saldo aki

		return nil
	}

	if err != service.ErrCategoryNotFound {
		return err
	}

	category, err = uc.merchantService.GetCategoryByMerchantName(inputAuthorizeTransactionDTO.Merchant)

	if err != nil {
		//retornar erro genérico de merchant não encontrado
		return err
	}

	//olhar saldo
	accountAmount, err := uc.balanceService.GetAmountByAccountId(accountId, category.Id)

	if err != nil {
		return err
	}

	if ok := uc.balanceService.IsBalanceSufficient(amount, accountAmount); ok {
		//descontar saldo aki
		return nil
	}

	fallbackAmount, err := uc.balanceService.GetFallbackBalanceAmount(accountId)

	if err != nil {
		return err
	}

	if ok := uc.balanceService.IsBalanceSufficient(amount, fallbackAmount); !ok {
		//retornar saldo insuficiente
		return nil
	}

	//descontar saldo
	return nil
}
