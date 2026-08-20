package handler

import (
	"net/http"

	"github.com/GuilhermeW1/backend-suino/model"
	"github.com/GuilhermeW1/backend-suino/service"
	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	S *service.EventService
}

func (h *EventHandler) AddEvent(ctx *gin.Context) {
	var event model.Event

	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados invalidos"})
		return
	}

	createdEvent, err := h.S.AddEvent(ctx, &event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, createdEvent)
}

func (h *EventHandler) GetEvents(ctx *gin.Context) {
	events, err := h.S.GetEvents(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, events)
}

func (h *EventHandler) EditEvent(ctx *gin.Context) {

}
