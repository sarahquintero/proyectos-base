package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"ms-audios/dto"
	"ms-audios/repositories"
	"ms-audios/services"
)

type AudioController struct {
	service *services.AudioService
}

func NewAudioController(service *services.AudioService) *AudioController {
	return &AudioController{service: service}
}

func (c *AudioController) RegistrarRutas(r gin.IRouter) {
	r.GET("/audios", c.Listar)
	r.POST("/audios", c.Registrar)
}

func (c *AudioController) Listar(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.service.Listar())
}

func (c *AudioController) Registrar(ctx *gin.Context) {
	var d dto.AudioDTO
	if err := ctx.ShouldBindJSON(&d); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	creado, err := c.service.Registrar(d)
	if err != nil {
		if errors.Is(err, repositories.ErrAudioDuplicado) {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, creado)
}