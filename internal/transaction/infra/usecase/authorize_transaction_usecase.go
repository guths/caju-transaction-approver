package usecase

import (
	account_service "github.com/guths/caju-transaction-approver/internal/account/infra/service"
	merchant_service "github.com/guths/caju-transaction-approver/internal/merchant/infra/service"
	"github.com/guths/caju-transaction-approver/internal/transaction/domain"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/repository"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/service"
	"github.com/guths/caju-transaction-approver/internal/validator"
	"github.com/shopspring/decimal"
)

type AuthorizeTransactionUseCase struct {
	balanceService  service.BalanceService
	merchantService merchant_service.MerchantService
	mccService      service.MccService
	accountService  account_service.AccountService
}

type InputTransactionDTO struct {
	Account     string  `json:"account"`
	TotalAmount float64 `json:"totalAmount"`
	Mcc         string  `json:"mcc"`
	Merchant    string  `json:"merchant"`
}

func ValidateInputTransaction(v *validator.Validator, input InputTransactionDTO) {
	v.Check(input.Account != "", "account", "account is required")
	v.Check(validator.IsString(input.Account), "account", "account must be a string")
	v.Check(input.TotalAmount > 0.00, "total_amount", "total_amount must be greater than 0")
	v.Check(validator.IsValidFloatPattern(input.TotalAmount), "total_amount", "total_amount must have 2 decimal places")
	v.Check(input.Merchant != "", "merchant", "merchant is required")
	v.Check(validator.IsString(input.Merchant), "merchant", "merchant must be a string")
	v.Check(input.Mcc != "", "mcc", "mcc is required")
	v.Check(validator.IsString(input.Mcc), "mcc", "mcc must be a string")
}

type OutputTransactionDTO struct {
	Success          bool            `json:"success"`
	Message          string          `json:"message"`
	Account          int             `json:"account"`
	Merchant         string          `json:"merchant"`
	AmountTransfered decimal.Decimal `json:"amount_transfered"`
	Category         string          `json:"category"`
	TransactionID    int             `json:"transaction_id"`
}

func NewAuthorizeTransactionUseCase(
	balanceService service.BalanceService,
	merchantService merchant_service.MerchantService,
	mccService service.MccService,
	accountService account_service.AccountService,
) AuthorizeTransactionUseCase {
	return AuthorizeTransactionUseCase{
		balanceService:  balanceService,
		merchantService: merchantService,
		mccService:      mccService,
		accountService:  accountService,
	}
}

func (uc *AuthorizeTransactionUseCase) Execute(inputAuthorizeTransactionDTO InputTransactionDTO) domain.Response {
	amount := decimal.NewFromFloat(inputAuthorizeTransactionDTO.TotalAmount)
	account, err := uc.accountService.GetAccountByIdentifier(inputAuthorizeTransactionDTO.Account)

	if err != nil {
		return domain.GetGenericResponseError(err.Error())
	}

	fallbackCategory, err := uc.mccService.GetFallbackCategory()

	if err != nil {
		return domain.GetGenericResponseError(err.Error())
	}

	category, err := uc.mccService.GetCategoryByMcc(inputAuthorizeTransactionDTO.Mcc)

	if err == nil {
		transaction, err := uc.balanceService.DebitAmount(account.Id, category.Id, amount)

		if err == nil {
			return domain.GetApprovedResponse()
		}

		if err != repository.ErrInsufficientFunds && transaction == nil {
			return domain.GetGenericResponseError(err.Error())
		}

		_, err = uc.balanceService.DebitAmount(account.Id, fallbackCategory.Id, amount)

		if err == nil {
			return domain.GetApprovedResponse()
		}

		return domain.GetRejectedResponse()
	}

	if err != repository.ErrCategoryNotFound {
		return domain.GetGenericResponseError(err.Error())
	}

	c, err := uc.merchantService.GetCategoryByMerchantName(inputAuthorizeTransactionDTO.Merchant)

	if err != nil {
		return domain.GetGenericResponseError(err.Error())
	}

	_, err = uc.balanceService.DebitAmount(account.Id, c.Id, amount)

	if err == nil {
		return domain.GetApprovedResponse()
	}

	if err != repository.ErrInsufficientFunds {
		return domain.GetGenericResponseError(err.Error())
	}

	if c.Id == fallbackCategory.Id {
		return domain.GetRejectedResponse()
	}

	_, err = uc.balanceService.DebitAmount(account.Id, fallbackCategory.Id, amount)

	if err != nil {
		return domain.GetRejectedResponse()
	}

	return domain.GetApprovedResponse()
}
