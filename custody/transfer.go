package custody

import (
	"fmt"
	"net/url"

	"github.com/matn/go-nowpayments/config"
	"github.com/matn/go-nowpayments/core"
	"github.com/rotisserie/eris"
)

// ListTransfersOption are options applying to the list of transfers
type ListTransfersOption struct {
	Id     int64
	Status string
	Limit  int64
	Offset int64
	Order  string
}

type Transfer struct {
	Id        string `json:"id,omitempty"`
	FromSubID string `json:"from_sub_id,omitempty"`
	ToSubID   string `json:"to_sub_id,omitempty"`
	Status    string `json:"status,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Amount    string `json:"amount,omitempty"`
	Currency  string `json:"currency,omitempty"`
}

// Transfer with return a list of all transfers based on supplied options (which can be nil)
// JWT is required for this request
func ListTransfers(o *ListTransfersOption) ([]*Transfer, error) {
	u := url.Values{}

	if o != nil {
		if o.Id != 0 {
			u.Set("id", fmt.Sprintf("%d", o.Id))
		}
		if o.Status != "" {
			u.Set("status", o.Status)
		}
		if o.Limit != 0 {
			u.Set("limit", fmt.Sprintf("%d", o.Limit))
		}
		if o.Offset != 0 {
			u.Set("offset", fmt.Sprintf("%d", o.Offset))
		}
		if o.Order != "" {
			u.Set("order", o.Order)
		}
	}

	tok, err := core.Authenticate(config.Login(), config.Password())
	if err != nil {
		return nil, eris.Wrap(err, "list")
	}

	type plist struct {
		Result []*Transfer `json:"result"`
	}

	pl := &plist{Result: make([]*Transfer, 0)}
	par := &core.SendParams{
		RouteName: "custody-list-transfers",
		Into:      pl,
		Values:    u,
		Token:     tok,
	}

	err = core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return pl.Result, nil
}
