package service

import (
	"context"
	"fmt"
	"time"

	"github.com/GuilhermeW1/backend-suino/model"
	"github.com/GuilhermeW1/backend-suino/repository"
	"github.com/GuilhermeW1/backend-suino/service/dto"
)

type EventService struct {
	R            *repository.EventRepository
	CycleService *CycleService
}

// defines 14 days to get a return
const maxGapSameCycle = 14

func (s *EventService) AddEvent(ctx context.Context, event *model.Event) (*model.Event, error) {
	activeCycle, err := s.CycleService.GetActiveCycleBySowId(ctx, event.SowID)
	if err != nil {
		return nil, fmt.Errorf("service error: getting active cycle: %w", err)
	}

	resolvedCycle, err := s.resolveCycle(ctx, event, activeCycle)
	if err != nil {
		return nil, fmt.Errorf("Service error: error closing cycle: %w", err)
	}

	if resolvedCycle != nil {
		event.Cycle = resolvedCycle
		event.CycleID = &resolvedCycle.ID
	}

	newEvent, err := s.R.CreateEvent(ctx, event)
	if err != nil {
		return nil, fmt.Errorf("service error: creating event: %w", err)
	}

	return newEvent, nil
}

func (s *EventService) GetEventsBySowId(ctx context.Context, sowId uint) ([]dto.EventResponseDto, error) {
	events, err := s.R.GetEventsBySowId(ctx, sowId)
	if err != nil {
		return nil, fmt.Errorf("service error: getting event: %w", err)
	}

	return events, nil
}

// TODO: fazer
func (s *EventService) GetEvents(ctx context.Context) ([]dto.EventResponseDto, error) {
	events, err := s.R.GetAllEvents(ctx)
	if err != nil {
		return nil, fmt.Errorf("Service error: error getting events: %w", err)
	}

	return events, err
}

func (s *EventService) resolveCycle(ctx context.Context, event *model.Event, activeCycle *model.Cycle) (*model.Cycle, error) {
	switch event.Type {

	case model.EventTypeAbortion:
		if activeCycle == nil {
			return nil, fmt.Errorf("service error: abortion event without an activeCycle cycle")
		}

		if err := s.CycleService.DisableCycle(ctx, activeCycle.ID); err != nil {
			return nil, fmt.Errorf("service error: disabling cycle: %w", err)
		}

		return activeCycle, nil

	case model.EventTypeInsemination:
		if activeCycle == nil {

			return s.CycleService.CreateCycle(ctx, event.SowID)
		}

		closes, err := s.IsInseminationReturn(ctx, activeCycle, event.EventDate)
		if err != nil {
			return nil, fmt.Errorf("service error: checking return window: %w", err)
		}

		if closes {
			if err := s.CycleService.DisableCycle(ctx, activeCycle.ID); err != nil {
				return nil, fmt.Errorf("service error: disabling cycle: %w", err)
			}

			return s.CycleService.CreateCycle(ctx, event.SowID)
		}

		return activeCycle, nil

	default:
		return activeCycle, nil
	}

}

func (s *EventService) IsInseminationReturn(ctx context.Context, cycle *model.Cycle, newEventDate time.Time) (bool, error) {
	lastInsemination, err := s.R.GetLastEventByCycleAndType(ctx, cycle.ID, model.EventTypeInsemination)
	if err != nil {
		return false, fmt.Errorf("Service error: error getting last event")
	}

	if lastInsemination == nil {
		return false, nil
	}

	gap := newEventDate.Sub(lastInsemination.EventDate)
	return gap > maxGapSameCycle, nil
}
