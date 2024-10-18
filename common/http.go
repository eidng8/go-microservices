package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RespondJSON(c *gin.Context, data interface{}) {
	RespondWithCodeJSON(c, http.StatusOK, data)
}

func RespondWithCodeJSON(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}

func ErrorJSON(c *gin.Context, err error) {
	ErrorWithCodeJSON(c, http.StatusInternalServerError, err)
}

func Error400JSON(c *gin.Context, err error) {
	ErrorWithCodeJSON(c, http.StatusBadRequest, err)
}

func Error401JSON(c *gin.Context, err error) {
	ErrorWithCodeJSON(c, http.StatusUnauthorized, err)
}

func Error403JSON(c *gin.Context, err error) {
	ErrorWithCodeJSON(c, http.StatusForbidden, err)
}

func Error422JSON(c *gin.Context, err error) {
	ErrorWithCodeJSON(c, http.StatusUnprocessableEntity, err)
}

func ErrorWithCodeJSON(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{"error": err.Error()})
}
