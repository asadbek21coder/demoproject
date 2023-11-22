package handler

import (
	"net/http"
	"strconv"

	"github.com/asadbek21coder/demoproject/internal/entities"
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

func (h *Handler) getAuthorById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if HandleBadRequestErrWithMessage(c, h.log, err, "strconv.Atoi(c.Param(id))") {
		return
	}

	res, err := h.services.Authors.GetAuthorById(id)
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.services.Authors.GetAuthorById(id)") {
		return
	}
	// if err != nil {
	// 	log.Print("error : \n", err.Error())
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, entities.Response{
	// 		ErrorCode:    entities.ErrorCodeInternal,
	// 		ErrorMessage: "Ooops something went wrong",
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, entities.GetAuthorByIdRes{
		ErrorCode:    entities.ErrorSuccessCode,
		ErrorMessage: "",
		Body:         res,
	})
}

func (h *Handler) createAuthor(c *gin.Context) {
	body := &entities.CreateAuthorReq{}

	err := c.ShouldBindJSON(&body)
	if HandleBadRequestErrWithMessage(c, h.log, err, "c.ShouldBindJSON(&body)") {
		return
	}

	res, err := h.services.Authors.CreateAuthor(&entities.CreateAuthorReq{
		Name: body.Name,
	})
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.services.Authors.CreateAuthor(&entities.CreateAuthorReq") {
		return
	}

	c.JSON(201, entities.AuthorRes{
		ErrorCode:    entities.ErrorSuccessCode,
		ErrorMessage: "",
		Body:         res,
	})
}

func (h *Handler) updateAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if HandleBadRequestErrWithMessage(c, h.log, err, "strconv.Atoi(c.Param(id))") {
		return
	}

	body := &entities.UpdateAuthorReq{}

	err = c.ShouldBindJSON(&body)
	if HandleBadRequestErrWithMessage(c, h.log, err, "c.ShouldBindJSON(&body)") {
		return
	}

	res, err := h.services.Authors.UpdateAuthor(id, &entities.UpdateAuthorReq{
		Name: body.Name,
	})
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.services.Authors.UpdateAuthor(id, &entities.UpdateAuthorReq") {
		return
	}

	c.JSON(201, entities.AuthorRes{
		ErrorCode:    entities.ErrorSuccessCode,
		ErrorMessage: "",
		Body:         res,
	})
}

func (h *Handler) deleteAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if HandleBadRequestErrWithMessage(c, h.log, err, "strconv.Atoi(c.Param(id))") {
		return
	}

	err = h.services.Authors.DeleteAuthor(id)
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.services.Authors.DeleteAuthor(id)") {
		return
	}

	c.JSON(201, entities.DeleteAuthorRes{
		ErrorCode:    entities.ErrorSuccessCode,
		ErrorMessage: "",
	})
}
