package handler

import (
	"net/http"
	"strconv"

	"github.com/GuilhermeW1/backend-suino/model"
	"github.com/GuilhermeW1/backend-suino/service"
	"github.com/gin-gonic/gin"
)

type SowHandler struct {
	Service *service.SowService
}

func (h *SowHandler) Create(c *gin.Context) {
	var sow model.Sow

	if err := c.ShouldBindJSON(&sow); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados invalidos"})
		return
	}

	createdSow, err := h.Service.Create(c.Request.Context(), &sow)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdSow)
}

func (h *SowHandler) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid ID"})
	}

	sow, err := h.Service.GetById(ctx.Request.Context(), uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, sow)
}

func (h *SowHandler) GetByEarTag(ctx *gin.Context) {
	earTagParam := ctx.Param("earTag")

	sow, err := h.Service.GetByEarTag(ctx, earTagParam)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, sow)

}

func (h *SowHandler) GetAllActive(ctx *gin.Context) {
	sows, err := h.Service.GetAllActive(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, sows)
}

func (h *SowHandler) Delete(ctx *gin.Context) {
	sowId := ctx.Param("id")

	id, err := strconv.ParseInt(sowId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid ID"})
	}

	err = h.Service.Delete(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
