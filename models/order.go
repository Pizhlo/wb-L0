package models

import (
	"database/sql"
	"time"

	uuid "github.com/google/uuid"
)

type Order struct {
	ID                int
	OrderUIID         uuid.UUID      `json:"order_uid" db:"order_id"`
	TrackNumber       string         `json:"track_number"`
	Entry             string         `json:"entry"`
	Delivery          Delivery       `json:"delivery"`
	Payment           Payment        `json:"payment"`
	Items             []Item         `json:"items"`
	Locale            string         `json:"locale"`
	InternalSignature sql.NullString `json:"internal_signature"`
	CustomerID        string         `json:"customer_id"`
	DeliveryService   string         `json:"delivery_service"`
	ShardKey          string         `json:"shard_key"`
	SmID              int            `json:"sm_id"`
	DateCreated       time.Time      `json:"date_created"`
	OofShard          string         `json:"oof_shard"`
}
