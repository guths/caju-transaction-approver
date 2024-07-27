package http_test

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/guths/caju-transaction-approver/cmd"
	"github.com/guths/caju-transaction-approver/configs"
	account_factory "github.com/guths/caju-transaction-approver/internal/account/infra/factory"
	"github.com/guths/caju-transaction-approver/internal/transaction/infra/factory"
	"github.com/shopspring/decimal"
)

var (
	server                 *gin.Engine
	invalidRequestsFormats []interface{} = []interface{}{
		map[string]interface{}{
			"account":     "xxx",
			"merchant":    "123",
			"mcc":         "123",
			"totalAmount": -10.00,
		},
		map[string]interface{}{
			"account":     1231231,
			"merchant":    "123",
			"mcc":         "123",
			"totalAmount": 10.00,
		},
		map[string]interface{}{
			"account":     "xxx",
			"merchant":    "123",
			"mcc":         "123",
			"totalAmount": 10.001,
		},
		map[string]interface{}{
			"account":     "xxx",
			"mcc":         "123",
			"totalAmount": 10.00,
		},
		map[string]interface{}{
			"account":     "xxx",
			"merchant":    "123",
			"totalAmount": 10.00,
		},
		map[string]interface{}{
			"account":  "xxx",
			"merchant": "123",
			"mcc":      "123",
		},
	}
)

func init() {
	server = cmd.GetServer()
}

// func TestAuthorizeTransactionWithInvalidRequestFormat(t *testing.T) {
// 	w := httptest.NewRecorder()

// 	for _, invalidRequestFormat := range invalidRequestsFormats {
// 		json, err := json.Marshal(invalidRequestFormat)

// 		if err != nil {
// 			t.Errorf("error marshalling payload: %v", err)
// 		}

// 		req := httptest.NewRequest("POST", "/authorize-transaction", strings.NewReader(string(json)))

// 		server.ServeHTTP(w, req)

// 		if w.Code != 200 {
// 			t.Errorf("expected status code 200, got %d", w.Code)
// 		}

// 		if !strings.Contains(w.Body.String(), "invalid request format") {
// 			t.Errorf("expected response to contain 'invalid request format'")
// 		}
// 	}
// }

func TestAuthorizeTransactionWithSuccess(t *testing.T) {
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

	_, err = balanceFactory.CreateBalance(account.Id, 2, decimal.New(100, 0))

	if err != nil {
		t.Errorf("Error creating balance: %v", err)
	}

	w := httptest.NewRecorder()

	payload := map[string]interface{}{
		"account":     account.AccountIdentifier,
		"merchant":    "123",
		"mcc":         "5411",
		"totalAmount": 10.00,
	}

	json, err := json.Marshal(payload)

	if err != nil {
		t.Errorf("error marshalling payload: %v", err)
	}

	req := httptest.NewRequest("POST", "/authorize-transaction", strings.NewReader(string(json)))

	server.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("expected status code 200, got %d", w.Code)
	}

	fmt.Println(w.Body.String())

	if !strings.Contains(w.Body.String(), "approved") {
		t.Errorf("expected response to contain 'approved'")
	}
}
