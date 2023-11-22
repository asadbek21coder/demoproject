package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllAuthors(c *gin.Context) {
	limit, err := h.ParseLimitQueryParam(c)
	if HandleBadRequestErrWithMessage(c, h.log, err, "h.ParseLimitQueryParam(c)") {
		return
	}
	page, err := h.ParsePageQueryParam(c)
	if HandleBadRequestErrWithMessage(c, h.log, err, "h.ParsePageQueryParam(c)") {
		return
	}

	authors, err := h.services.Authors.GetAllAuthors(page, limit)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"authors": authors,
	})
}
