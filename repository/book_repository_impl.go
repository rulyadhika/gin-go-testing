package repository

import (
	"database/sql"
	"errors"
	"gin-go-testing/model/domain"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/go-custom-err/errs"
)

type bookRepositoryImpl struct{}

func NewBookRepositoryImpl() BookRepository {
	return &bookRepositoryImpl{}
}
func (b *bookRepositoryImpl) Create(ctx *gin.Context, db *sql.DB, book *domain.Book) (*domain.Book, errs.CustomError) {
	err := db.QueryRowContext(ctx, createQuery, book.Title, book.Author).Scan(&book.Id)

	if err != nil {
		log.Printf("[CreateBook - Repo] err: %s", err.Error())
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return book, nil
}

func (b *bookRepositoryImpl) FindOneById(ctx *gin.Context, db *sql.DB, bookId uint) (*domain.Book, errs.CustomError) {
	book := new(domain.Book)

	err := db.QueryRowContext(ctx, findOneByIdQuery, bookId).Scan(&book.Id, &book.Title, &book.Author)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError("data not found")
		}

		log.Printf("[FindOneBookById - Repo] err: %s", err.Error())
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return book, nil
}

func (b *bookRepositoryImpl) FindAll(ctx *gin.Context, db *sql.DB) ([]*domain.Book, errs.CustomError) {
	books := []*domain.Book{}

	rows, err := db.QueryContext(ctx, findOneByIdQuery)
	if err != nil {
		log.Printf("[FindAllBook - Repo] err: %s", err.Error())

		return nil, errs.NewInternalServerError("something went wrong")
	}

	for rows.Next() {
		book := &domain.Book{}

		if err := rows.Scan(&book.Id, &book.Title, &book.Author); err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}

		books = append(books, book)
	}

	// if the result is empty
	if len(books) == 0 {
		return nil, errs.NewNotFoundError("not data found")
	}

	return books, nil
}
