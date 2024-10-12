package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RespondJSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func RespondWithCodeJSON(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}

func ErrorJSON(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func ErrorWithCodeJSON(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{"error": err.Error()})
}
