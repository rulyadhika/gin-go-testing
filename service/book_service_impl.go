package service

import (
	"database/sql"
	"gin-go-testing/model/domain"
	"gin-go-testing/model/dto"
	"gin-go-testing/repository"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/go-custom-err/errs"
)

type bookServiceImpl struct {
	br repository.BookRepository
	db *sql.DB
}

func NewBookServiceImpl(br repository.BookRepository, db *sql.DB) BookService {
	return &bookServiceImpl{br, db}
}

func (b *bookServiceImpl) Create(ctx *gin.Context, bookDto *dto.NewBookRequest) (*dto.BookResponse, errs.CustomError) {
	book := &domain.Book{Title: bookDto.Title, Author: bookDto.Author}

	result, err := b.br.Create(ctx, b.db, book)

	if err != nil {
		return nil, err
	}

	return &dto.BookResponse{Id: result.Id, Title: result.Title, Author: result.Author}, nil
}

func (b *bookServiceImpl) FindOneById(ctx *gin.Context, bookId uint) (*dto.BookResponse, errs.CustomError) {
	result, err := b.br.FindOneById(ctx, b.db, bookId)

	if err != nil {
		return nil, err
	}

	return &dto.BookResponse{Id: result.Id, Title: result.Title, Author: result.Author}, nil
}

func (b *bookServiceImpl) FindAll(ctx *gin.Context) ([]*dto.BookResponse, errs.CustomError) {
	result, err := b.br.FindAll(ctx, b.db)

	if err != nil {
		return nil, err
	}

	booksDto := []*dto.BookResponse{}

	for _, e := range result {
		booksDto = append(booksDto, &dto.BookResponse{Id: e.Id, Title: e.Title, Author: e.Author})
	}

	return booksDto, nil
}
