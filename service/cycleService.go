package service

import (
	"context"
	"fmt"

	"github.com/GuilhermeW1/backend-suino/model"
	"github.com/GuilhermeW1/backend-suino/repository"
)

type CycleService struct {
	R *repository.CycleRepository
}

func (c *CycleService) DisableCycle(ctx context.Context, cycleId uint64) error {
	_, err := c.R.GetByID(ctx, cycleId)
	if err != nil {
		return err
	}

	return c.R.CloseCycle(ctx, cycleId)
}

func (c CycleService) CreateCycle(ctx context.Context, sowId uint64) (*model.Cycle, error) {
	cycle, err := c.R.CreateCycle(ctx, sowId)
	if err != nil {
		return nil, fmt.Errorf("Cycle service error creating cycle: %w", err)
	}

	return cycle, nil
}

func (c CycleService) GetCycleBySowId(ctx context.Context, sowId uint64) (*model.Cycle, error) {
	cycle, err := c.R.GetCycleBySowId(ctx, sowId)
	if err != nil {
		return nil, fmt.Errorf("Cycle service error getting cycle: %w", err)
	}

	return cycle, nil
}

func (c CycleService) GetActiveCycleBySowId(ctx context.Context, sowId uint64) (*model.Cycle, error) {
	cycle, err := c.R.GetActiveCycleBySowId(ctx, sowId)
	if err != nil {
		return nil, fmt.Errorf("Cycle service error getting cycle: %w", err)
	}

	return cycle, nil
}
