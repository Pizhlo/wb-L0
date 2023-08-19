package models

import "database/sql"

type Payment struct {
	ID           int
	Transaction  string         `json:"transaction"`
	RequestID    sql.NullString `json:"request_id"`
	Currency     string         `json:"currency"`
	Provider     string         `json:"provider"`
	Amount       int            `json:"amount"`
	PaymentDate  string         `json:"payment_dt"`
	Bank         string         `json:"bank"`
	DeliveryCost int            `json:"delivery_cost"`
	GoodsTotal   int            `json:"goods_total"`
	CustomFee    int            `json:"custom_fee"`
}
