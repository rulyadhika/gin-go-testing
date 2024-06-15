package repository

import (
	"gin-go-testing/model/domain"

	"github.com/gin-gonic/gin"
)

type BookRepository interface {
	Create(ctx *gin.Context, book *domain.Book) (*domain.Book, error)
	FindOneById(ctx *gin.Context, bookId uint) (*domain.Book, error)
	FindAll(ctx *gin.Context) ([]*domain.Book, error)
}
