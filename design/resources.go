package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Printer API
var _ = Resource("book", func() {
	DefaultMedia(Book)
	BasePath("/books")

	Action("list", func() {
		Routing(GET(""))
		Description("Retrieve all books.")
		Response(InternalServerError, ErrorMedia)
		Response(OK, CollectionOf(Book))
	})

	Action("show", func() {
		Routing(GET("/:bookID"))
		Description("Show a book.")
		Params(func() {
			Param("bookID", Integer, "Book ID", func() {
				Minimum(1)
			})
		})
		Response(OK)
		Response(NotFound)
		Response(InternalServerError, ErrorMedia)
	})

	Action("create", func() {
		Routing(POST(""))
		Description("Create a new book")
		Payload(BookPayload)
		Response(Created, "/books/[0-9]+")
		Response(InternalServerError, ErrorMedia)
	})

	Action("update", func() {
		Routing(PUT("/:bookID"))
		Description("Change book data")
		Params(func() {
			Param("bookID", Integer, "Book ID")
		})
		Payload(BookPayload)
		Response(NoContent)
		Response(NotFound)
		Response(InternalServerError, ErrorMedia)
	})

	Action("delete", func() {
		Routing(DELETE("/:bookID"))
		Params(func() {
			Param("bookID", Integer, "Book ID", func() {
				Minimum(1)
			})
			Required("bookID")
		})
		Response(NoContent)
		Response(NotFound)
		Response(InternalServerError, ErrorMedia)
	})
})
