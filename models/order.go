package models

import "github.com/google/uuid"

type OrderType int

const (
	Buy OrderType = iota
	Sell
)

type Contract struct {
	ID        uuid.UUID `json:"id"`
	Issuer    uuid.UUID `json:"issuer"`
	Created   int64     `json:"created"`
	Expires   int64     `json:"expires"`
	Accepted  int64     `json:"accepted"`
	Closed    int64     `json:"closed"`
	Fulfilled func(args interface{}) bool
}

type MarketOrder struct {
	Issuer         uuid.UUID `json:"issuer"`
	Representative uuid.UUID `json:"representative"`
	MarketID       uuid.UUID `json:"market_id"`
	ItemID         uuid.UUID `json:"item_id"`
	OrderType      OrderType `json:"order_type"`
	Amount         int       `json:"amount"`
	Created        int64     `json:"created"`
	Accepted       int64     `json:"accepted"`
	Closed         int64     `json:"closed"`
	Fulfilled      func(args interface{}) bool
}
