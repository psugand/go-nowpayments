package custody

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/matn/go-nowpayments/config"
	"github.com/matn/go-nowpayments/core"
	"github.com/rotisserie/eris"
)

type CustodyUserAccountArgs struct {
	Name string `json:"name"`
}

type CustodyUserAccount struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// CustodyBalances hold multiple balances for Custody user account
type CustodyBalances struct {
	Usddtrc20 CustodyBalance `json:"usddtrc20"`
	Usdtbsc   CustodyBalance `json:"usdtbsc"`
}

// CustodyBalances single balance for Custody user account
type CustodyBalance struct {
	Amount        float64 `json:"amount"`
	PendingAmount float64 `json:"pendingAmount"`
}

// CustodyUserAccountBalance hold response for a specific Custody user account balance
type CustodyUserAccountBalance struct {
	SubPartnerID string          `json:"subPartnerId"`
	Balances     CustodyBalances `json:"balances"`
}

// Create will initiate new user account from a unique ID
// JWT is required for this request
func Create(cu *CustodyUserAccountArgs) (*CustodyUserAccount, error) {
	if cu == nil {
		return nil, errors.New("nil custody user account args")
	}

	d, err := json.Marshal(cu)
	if err != nil {
		return nil, eris.Wrap(err, "custody user account args")
	}

	tok, err := core.Authenticate(config.Login(), config.Password())
	if err != nil {
		return nil, eris.Wrap(err, "custody user account")
	}

	// CONSISTENCY PROBLEM ON THEIR SIDE: for some requests response is put under result object
	type result struct {
		Result *CustodyUserAccount `json:"result"`
	}

	us := &result{}

	par := &core.SendParams{
		RouteName: "custody-create-account",
		Into:      &us,
		Body:      strings.NewReader(string(d)),
		Token:     tok,
	}

	err = core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return us.Result, nil
}

// Status gets the actual information about the payment. You need to provide the payment ID
// This endpoint will work only if IP is whitelisted (or white IP restrictions are disabled)
func Balance(userAccountID string) (*CustodyUserAccountBalance, error) {
	if userAccountID == "" {
		return nil, eris.New("empty user account ID")
	}

	// CONSISTENCY PROBLEM ON THEIR SIDE: for some requests response is put under result object
	type result struct {
		Result *CustodyUserAccountBalance `json:"result"`
	}

	bl := &result{}

	par := &core.SendParams{
		RouteName: "custody-account-balance",
		Path:      userAccountID,
		Into:      &bl,
	}

	err := core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return bl.Result, nil
}
