package models

import (
	"time"
)

type Company struct {
	Symbol    string    `json:"symbol"`
	Name      string    `json:"name"`
	Exchange  string    `json:"exchange"`
	Sector    string    `json:"sector"`
	Industry  string    `json:"industry"`
	Country   string    `json:"country"`
	Currency  string    `json:"currency"`
	UpdatedAt time.Time `json:"updated_at"`
}
