package models

import (
	"time"

	"github.com/google/uuid"
)

type PortfolioSnapshot struct {
	ID               uuid.UUID `json:"id"`
	PortfolioID      uuid.UUID `json:"portfolio_id"`
	TotalInvestment  float64   `json:"total_investment"`
	InitialValue     float64   `json:"initial_value"`
	CurrentValue     float64   `json:"current_value"`
	ProfitLoss       float64   `json:"profit_loss"`
	ReturnPercentage float64   `json:"return_percentage"`
	SnapshotDate     time.Time `json:"snapshot_date"`
}
