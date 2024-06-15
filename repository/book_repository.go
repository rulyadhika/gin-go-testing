package repository

import (
	"database/sql"
	"gin-go-testing/model/domain"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/go-custom-err/errs"
)

type BookRepository interface {
	Create(ctx *gin.Context, db *sql.DB, book *domain.Book) (*domain.Book, errs.CustomError)
	FindOneById(ctx *gin.Context, db *sql.DB, bookId uint) (*domain.Book, errs.CustomError)
	FindAll(ctx *gin.Context, db *sql.DB) ([]*domain.Book, errs.CustomError)
}
