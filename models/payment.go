package models

import (
	"database/sql"

	uuid "github.com/google/uuid"
)

type Payment struct {
	ID           int
	Transaction  uuid.UUID      `json:"transaction"`
	RequestID    sql.NullString `json:"request_id"`
	Currency     string         `json:"currency"`
	Provider     string         `json:"provider"`
	Amount       int            `json:"amount"`
	PaymentDate  int64          `json:"payment_dt"`
	Bank         string         `json:"bank"`
	DeliveryCost int            `json:"delivery_cost"`
	GoodsTotal   int            `json:"goods_total"`
	CustomFee    int            `json:"custom_fee"`
}
