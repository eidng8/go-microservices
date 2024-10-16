package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter(env *Env) *gin.Engine {
	r := gin.Default()
	r.GET(
		"/",
		func(c *gin.Context) {
			list(c, env)
		},
	)
	r.POST("/", create)
	r.PUT("/", update)
	r.GET("/:user", detail)
	r.DELETE("/:user", remove)
	r.PATCH("/:user", restore)
	return r
}
