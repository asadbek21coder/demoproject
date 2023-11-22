package handler

import (
	"github.com/asadbek21coder/demoproject/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/", h.pong)

	books := router.Group("/books")
	{
		books.GET("/", h.getAllBooks)
	}

	return router
}
