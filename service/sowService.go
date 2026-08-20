package service

import (
	"context"
	"fmt"

	"github.com/GuilhermeW1/backend-suino/model"
	"github.com/GuilhermeW1/backend-suino/repository"
)

type SowService struct {
	R *repository.SowRepository
}

func (s *SowService) GetById(ctx context.Context, id uint) (*model.Sow, error) {
	sow, err := s.R.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return sow, nil
}

func (s *SowService) Create(ctx context.Context, sow *model.Sow) (*model.Sow, error) {
	err := s.R.Create(ctx, sow)
	if err != nil {
		return nil, fmt.Errorf("Service error, create sow: %w", err)
	}

	return sow, nil
}

func (s *SowService) GetAllActive(ctx context.Context) ([]model.Sow, error) {
	sow, err := s.R.GetAllActiveSows(ctx)
	if err != nil {
		return nil, fmt.Errorf("Service error getAllActive sows: %w", err)
	}

	return sow, nil
}

func (s *SowService) Delete(ctx context.Context, sowId uint) error {
	err := s.R.Delete(ctx, sowId)
	if err != nil {
		return fmt.Errorf("Service error delete: %w", err)
	}

	return nil
}
