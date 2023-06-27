package controllers

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"taiwodevops.com/crud_be/data/request"
	"taiwodevops.com/crud_be/data/response"
	"taiwodevops.com/crud_be/helper"
	"taiwodevops.com/crud_be/services"
)

type BookController struct {
	BookService services.BookService
}

// helper function for creating an object of Book service
func NewBookController(bS services.BookService) *BookController {
	return &BookController{BookService: bS}
}

// create func literals (anonymous func)
func (c *BookController) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookCreateRequest := request.CreateBookReq{}
	helper.ReadRequestBody(r, &bookCreateRequest)

	c.BookService.Create(r.Context(), bookCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}

	helper.WriteResponseBody(w, webResponse)

}

func (c *BookController) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookUpdateRequest := request.UpdateBookReq{}
	helper.ReadRequestBody(r, &bookUpdateRequest)

	// get the ID of the book  from params and parse it as int
	bookId := p.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)
	bookUpdateRequest.Id = id

	c.BookService.Update(r.Context(), bookUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}

	helper.WriteResponseBody(w, webResponse)

}

func (c *BookController) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// get the ID of the book  from params and parse it as int
	bookId := p.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	c.BookService.Delete(r.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}

	helper.WriteResponseBody(w, webResponse)

}

func (c *BookController) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	results := c.BookService.FindAll(r.Context())

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   results,
	}

	helper.WriteResponseBody(w, webResponse)

}

func (c *BookController) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// get the ID of the book from params and parse it as int
	bookId := p.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	c.BookService.FindById(r.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}

	helper.WriteResponseBody(w, webResponse)

}
