package custody

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/CIDgravity/go-nowpayments/config"
	"github.com/CIDgravity/go-nowpayments/core"
	"github.com/CIDgravity/go-nowpayments/payments"
	"github.com/rotisserie/eris"
)

type DepositArgs struct {
	Currency     string
	Amount       float64
	SubPartnerID string
}

type DepositWithPaymentArgs struct {
	DepositArgs
	IsFixedRate     bool
	IsFeePaidByUser bool
	IpnCallbackURL  string
}

// NewDepositWithPayment will create a payment to deposit on a specific user account (refill account)
// The response doesn't provide the payment link, but can be built using https://nowpayments.io/payment/?iid=[INVOICE_ID]&paymentId=[PAYMENT_id]
// JWT is required for this request
func NewDepositWithPayment(da *DepositWithPaymentArgs) (*payments.Payment, error) {
	if da == nil {
		return nil, errors.New("nil deposit args")
	}

	d, err := json.Marshal(da)
	if err != nil {
		return nil, eris.Wrap(err, "deposit args")
	}

	tok, err := core.Authenticate(config.Login(), config.Password())
	if err != nil {
		return nil, eris.Wrap(err, "deposit with payment")
	}

	dp := &core.V2ResponseFormat[*payments.Payment]{}
	par := &core.SendParams{
		RouteName: "custody-deposit-with-payment",
		Into:      &dp,
		Body:      strings.NewReader(string(d)),
		JWTToken:  tok,
	}

	err = core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return dp.Result, nil
}

// NewDepositFroMasterAccount will create a deposit on a specific user account from a master account (no payment link, will use balance from master)
// JWT is required for this request
func NewDepositFroMasterAccount(da *DepositArgs) (*Transfer, error) {
	if da == nil {
		return nil, errors.New("nil deposit args")
	}

	d, err := json.Marshal(da)
	if err != nil {
		return nil, eris.Wrap(err, "deposit args")
	}

	tok, err := core.Authenticate(config.Login(), config.Password())
	if err != nil {
		return nil, eris.Wrap(err, "deposit from master account")
	}

	tr := &core.V2ResponseFormat[*Transfer]{}
	par := &core.SendParams{
		RouteName: "custody-deposit-from-master",
		Into:      &tr,
		Body:      strings.NewReader(string(d)),
		JWTToken:  tok,
	}

	err = core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return tr.Result, nil
}
