package dto

import (
	"github.com/GuilhermeW1/backend-suino/model"
)

type EventResponseDto struct {
	model.Event `gorm:"embedded"`
	SowEarTag   string `json:"sowEarTag"`
}
