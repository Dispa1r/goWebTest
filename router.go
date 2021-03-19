package main

import (
	"github.com/gin-gonic/gin"
	"goWeb/controller"
)

func collectRoute(r *gin.Engine) *gin.Engine{
	r.POST("/api/auth/register",controller.Register)
	return r

}