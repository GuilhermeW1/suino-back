package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/GuilhermeW1/backend-suino/model"
	"gorm.io/gorm"
)

type SowRepository struct {
	DB *gorm.DB
}

func (r *SowRepository) FindByEarTag(ctx context.Context, earTag string) (*model.Sow, error) {
	var sow model.Sow

	err := r.DB.WithContext(ctx).Where("ear_tag = ?", earTag).First(&sow).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Porca com o brinco %s nao encontrado: %w", earTag, err)
		}

		return nil, fmt.Errorf("Erro ao buscar porca com brinco %s: %w", earTag, err)
	}

	return &sow, nil
}

func (r *SowRepository) FindByID(ctx context.Context, id uint) (*model.Sow, error) {
	var sow model.Sow

	err := r.DB.WithContext(ctx).Where("id = ?", id).First(&sow).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Porca com o id %d nao encontrado: %w", id, err)
		}

		return nil, fmt.Errorf("Erro ao buscar porca com id: %d: %w", id, err)
	}

	return &sow, nil
}

func (r *SowRepository) GetAllActiveSows(ctx context.Context) ([]model.Sow, error) {
	var sows []model.Sow

	err := r.DB.WithContext(ctx).Where("status = ?", model.SowStatusActive).Find(&sows).Error
	if err != nil {
		return nil, fmt.Errorf("Erro ao buscar porcas ativas: %w", err)
	}

	return sows, nil
}

func (r *SowRepository) Create(ctx context.Context, sow *model.Sow) error {
	err := r.DB.WithContext(ctx).Create(sow).Error
	if err != nil {
		return fmt.Errorf("Erro ao inserir porca no banco %w", err)
	}

	return nil
}

func (r *SowRepository) Delete(ctx context.Context, sowId uint) error {
	err := r.DB.WithContext(ctx).Model(&model.Sow{}).Where("id = ?", sowId).Delete(sowId).Error
	if err != nil {
		return fmt.Errorf("Erro ao deletar porca: %w", err)
	}

	return nil
}
