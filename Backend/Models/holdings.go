package models

import (
	"time"

	"github.com/google/uuid"
)

type Holding struct {
	ID           uuid.UUID `json:"id"`
	PortfolioID  uuid.UUID `json:"portfolio_id"`

	Symbol       string    `json:"symbol"`
	Quantity     float64   `json:"quantity"`
	PurchasePrice float64   `json:"purchase_price"`

	PurchaseDate time.Time `json:"purchase_date"`

	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}