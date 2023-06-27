package repository

import (
	"context"
	"database/sql"
	"errors"

	"taiwodevops.com/crud_be/helper"
	"taiwodevops.com/crud_be/model"
)

type BookRepositoryImpl struct {
	Db *sql.DB
}

// Delete implements BookRepository.
func (b *BookRepositoryImpl) Delete(ctx context.Context, bookId int) {
	// instantiate DB instance
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollBack(tx)

	// write the sql for delete
	SQL := "Delete from book where id =$1"
	// use this queries to delete
	_, errCtx := tx.ExecContext(ctx, SQL, bookId)
	helper.PanicIfError(errCtx)
}

// FindAll implements BookRepository.
func (b *BookRepositoryImpl) FindAll(ctx context.Context) []model.Book {
	// instantiate DB instance
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollBack(tx)

	// statement for init find all
	SQL := "select id,name from"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var books []model.Book

	for result.Next() {
		book := model.Book{}
		err := result.Scan(&book.Id, &book.Name)
		helper.PanicIfError(err)

		books = append(books, book)
	}

	return books
}

// FindById implements BookRepository.
func (b *BookRepositoryImpl) FindById(ctx context.Context, bookId int) (model.Book, error) {
	// instantiate DB instance
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollBack(tx)

	// statement for init find all
	SQL := "select id,name from books where ID is =$1"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)
	defer result.Close()

	book := model.Book{}

	if result.Next() {
		err := result.Scan(&book.Id, &book.Name)
		helper.PanicIfError(err) // if error is found, this terminates the search
		return book, nil
	} else {
		return book, errors.New("book id not found")
	}

}

// Save implements BookRepository.
func (b *BookRepositoryImpl) Save(ctx context.Context, book model.Book) {
	// instantiate DB instance
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	// statement saving a new record
	SQL := "insert into book(name) and values ($1)"
	_, errQuery := tx.ExecContext(ctx, SQL)
	helper.PanicIfError(errQuery)

}

// Update implements BookRepository.
func (b *BookRepositoryImpl) Update(ctx context.Context, book model.Book) {
	// instantiate DB instance
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	// statement saving a new record
	SQL := "update book record name by=$1 where id=$2"
	_, errQuery := tx.ExecContext(ctx, SQL, book.Name, book.Id)
	helper.PanicIfError(errQuery)
}

// create a func that returns a book repositpry

func NewBookRepositoryImpl(DB *sql.DB) BookRepository {
	return &BookRepositoryImpl{Db: DB}
}
