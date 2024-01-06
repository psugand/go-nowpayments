package recurring_payments

import (
	"github.com/matm/go-nowpayments/config"
	"github.com/matm/go-nowpayments/core"
	"github.com/rotisserie/eris"
)

// RecurringPayment handle status of a specific recurring payment
type RecurringPayment struct {
	ID                 string     `json:"id"`
	SubscriptionPlanID string     `json:"subscription_plan_id"`
	IsActive           bool       `json:"is_active"`
	Status             string     `json:"status"`
	ExpireDate         string     `json:"expire_date"`
	Subscriber         Subscriber `json:"subscriber"`
	CreatedAt          string     `json:"created_at"`
	UpdatedAt          string     `json:"updated_at"`
}

// DeleteReccurringPayment handle status when deleting recurring payment
type DeleteReccurringPayment struct {
	Status string `json:"status"`
}

// Subscriber handle a subscriber to a specific plan
type Subscriber struct {
	Email        string `json:"email,omniempty"`
	SubPartnerID string `json:"sub_partner_id,omniempty"`
}

// Get return a single reccuring payment via it's ID
func Get(recurringPaymentID string) (*RecurringPayment, error) {
	if recurringPaymentID == "" {
		return nil, eris.New("empty recurring payment ID")
	}

	tok, err := core.Authenticate(config.Login(), config.Password())
	if err != nil {
		return nil, eris.Wrap(err, "status")
	}

	st := &RecurringPayment{}

	par := &core.SendParams{
		RouteName: "recurring-payment-single",
		Path:      recurringPaymentID,
		Into:      &st,
		Token:     tok,
	}

	err = core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return st, nil
}

// Delete remove a recurring payment via it's ID
func Delete(recurringPaymentID string) (*DeleteReccurringPayment, error) {
	if recurringPaymentID == "" {
		return nil, eris.New("empty recurring payment ID")
	}

	tok, err := core.Authenticate(config.Login(), config.Password())
	if err != nil {
		return nil, eris.Wrap(err, "status")
	}

	de := &DeleteReccurringPayment{}

	par := &core.SendParams{
		RouteName: "recurring-payment-delete",
		Path:      recurringPaymentID,
		Into:      &de,
		Token:     tok,
	}

	err = core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return de, nil
}
