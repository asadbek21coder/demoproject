package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllBooks(c *gin.Context) {
	limit, err := h.ParseLimitQueryParam(c)
	if HandleBadRequestErrWithMessage(c, h.log, err, "h.ParseLimitQueryParam(c)") {
		return
	}
	page, err := h.ParsePageQueryParam(c)
	if HandleBadRequestErrWithMessage(c, h.log, err, "h.ParsePageQueryParam(c)") {
		return
	}

	books, err := h.services.Books.GetAllBooks(page, limit)
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.services.Books.GetAllBooks") {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"books": books,
	})
}
