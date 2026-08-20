package model

import (
	"time"

	"gorm.io/gorm"
)

type SowStatus string

const (
	SowStatusActive  SowStatus = "ACTIVE"
	SowStatusDisable SowStatus = "DISABLE"
)

type Sow struct {
	ID     uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	EarTag string    `gorm:"type:varchar(50);not null" json:"earTag"`
	Status SowStatus `gorm:"type:varchar(30);default:'ACTIVE'" json:"status"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Events []Event `gorm:"foreignKey:SowID" json:"events,omitempty"`
	Cycles []Cycle `gorm:"foreignKey:SowID" json:"cycles,omitempty"`
}
