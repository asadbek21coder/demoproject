package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) pong(c *gin.Context) {

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "pong",
	})
}
