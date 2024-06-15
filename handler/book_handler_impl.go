package handler

import (
	"gin-go-testing/model/dto"
	"gin-go-testing/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/go-custom-err/errs"
)

type bookHandlerImpl struct {
	bs service.BookService
}

func NewBookHandlerImpl(bs service.BookService) BookHandler {
	return &bookHandlerImpl{bs}
}

func (b *bookHandlerImpl) Create(ctx *gin.Context) {
	bookDto := new(dto.NewBookRequest)

	if err := ctx.ShouldBindJSON(bookDto); err != nil {
		unprocessableEntityError := errs.NewUnprocessableEntityError("invalid json request body")
		ctx.AbortWithStatusJSON(unprocessableEntityError.StatusCode(), unprocessableEntityError)
		return
	}

	result, err := b.bs.Create(ctx, bookDto)
	if err != nil {
		ctx.AbortWithStatusJSON(err.StatusCode(), err)
		return
	}

	response := &dto.APIResponse{
		Status:     http.StatusText(http.StatusCreated),
		StatusCode: http.StatusCreated,
		Message:    "success",
		Data:       result,
	}

	ctx.JSON(http.StatusCreated, response)
}

func (b *bookHandlerImpl) FindOneById(ctx *gin.Context) {
	bookId, errConv := strconv.Atoi(ctx.Param("bookId"))
	if errConv != nil {
		unprocessableEntityErr := errs.NewUnprocessableEntityError("bookId param must be a valid number")
		ctx.AbortWithStatusJSON(unprocessableEntityErr.StatusCode(), unprocessableEntityErr)
		return
	}

	result, err := b.bs.FindOneById(ctx, uint(bookId))
	if err != nil {
		ctx.AbortWithStatusJSON(err.StatusCode(), err)
		return
	}

	response := &dto.APIResponse{
		Status:     http.StatusText(http.StatusOK),
		StatusCode: http.StatusOK,
		Message:    "success",
		Data:       result,
	}

	ctx.JSON(http.StatusOK, response)

}

func (b *bookHandlerImpl) FindAll(ctx *gin.Context) {
	result, err := b.bs.FindAll(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(err.StatusCode(), err)
		return
	}

	response := &dto.APIResponse{
		Status:     http.StatusText(http.StatusOK),
		StatusCode: http.StatusOK,
		Message:    "success",
		Data:       result,
	}

	ctx.JSON(http.StatusOK, response)
}
