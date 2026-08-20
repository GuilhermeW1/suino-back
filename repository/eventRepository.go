package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/GuilhermeW1/backend-suino/model"
	"gorm.io/gorm"
)

type EventRepository struct {
	DB *gorm.DB
}

func (r *EventRepository) FindAllEventsBySowId(ctx context.Context, sowID uint) ([]model.Event, error) {
	var events []model.Event

	err := r.DB.WithContext(ctx).Where("sow_id = ?", sowID).Find(&events).Error
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (r *EventRepository) GetAllEvents(ctx context.Context) ([]model.Event, error) {
	var events []model.Event

	err := r.DB.WithContext(ctx).Model(&model.Event{}).Find(&events).Error
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
