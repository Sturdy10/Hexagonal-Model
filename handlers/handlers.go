package handlers

import (
	"Hexagonal-Model/models"
	"Hexagonal-Model/services"

	"github.com/gin-gonic/gin"
)

type IHandlerPort interface {
	PostRegisterHandler(c *gin.Context)
}

type handlerAdapter struct {
	s services.ServicesPort
}

func NewHandlers(s services.ServicesPort) IHandlerPort {
	return &handlerAdapter{s: s}
}

func (h *handlerAdapter) PostRegisterHandler(c *gin.Context) {
	var req models.RequestRegister

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"status": "error", "message": err.Error()})
		return
	}
	err := h.s.PotsRegisterSer(req)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "Failed to register user"})
		return
	}

	c.JSON(200, gin.H{"status": "OK", "message": "User registered successfully"})
}
