package service

import (
	"errors"
	"gin-go-testing/mocks"
	"gin-go-testing/model/domain"
	"gin-go-testing/model/dto"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type unitTestBookServiceSuite struct {
	suite.Suite
	ctx *gin.Context
	brm *mocks.BookRepository
	bs  BookService
}

func TestUnitTestBookService(t *testing.T) {
	suite.Run(t, &unitTestBookServiceSuite{})
}

func (u *unitTestBookServiceSuite) SetupTest() {
	bookRepoMock := mocks.NewBookRepository(u.T())

	u.brm = bookRepoMock
	u.bs = NewBookServiceImpl(bookRepoMock)

	u.ctx = &gin.Context{}
}

func (u *unitTestBookServiceSuite) TestCreate_Success() {
	data := &domain.Book{Id: 2, Title: "The 7 Habits of Highly Effective People", Author: "Stephen R. Covey"}
	reqDto := &dto.NewBookRequest{Title: data.Title, Author: data.Author}
	expected := &dto.BookResponse{Id: data.Id, Title: data.Title, Author: data.Author}

	u.brm.On("Create", u.ctx, mock.Anything).Return(data, nil)

	result, err := u.bs.Create(u.ctx, reqDto)
	u.NoError(err)
	u.NotNil(result)
	u.Equal(expected, result)

	u.brm.AssertExpectations(u.T())
}

func (u *unitTestBookServiceSuite) TestCreate_Failed() {
	data := &domain.Book{Id: 2, Title: "The 7 Habits of Highly Effective People", Author: "Stephen R. Covey"}
	reqDto := &dto.NewBookRequest{Title: data.Title, Author: data.Author}

	u.brm.On("Create", u.ctx, mock.Anything).Return(nil, errors.New("something went wrong"))

	result, err := u.bs.Create(u.ctx, reqDto)
	u.Error(err)
	u.Nil(result)

	u.brm.AssertExpectations(u.T())
}

func (u *unitTestBookServiceSuite) TestFindOneById_Success() {
	data := &domain.Book{Id: 1, Title: "Atomic Habits: An Easy & Proven Way to Build Good Habits & Break Bad Ones", Author: "James Clear"}
	expected := &dto.BookResponse{Id: data.Id, Title: data.Title, Author: data.Author}

	u.brm.On("FindOneById", u.ctx, mock.Anything).Return(data, nil)

	result, err := u.bs.FindOneById(u.ctx, 1)

	u.NoError(err)
	u.NotNil(result)
	u.Equal(expected, result)

	u.brm.AssertExpectations(u.T())
}

func (u *unitTestBookServiceSuite) TestFindOneById_NotFound() {
	u.brm.On("FindOneById", u.ctx, mock.Anything).Return(nil, errors.New("no data found"))

	result, err := u.bs.FindOneById(u.ctx, 2)

	u.Error(err)
	u.Nil(result)

	u.brm.AssertExpectations(u.T())
}

func (u *unitTestBookServiceSuite) TestFindAll_Success() {
	data := []*domain.Book{
		{
			Id:     1,
			Title:  "Atomic Habits: An Easy & Proven Way to Build Good Habits & Break Bad Ones",
			Author: "James Clear",
		},
		{
			Id:     2,
			Title:  "The 7 Habits of Highly Effective People",
			Author: "Stephen R. Covey",
		},
	}

	var expected []*dto.BookResponse

	for _, e := range data {
		expected = append(expected, &dto.BookResponse{Id: e.Id, Title: e.Title, Author: e.Author})
	}

	u.brm.On("FindAll", u.ctx, mock.Anything).Return(data, nil)

	result, err := u.bs.FindAll(u.ctx)

	u.NotNil(result)
	u.NoError(err)
	u.Equal(expected, result)

	u.brm.AssertExpectations(u.T())
}

func (u *unitTestBookServiceSuite) TestFindAll_Failed() {
	u.brm.On("FindAll", u.ctx, mock.Anything).Return(nil, errors.New("something went wrong"))

	result, err := u.bs.FindAll(u.ctx)

	u.Nil(result)
	u.Error(err)

	u.brm.AssertExpectations(u.T())
}
