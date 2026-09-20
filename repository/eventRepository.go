package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/GuilhermeW1/backend-suino/model"
	"github.com/GuilhermeW1/backend-suino/service/dto"
	"gorm.io/gorm"
)

type EventRepository struct {
	DB *gorm.DB
}

func (r *EventRepository) GetEventsBySowId(ctx context.Context, sowID uint) ([]dto.EventResponseDto, error) {
	events := make([]dto.EventResponseDto, 0)

	err := r.DB.WithContext(ctx).
		Model(&model.Event{}).
		Select("events.*, sows.ear_tag AS sow_ear_tag").
		Joins("INNER JOIN sows ON sows.id = events.sow_id").
		Where("events.sow_id = ? AND events.deleted_at IS NULL", sowID).
		Order("events.event_date DESC").
		Scan(&events).
		Error

	if err != nil {
		return nil, err
	}

	return events, nil
}

func (r *EventRepository) GetAllEvents(ctx context.Context) ([]dto.EventResponseDto, error) {
	events := make([]dto.EventResponseDto, 0)

	err := r.DB.WithContext(ctx).Debug().Model(&model.Event{}).
		Select("events.*, sows.ear_tag AS sow_ear_tag").
		Joins("INNER JOIN sows ON sows.id = events.sow_id").
		Where("events.deleted_at IS NULL").
		Order("events.event_date DESC").
		Scan(&events).
		Error

	if err != nil {
		return nil, err
	}

	return events, nil

}

func (r *EventRepository) CreateEvent(ctx context.Context, event *model.Event) (*model.Event, error) {
	err := r.DB.WithContext(ctx).Create(event).Error
	if err != nil {
		return nil, fmt.Errorf("Repository error: error creating event, %w", err)
	}

	return event, nil
}

func (r *EventRepository) GetLastEventByCycleAndType(ctx context.Context, cycleId uint64, eventTye model.EventType) (*model.Event, error) {
	var event model.Event

	err := r.DB.WithContext(ctx).Where("cycle_id = ? AND type = ?", cycleId, eventTye).Find(&event).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("Repository error: error getting cycle by cycle id and type: %w", err)
	}

	return &event, nil
}
