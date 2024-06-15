package repository

import (
	"gin-go-testing/model/domain"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/go-custom-err/errs"
)

type BookRepository interface {
	Create(ctx *gin.Context, book *domain.Book) (*domain.Book, errs.CustomError)
	FindOneById(ctx *gin.Context, bookId uint) (*domain.Book, errs.CustomError)
	FindAll(ctx *gin.Context) ([]*domain.Book, errs.CustomError)
}
