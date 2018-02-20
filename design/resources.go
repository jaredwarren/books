package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Printer API
var _ = Resource("book", func() {
	DefaultMedia(BookPayload)
	BasePath("/books")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Retrieve all books.")
		Response(OK, CollectionOf(BookPayload))
	})

	Action("show", func() {
		Routing(
			GET("/:bookID"),
		)
		Description("Show a book.")
		Params(func() {
			Param("bookID", Integer, "Book ID", func() {
				Minimum(1)
			})
		})
		Response(OK)
		Response(NotFound)
		// Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create a new book")
		Payload(func() {
			Member("title")
			Required("title")
		})
		Response(Created, "/books/[0-9]+")
		// Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PUT("/:bookID"),
		)
		Description("Change book data")
		Params(func() {
			Param("bookID", Integer, "Book ID")
		})
		Payload(func() {
			Member("title")
			Required("title")
		})
		Response(NoContent)
		Response(NotFound)
		// Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:bookID"),
		)
		Params(func() {
			Param("bookD", Integer, "Book ID")
		})
		Response(NoContent)
		Response(NotFound)
		// Response(BadRequest, ErrorMedia)
	})
})
