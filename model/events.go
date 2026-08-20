package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type EventType string

const (
	EventTypeInsemination   EventType = "INSEMINATION"
	EventTypePregnancyCheck EventType = "PREGNANCY_CHECK"
	EventTypeFarrowing      EventType = "FARROWING"
	EventTypeWeaning        EventType = "WEANING"
	EventTypeAbortion       EventType = "ABORTION"
	EventTypeMedication     EventType = "MEDICATION"
)

type Event struct {
	ID          uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	Type        EventType      `gorm:"type:text;not null;index:idx_events_type_date" json:"type"`
	SowID       uint64         `gorm:"not null;index:idx_events_sow_date;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"sowId"`
	CycleID     *uint64        `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"cycle_id,omitempty"`
	RefEventID  *uint64        `json:"ref_event_id,omitempty"`
	Description *string        `gorm:"type:text" json:"description,omitempty"`
	Payload     datatypes.JSON `gorm:"type:jsonb" json:"payload,omitempty"`
	EventDate   time.Time      `gorm:"not null;index:idx_events_sow_date;index:idx_events_type_date" json:"event_date"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Sow      *Sow   `gorm:"foreignKey:SowID" json:"sow,omitempty"`
	Cycle    *Cycle `gorm:"foreignKey:CycleID" json:"cycle,omitempty"`
	RefEvent *Event `gorm:"foreignKey:RefEventID" json:"ref_event,omitempty"`
}

func (e *Event) CheckIsEventCloseCycle() bool {
	if e.Type == EventTypeAbortion || e.Type == EventTypeInsemination {
		return true
	}

	return false
}

func (e *Event) IsInseminationEventCloseCycle() {

}
