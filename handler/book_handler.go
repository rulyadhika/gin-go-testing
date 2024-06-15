package handler

import "github.com/gin-gonic/gin"

type BookHandler interface {
	Create(ctx *gin.Context)
	FindOneById(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}
