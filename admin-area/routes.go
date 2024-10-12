package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", list)
	r.POST("/", create)
	r.PUT("/", update)
	r.GET("/:user", detail)
	r.DELETE("/:user", remove)
	r.PATCH("/:user", restore)
	return r
}
