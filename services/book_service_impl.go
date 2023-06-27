package services

import (
	"context"

	"taiwodevops.com/crud_be/data/request"
	"taiwodevops.com/crud_be/data/response"
	"taiwodevops.com/crud_be/helper"
	"taiwodevops.com/crud_be/model"
	"taiwodevops.com/crud_be/repository"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
}

// Create implements BookService.
func (b *BookServiceImpl) Create(context context.Context, request request.CreateBookReq) {
	book := model.Book{
		Name: request.Name,
	}

	b.BookRepository.Save(context, book)
}

// Delete implements BookService.
func (b *BookServiceImpl) Delete(context context.Context, bookId int) {
	book, err := b.BookRepository.FindById(context, bookId)
	helper.PanicIfError(err)
	b.BookRepository.Delete(context, book.Id)
}

// FindAll implements BookService.
func (b *BookServiceImpl) FindAll(context context.Context) []response.BookResponse {
	books := b.BookRepository.FindAll(context)

	var booksResponse []response.BookResponse

	for _, value := range books {
		book := response.BookResponse{Id: value.Id, Name: value.Name}
		booksResponse = append(booksResponse, book)

	}
	return booksResponse

}

// FindById implements BookService.
func (b *BookServiceImpl) FindById(context context.Context, bookId int) response.BookResponse {
	book, err := b.BookRepository.FindById(context, bookId)
	helper.PanicIfError(err)
	return response.BookResponse(book)
}

// Update implements BookService.
func (b *BookServiceImpl) Update(context context.Context, request request.UpdateBookReq) {
	// first find if the book exist based on the request ID
	book, err := b.BookRepository.FindById(context, request.Id)
	helper.PanicIfError(err)

	// because app didn't quit, book exist. So take the (properties )name in that request and update/ assign to book name
	book.Name = request.Name
	b.BookRepository.Update(context, book)

}

func NewBookServiceImpl(bookRepository repository.BookRepository) BookService {
	return &BookServiceImpl{BookRepository: bookRepository}
}
