package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/GuilhermeW1/backend-suino/model"
	"gorm.io/gorm"
)

type CycleRepository struct {
	DB *gorm.DB
}

func (r *CycleRepository) CreateCycle(ctx context.Context, sowId uint64) (*model.Cycle, error) {
	prevCycle, _ := r.GetCycleBySowId(ctx, sowId)

	// cycles start at 1 by default
	cycleCount := 1

	if prevCycle != nil {
		cycleCount = prevCycle.GetNextCycleNumber()
	}

	newCycle := &model.Cycle{
		SowID:       uint64(sowId),
		CycleNumber: cycleCount,
		Status:      model.CycleStatusActive,
		StartDate:   time.Now(),
		EndDate:     nil,
		CreatedAt:   time.Now(),
	}

	err := r.DB.WithContext(ctx).Create(newCycle).Error
	if err != nil {
		return nil, fmt.Errorf("Error creating new Cyclce: %s", err)
	}

	return newCycle, nil
}

func (r *CycleRepository) GetCycleBySowId(ctx context.Context, sowId uint64) (*model.Cycle, error) {
	var cycle model.Cycle

	err := r.DB.WithContext(ctx).Where("sow_id = ?", sowId).First(&cycle).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &cycle, nil
}

func (r *CycleRepository) GetActiveCycleBySowId(ctx context.Context, sowId uint64) (*model.Cycle, error) {
	var cycle model.Cycle

	err := r.DB.WithContext(ctx).Where("sow_id = ? AND status = ?", sowId, model.CycleStatusActive).First(&cycle).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &cycle, nil
}

func (r *CycleRepository) GetByID(ctx context.Context, cycleId uint64) (*model.Cycle, error) {
	var cycle model.Cycle

	err := r.DB.WithContext(ctx).Where("id = ?", cycleId).First(&cycle).Error
	if err != nil {
		return nil, fmt.Errorf("Cycle repository: Error getting cycle: %s", err)
	}

	return &cycle, nil
}

func (r *CycleRepository) CloseCycle(ctx context.Context, cycleId uint64) error {
	err := r.DB.WithContext(ctx).
		Model(model.Cycle{}).
		Where("id = ?", cycleId).
		Update("status", model.CycleStatusClosed).
		Error

	if err != nil {
		return fmt.Errorf("Repository error closing cycle: %s", err)
	}

	return nil
}
