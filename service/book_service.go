package service

import (
	"gin-go-testing/model/dto"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/go-custom-err/errs"
)

type BookService interface {
	Create(ctx *gin.Context, bookDto *dto.NewBookRequest) (*dto.BookResponse, errs.CustomError)
	FindOneById(ctx *gin.Context, bookId uint) (*dto.BookResponse, errs.CustomError)
	FindAll(ctx *gin.Context) ([]*dto.BookResponse, errs.CustomError)
}
