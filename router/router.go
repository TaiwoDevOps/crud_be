package router

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"taiwodevops.com/crud_be/controllers"
	"taiwodevops.com/crud_be/helper"
)

func getHome(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\nThis is your landing page.", ps.ByName("name"))
}

func NewRouter(c *controllers.BookController) *httprouter.Router {

	routes := httprouter.New()

	routes.GET("/:name", getHome)
	server := http.ListenAndServe("localhost:8888", routes)

	helper.PanicIfError(server)

	routes.GET("/api/book", c.FindAll)
	routes.GET("/api/book/:bookId", c.FindById)
	routes.PATCH("/api/book/:bookId", c.Update)
	routes.DELETE("/api/book/:bookId", c.Delete)
	routes.POST("/api/book", c.Create)

	return routes

}
