package services

import (
	"context"

	"taiwodevops.com/crud_be/data/request"
	"taiwodevops.com/crud_be/data/response"
)

type BookService interface {
	Create(context context.Context, request request.CreateBookReq)
	Update(context context.Context, request request.UpdateBookReq)
	Delete(context context.Context, bookId int)
	FindById(context context.Context, bookId int) response.BookResponse
	FindAll(context context.Context) []response.BookResponse
}
