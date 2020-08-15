package domains

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessResponseWithData(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func SuccessResponseWithoutData(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"message": message})
}

func ErrorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"message": message})
}

func InternalServerError(c *gin.Context) {
	ErrorResponse(c, http.StatusInternalServerError, ErrInternalServerError.Error())
}

func NotFound(c *gin.Context) {
	ErrorResponse(c, http.StatusNotFound, ErrNotFound.Error())
}

func Conflict(c *gin.Context) {
	ErrorResponse(c, http.StatusConflict, ErrConflict.Error())
}

func BadRequest(c *gin.Context, errors interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{"message": ErrBadRequest.Error(), "errors": errors})
}

func UnprocessableEntity(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusUnprocessableEntity, message)
}

func Unauthenticated(c *gin.Context) {
	ErrorResponse(c, http.StatusUnauthorized, ErrUnauthenticate.Error())
}

func Unauthorized(c *gin.Context) {
	ErrorResponse(c, http.StatusForbidden, ErrUnauthorized.Error())
}
