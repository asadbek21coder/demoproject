package handler

import (
	"net/http"
	"strconv"

	"github.com/asadbek21coder/demoproject/internal/entities"
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
	}

	authors := router.Group("/authors")
	{
		authors.GET("/", h.getAllAuthors)
	}

	return router
}

func HandleInternalWithMessage(c *gin.Context, l *logger.Logger, err error, message string, args ...interface{}) bool {
	if err != nil {
		l.Error(message, err, args)
		c.AbortWithStatusJSON(http.StatusInternalServerError, entities.Response{
			ErrorCode:    entities.ErrorCodeInternal,
			ErrorMessage: "Oops something went wrong",
		})
		return true
	}

	return false
}

// func HandleDatabaseLevelWithMessage(c *gin.Context, l *logger.Logger, err error, message string, args ...interface{}) bool {
// 	status_err, _ := status.FromError(err)
// 	if err != nil {
// 		errorCode := entities.ErrorCodeInternal
// 		statuscode := http.StatusInternalServerError
// 		message := status_err.Message()
// 		switch status_err.Code() {
// 		case codes.NotFound:
// 			errorCode = entities.ErrorCodeNotFound
// 			statuscode = http.StatusNotFound
// 		case codes.Unknown:
// 			errorCode = entities.ErrorCodeBadRequest
// 			statuscode = http.StatusBadRequest
// 			message = "Ooops something went wrong"
// 		case codes.Aborted:
// 			errorCode = entities.ErrorCodeBadRequest
// 			statuscode = http.StatusBadRequest
// 		case codes.InvalidArgument:
// 			errorCode = entities.ErrorCodeBadRequest
// 			statuscode = http.StatusBadRequest
// 		}
// 		l.Error(message, err, args)
// 		c.AbortWithStatusJSON(statuscode, entities.Response{
// 			ErrorCode:    errorCode,
// 			ErrorMessage: message,
// 		})
// 		return true
// 	}
// 	return false
// }

func HandleBadRequestErrWithMessage(c *gin.Context, l *logger.Logger, err error, message string, args ...interface{}) bool {
	if err != nil {
		l.Error(message, err, args)
		c.AbortWithStatusJSON(http.StatusBadRequest, entities.Response{
			ErrorCode:    entities.ErrorCodeBadRequest,
			ErrorMessage: "Please enter right information",
		})
		return true
	}
	return false
}
func (h *Handler) ParseLimitQueryParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.DefaultQuery("limit", "10"))
}

func (h *Handler) ParsePageQueryParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.DefaultQuery("page", "1"))
}
