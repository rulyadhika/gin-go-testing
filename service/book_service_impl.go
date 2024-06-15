package service

import (
	"gin-go-testing/model/domain"
	"gin-go-testing/model/dto"
	"gin-go-testing/repository"

	"github.com/gin-gonic/gin"
)

type bookServiceImpl struct {
	br repository.BookRepository
}

func NewBookServiceImpl(br repository.BookRepository) BookService {
	return &bookServiceImpl{br}
}

func (b *bookServiceImpl) Create(ctx *gin.Context, bookDto *dto.NewBookRequest) (*dto.BookResponse, error) {
	book := &domain.Book{Title: bookDto.Title, Author: bookDto.Author}

	result, err := b.br.Create(ctx, book)

	if err != nil {
		return nil, err
	}

	return &dto.BookResponse{Id: result.Id, Title: result.Title, Author: result.Author}, nil
}

func (b *bookServiceImpl) FindOneById(ctx *gin.Context, bookId uint) (*dto.BookResponse, error) {
	result, err := b.br.FindOneById(ctx, bookId)

	if err != nil {
		return nil, err
	}

	return &dto.BookResponse{Id: result.Id, Title: result.Title, Author: result.Author}, nil
}

func (b *bookServiceImpl) FindAll(ctx *gin.Context) ([]*dto.BookResponse, error) {
	result, err := b.br.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	booksDto := []*dto.BookResponse{}

	for _, e := range result {
		booksDto = append(booksDto, &dto.BookResponse{Id: e.Id, Title: e.Title, Author: e.Author})
	}

	return booksDto, nil
}
