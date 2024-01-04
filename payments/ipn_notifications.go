package payments

type IPNPaymentFees struct {
	Currency      string  `json:"currency"`
	DepositFee    float64 `json:"depositFee"`
	ServiceFee    float64 `json:"serviceFee"`
	WithdrawalFee float64 `json:"withdrawalFee"`
}

// PaymentStatus holds payment status related information
// Docs found on https://documenter.getpostman.com/view/7907941/2s93JusNJt#62a6d281-478d-4927-8cd0-f96d677b8de6
// Docs said IPN response is similar to PaymentStatus, but it's not the case
// And because this struct is used to compare a signature (callback), must be exactly the same
type IPNPaymentStatus struct {
	ActuallyPaid       float64        `json:"actually_paid"`
	ActuallyPaidAtFiat float64        `json:"actually_paid_at_fiat"`
	Fee                IPNPaymentFees `json:"fee"`
	InvoiceID          int64          `json:"invoice_id"`
	OrderDescription   string         `json:"order_description"`
	OrderID            string         `json:"order_id"`
	OutcomeAmount      float64        `json:"outcome_amount"`
	OutcomeCurrency    string         `json:"outcome_currency"`
	ParentPaymentId    *int64         `json:"parent_payment_id"`
	PayAddress         string         `json:"pay_address"`
	PayAmount          float64        `json:"pay_amount"`
	PayCurrency        string         `json:"pay_currency"`
	PayinExtraID       *int64         `json:"payin_extra_id"`
	PaymentExtraIds    []int64        `json:"payment_extra_ids"`
	PaymentID          int64          `json:"payment_id"`
	PaymentStatus      string         `json:"payment_status"`
	PriceAmount        float64        `json:"price_amount"`
	PriceCurrency      string         `json:"price_currency"`
	PurchaseID         string         `json:"purchase_id"`
}
