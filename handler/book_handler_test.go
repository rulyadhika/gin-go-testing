package handler

import (
	"bytes"
	"encoding/json"
	"gin-go-testing/mocks"
	"gin-go-testing/model/dto"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/go-custom-err/errs"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type unitTestBookHandlerSuite struct {
	suite.Suite
	bh     BookHandler
	bsm    *mocks.BookService
	ctx    *gin.Context
	writer *httptest.ResponseRecorder
}

func TestUnitTestBookHandler(t *testing.T) {
	suite.Run(t, &unitTestBookHandlerSuite{})
}

func (u *unitTestBookHandlerSuite) SetupTest() {
	bookServMock := mocks.NewBookService(u.T())
	u.bsm = bookServMock

	u.bh = NewBookHandlerImpl(bookServMock)

	// set context and response recorder
	gin.SetMode(gin.TestMode)

	writer := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(writer)
	u.ctx = ctx
	u.writer = writer
}

func (u *unitTestBookHandlerSuite) TestFindOneById_Success() {
	// setup expected result
	bookId := uint(1)

	data := &dto.BookResponse{
		Id:     bookId,
		Title:  "Atomic Habits: An Easy & Proven Way to Build Good Habits & Break Bad Ones",
		Author: "James Clear",
	}

	expectedDataMap := map[string]any{
		"id":     float64(data.Id),
		"title":  data.Title,
		"author": data.Author,
	}

	expected := dto.APIResponse{
		Status:     http.StatusText(http.StatusOK),
		StatusCode: http.StatusOK,
		Message:    "success",
		Data:       expectedDataMap,
	}

	// mock service method
	u.bsm.On("FindOneById", u.ctx, bookId).Return(data, nil)

	// set route params
	u.ctx.Params = gin.Params{{Key: "bookId", Value: strconv.Itoa(int(bookId))}}

	// call the method
	u.bh.FindOneById(u.ctx)

	var apiResponse dto.APIResponse
	err := json.Unmarshal(u.writer.Body.Bytes(), &apiResponse)

	u.NoError(err)
	u.Equal(expected, apiResponse)

	u.bsm.AssertExpectations(u.T())
}

func (u *unitTestBookHandlerSuite) TestFindOneById_NotFound() {
	// setup expected result
	bookId := uint(1)

	expected := dto.APIResponse{
		Status:     http.StatusText(http.StatusNotFound),
		StatusCode: http.StatusNotFound,
		Message:    "data not found",
		Data:       nil,
	}

	// mock service method
	u.bsm.On("FindOneById", u.ctx, bookId).Return(nil, errs.NewNotFoundError("data not found"))

	// set route params
	u.ctx.Params = gin.Params{{Key: "bookId", Value: strconv.Itoa(int(bookId))}}

	// call the method
	u.bh.FindOneById(u.ctx)

	var apiResponse dto.APIResponse
	err := json.Unmarshal(u.writer.Body.Bytes(), &apiResponse)

	u.NoError(err)
	u.Equal(expected, apiResponse)

	u.bsm.AssertExpectations(u.T())
}

func (u *unitTestBookHandlerSuite) TestCreate_Success() {
	data := &dto.BookResponse{
		Id:     1,
		Title:  "Atomic Habits: An Easy & Proven Way to Build Good Habits & Break Bad Ones",
		Author: "James Clear",
	}

	expectedDataMap := map[string]any{
		"id":     float64(data.Id),
		"title":  data.Title,
		"author": data.Author,
	}

	expected := dto.APIResponse{
		Status:     http.StatusText(http.StatusCreated),
		StatusCode: http.StatusCreated,
		Message:    "success",
		Data:       expectedDataMap,
	}

	u.bsm.On("Create", u.ctx, mock.Anything).Return(data, nil)

	// create request body
	requestData := dto.NewBookRequest{
		Title:  data.Title,
		Author: data.Author,
	}

	requestBody, _ := json.Marshal(requestData)
	u.ctx.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))

	u.bh.Create(u.ctx)

	var apiResponse dto.APIResponse
	err := json.Unmarshal(u.writer.Body.Bytes(), &apiResponse)

	u.NoError(err)
	u.Equal(expected, apiResponse)

	u.bsm.AssertExpectations(u.T())
}

func (u *unitTestBookHandlerSuite) TestCreate_Failed() {
	data := &dto.BookResponse{
		Id:     1,
		Title:  "Atomic Habits: An Easy & Proven Way to Build Good Habits & Break Bad Ones",
		Author: "James Clear",
	}

	expected := dto.APIResponse{
		Status:     http.StatusText(http.StatusInternalServerError),
		StatusCode: http.StatusInternalServerError,
		Message:    "something went wrong",
		Data:       nil,
	}

	u.bsm.On("Create", u.ctx, mock.Anything).Return(nil, errs.NewInternalServerError("something went wrong"))

	// create request body
	requestData := dto.NewBookRequest{
		Title:  data.Title,
		Author: data.Author,
	}

	requestBody, _ := json.Marshal(requestData)
	u.ctx.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))

	u.bh.Create(u.ctx)

	var apiResponse dto.APIResponse
	err := json.Unmarshal(u.writer.Body.Bytes(), &apiResponse)

	u.NoError(err)
	u.Equal(expected, apiResponse)

	u.bsm.AssertExpectations(u.T())
}

func (u *unitTestBookHandlerSuite) TestFindAll_Success() {
	data := []*dto.BookResponse{
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

	u.bsm.On("FindAll", u.ctx).Return(data, nil)

	var expectedDataMap []any

	for _, e := range data {
		expectedDataMap = append(expectedDataMap, map[string]any{
			"id":     float64(e.Id),
			"title":  e.Title,
			"author": e.Author,
		})
	}

	expected := dto.APIResponse{
		Status:     http.StatusText(http.StatusOK),
		StatusCode: http.StatusOK,
		Message:    "success",
		Data:       expectedDataMap,
	}

	u.bh.FindAll(u.ctx)

	var apiResponse dto.APIResponse
	err := json.Unmarshal(u.writer.Body.Bytes(), &apiResponse)
	u.NoError(err)
	u.Equal(expected, apiResponse)

	u.bsm.AssertExpectations(u.T())
}

func (u *unitTestBookHandlerSuite) TestFindAll_Failed() {
	u.bsm.On("FindAll", u.ctx).Return(nil, errs.NewInternalServerError("something went wrong"))

	expected := dto.APIResponse{
		Status:     http.StatusText(http.StatusInternalServerError),
		StatusCode: http.StatusInternalServerError,
		Message:    "something went wrong",
		Data:       nil,
	}

	u.bh.FindAll(u.ctx)

	var apiResponse dto.APIResponse
	err := json.Unmarshal(u.writer.Body.Bytes(), &apiResponse)
	u.NoError(err)
	u.Equal(expected, apiResponse)

	u.bsm.AssertExpectations(u.T())
}
