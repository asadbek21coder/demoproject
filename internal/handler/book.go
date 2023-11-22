package handler

import (
	"net/http"
	"strconv"

	"github.com/asadbek21coder/demoproject/internal/entities"
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

func (h *Handler) getBookById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if HandleBadRequestErrWithMessage(c, h.log, err, "strconv.Atoi(c.Param(id))") {
		return
	}

	res, err := h.services.Books.GetBookById(id)
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.services.Books.GetBookById(id)") {
		return
	}

	c.JSON(http.StatusOK, entities.GetBookByIdRes{
		ErrorCode:    entities.ErrorSuccessCode,
		ErrorMessage: "",
		Body:         res,
	})
}

func (h *Handler) createBook(c *gin.Context) {
	body := &entities.CreateBookReq{}

	err := c.ShouldBindJSON(&body)
	if HandleBadRequestErrWithMessage(c, h.log, err, "c.ShouldBindJSON(&body)") {
		return
	}

	res, err := h.services.Books.CreateBook(&entities.CreateBookReq{
		Name: body.Name,
	})
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.services.Books.CreateBook(&entities.CreateBookReq") {
		return
	}

	c.JSON(201, entities.BookRes{
		ErrorCode:    entities.ErrorSuccessCode,
		ErrorMessage: "",
		Body:         res,
	})
}

func (h *Handler) updateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if HandleBadRequestErrWithMessage(c, h.log, err, "strconv.Atoi(c.Param(id))") {
		return
	}

	body := &entities.UpdateBookReq{}

	err = c.ShouldBindJSON(&body)
	if HandleBadRequestErrWithMessage(c, h.log, err, "c.ShouldBindJSON(&body)") {
		return
	}

	res, err := h.services.Books.UpdateBook(id, &entities.UpdateBookReq{
		Name: body.Name,
	})
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.services.Books.UpdateBook(id, &entities.UpdateBookReq") {
		return
	}

	c.JSON(201, entities.BookRes{
		ErrorCode:    entities.ErrorSuccessCode,
		ErrorMessage: "",
		Body:         res,
	})
}

func (h *Handler) deleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if HandleBadRequestErrWithMessage(c, h.log, err, "strconv.Atoi(c.Param(id))") {
		return
	}

	err = h.services.Books.DeleteBook(id)
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.services.Books.DeleteBook(id)") {
		return
	}

	c.JSON(201, entities.DeleteBookRes{
		ErrorCode:    entities.ErrorSuccessCode,
		ErrorMessage: "",
	})
}
