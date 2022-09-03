package controllers

import "github.com/gin-gonic/gin"

type Controller interface {
	Handle(ctx *gin.Context)
}
