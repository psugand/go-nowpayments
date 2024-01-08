package subscriptions

import (
	"fmt"
	"net/url"

	"github.com/CIDgravity/go-nowpayments/core"
)

// ListOption are options applying to the list of subscriptions
type ListOption struct {
	Limit  int
	Offset int
}

// List returns a list of all subscription plans, depending on the supplied options (which can be nil).
func List(o *ListOption) ([]*Subscription, error) {
	u := url.Values{}

	if o != nil {
		if o.Limit != 0 {
			u.Set("limit", fmt.Sprintf("%d", o.Limit))
		}
		if o.Offset != 0 {
			u.Set("offset", fmt.Sprintf("%d", o.Offset))
		}
	}

	type slist struct {
		Data []*Subscription `json:"data"`
	}

	pl := &slist{Data: make([]*Subscription, 0)}
	par := &core.SendParams{
		RouteName: "subscription-list",
		Into:      pl,
		Values:    u,
	}

	err := core.HTTPSend(par)
	if err != nil {
		return nil, err
	}

	return pl.Data, nil
}
