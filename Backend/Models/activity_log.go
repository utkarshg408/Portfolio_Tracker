package models

import (
	"time"

	"github.com/google/uuid"
)

type ActivityLog struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"userid"`
	Action      string    `json:"action"`
	Symbol      string    `json:"symbol"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
