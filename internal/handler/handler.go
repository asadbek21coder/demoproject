package handler

import (
	"github.com/asadbek21coder/demoproject/internal/logger"
	"github.com/asadbek21coder/demoproject/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
	log      *logger.Logger
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/ping", h.pong)

	books := router.Group("/books")
	{
		books.GET("/", h.getAllBooks)
		books.GET("/:id", h.getBookById)
		books.POST("/", h.createBook)
		books.PUT("/:id", h.updateBook)
		books.DELETE("/:id", h.deleteBook)
	}

	authors := router.Group("/authors")
	{
		authors.GET("/", h.getAllAuthors)
		authors.GET("/:id", h.getAuthorById)
		authors.POST("/", h.createAuthor)
		authors.PUT("/:id", h.updateAuthor)
		authors.DELETE("/:id", h.deleteAuthor)
	}

	return router
}
