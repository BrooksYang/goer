package models

import (
	"time"
)

type BaseModel struct {
	ID uint64 `gorm:"primaryKey" json:"id"`
}

type TimestampsField struct {
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
