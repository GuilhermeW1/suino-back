package dto

import (
	"time"

	"github.com/GuilhermeW1/backend-suino/model"
	"gorm.io/datatypes"
)

type EventReqDto struct {
	Type        model.EventType `gorm:"type:text;not null;index:idx_events_type_date" json:"type"`
	SowID       uint64          `gorm:"not null;index:idx_events_sow_date;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"sow_id"`
	CycleID     *uint64         `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"cycle_id,omitempty"`
	RefEventID  *uint64         `json:"ref_event_id,omitempty"`
	Description *string         `gorm:"type:text" json:"description,omitempty"`
	Payload     datatypes.JSON  `gorm:"type:jsonb" json:"payload,omitempty"`
	EventDate   time.Time       `gorm:"not null;index:idx_events_sow_date;index:idx_events_type_date" json:"event_date"`
}
