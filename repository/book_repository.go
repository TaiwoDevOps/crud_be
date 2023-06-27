package repository

import (
	"context"

	"taiwodevops.com/crud_be/model"
)

// use an interface to design the method for the repository

type BookRepository interface {
	Save(ctx context.Context, book model.Book)
	Update(ctx context.Context, book model.Book) // how do you update a book without an ID though
	Delete(ctx context.Context, bookId int)
	FindById(ctx context.Context, bookId int) (model.Book, error) // returns error is book not found by Id or returns book
	FindAll(ctx context.Context) []model.Book                     // returns an array of books
}
