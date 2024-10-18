package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter(env *Env) *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) { list(c, env) })
	r.POST("/", func(c *gin.Context) { create(c, env) })
	r.PUT("/", func(c *gin.Context) { update(c, env) })
	r.GET("/:user", func(c *gin.Context) { detail(c, env) })
	r.DELETE("/:user", func(c *gin.Context) { remove(c, env) })
	r.PATCH("/:user", func(c *gin.Context) { restore(c, env) })
	return r
}
