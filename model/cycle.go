package model

import (
	"time"

	"gorm.io/gorm"
)

type CycleStatus string

const (
	CycleStatusActive CycleStatus = "ACTIVE"
	CycleStatusClosed CycleStatus = "CLOSED"
)

type Cycle struct {
	ID          uint64      `gorm:"primaryKey;autoIncrement" json:"id"`
	SowID       uint64      `gorm:"not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"sowId"`
	CycleNumber int         `gorm:"not null" json:"cycle_number"`
	Status      CycleStatus `gorm:"type:varchar(20);not null;default:'ACTIVE'" json:"status"`
	StartDate   time.Time   `gorm:"not null" json:"start_date"`
	EndDate     *time.Time  `json:"end_date,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Sow    *Sow    `gorm:"foreignKey:SowID" json:"sow,omitempty"`
	Events []Event `gorm:"foreignKey:CycleID" json:"events,omitempty"`
}

func (c *Cycle) CanStartNewCycle() bool {
	return c == nil || c.Status != CycleStatusActive
}

func (c *Cycle) GetNextCycleNumber() int {
	if c == nil {
		return 1
	}

	return c.CycleNumber + 1
}
