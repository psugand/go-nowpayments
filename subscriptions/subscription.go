package subscriptions

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/matn/go-nowpayments/config"
	"github.com/matn/go-nowpayments/core"
	recurringPayment "github.com/matn/go-nowpayments/recurring_payments"
	"github.com/rotisserie/eris"
)

// SubscriptionArgs handle args to create a subscription plan
type SubscriptionArgs struct {
	Title       float64 `json:"title"`
	IntervalDay int64   `json:"interval_day"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
}

// EmailSubscriptionArgs handle args to create a subscription with an email
type EmailSubscriptionArgs struct {
	SubscriptionPlanID int64  `json:"subscription_plan_id"`
	Email              string `json:"email"`
}

// Subscription handle subscription plan
type Subscription struct {
	ID               string  `json:"id"`
	Title            string  `json:"title"`
	IntervalDay      string  `json:"interval_day"`
	IpnCallbackURL   string  `json:"ipn_callback_url"`
	SuccessURL       string  `json:"success_url"`
	CancelURL        string  `json:"cancel_url"`
	PartiallyPaidURL string  `json:"partially_paid_url"`
	Amount           float64 `json:"amount"`
	Currency         string  `json:"currency"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
}

// New create a subscription plan
// JWT is required for this request
func New(su *SubscriptionArgs) (*Subscription, error) {
	if su == nil {
		return nil, errors.New("nil subscription args")
	}

	d, err := json.Marshal(su)
	if err != nil {
		return nil, eris.Wrap(err, "subscription args")
	}

	tok, err := core.Authenticate(config.Login(), config.Password())
	if err != nil {
		return nil, eris.Wrap(err, "subscription")
	}

	s := &Subscription{}

	par := &core.SendParams{
		RouteName: "subscription-create",
		Into:      &s,
		Body:      strings.NewReader(string(d)),
		JWTToken:  tok,
	}

	err = core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// NewWithEmail create an email subscription with specific plan ID
// JWT is required for this request
func NewWithEmail(su *EmailSubscriptionArgs) (*recurringPayment.RecurringPayment, error) {
	if su == nil {
		return nil, errors.New("nil subscription email args")
	}

	d, err := json.Marshal(su)
	if err != nil {
		return nil, eris.Wrap(err, "subscription email args")
	}

	tok, err := core.Authenticate(config.Login(), config.Password())
	if err != nil {
		return nil, eris.Wrap(err, "subscription")
	}

	s := &recurringPayment.RecurringPayment{}

	par := &core.SendParams{
		RouteName: "subscription-create-email",
		Into:      &s,
		JWTToken:  tok,
		Body:      strings.NewReader(string(d)),
	}

	err = core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// Update update a subscription plan
// JWT is required for this request
func Update(su *SubscriptionArgs) (*Subscription, error) {
	if su == nil {
		return nil, errors.New("nil subscription args")
	}

	d, err := json.Marshal(su)
	if err != nil {
		return nil, eris.Wrap(err, "subscription args")
	}

	tok, err := core.Authenticate(config.Login(), config.Password())
	if err != nil {
		return nil, eris.Wrap(err, "subscription")
	}

	s := &Subscription{}

	par := &core.SendParams{
		RouteName: "subscription-update",
		Into:      &s,
		JWTToken:  tok,
		Body:      strings.NewReader(string(d)),
	}

	err = core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// Get return a single subscription plan by ID
func Get(subscriptionPlanID string) (*Subscription, error) {
	if subscriptionPlanID == "" {
		return nil, eris.New("empty subscription plan ID")
	}

	st := &Subscription{}

	par := &core.SendParams{
		RouteName: "subscription-single",
		Path:      subscriptionPlanID,
		Into:      &st,
	}

	err := core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return st, nil
}
