package service

import (
	"gin-go-testing/model/dto"

	"github.com/gin-gonic/gin"
)

type BookService interface {
	Create(ctx *gin.Context, bookDto *dto.NewBookRequest) (*dto.BookResponse, error)
	FindOneById(ctx *gin.Context, bookId uint) (*dto.BookResponse, error)
	FindAll(ctx *gin.Context) ([]*dto.BookResponse, error)
}
