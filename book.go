package main

import (
	"github.com/goadesign/goa"
	"github.com/jaredwarren/redeam/app"
	"github.com/jaredwarren/redeam/db"
)

// BookController implements the book resource.
type BookController struct {
	*goa.Controller
	DB *db.BookStore
}

// NewBookController creates a book controller.
func NewBookController(service *goa.Service, db *db.BookStore) *BookController {
	return &BookController{
		Controller: service.NewController("BookController"),
		DB:         db,
	}
}

// Create runs the create action.
func (c *BookController) Create(ctx *app.CreateBookContext) error {
	book := &app.Book{}
	if ctx.Payload.Title != nil {
		book.Title = ctx.Payload.Title
	}
	if ctx.Payload.Author != nil {
		book.Author = ctx.Payload.Author
	}
	if ctx.Payload.Publisher != nil {
		book.Publisher = ctx.Payload.Publisher
	}
	if ctx.Payload.PublishDate != nil {
		book.PublishDate = ctx.Payload.PublishDate
	}
	book.Rating = ctx.Payload.Rating
	if ctx.Payload.Status != nil {
		book.Status = ctx.Payload.Status
	}

	err := c.DB.Insert(book)
	if err != nil {
		return ctx.InternalServerError(err)
	}
	ctx.ResponseData.Header().Set("Location", app.BookHref(book.ID))
	return ctx.Created()
}

// Delete runs the delete action.
func (c *BookController) Delete(ctx *app.DeleteBookContext) error {
	err := c.DB.Delete(ctx.BookID)
	if err != nil {
		return ctx.InternalServerError(err)
	}

	return ctx.NoContent()
}

// List runs the list action.
func (c *BookController) List(ctx *app.ListBookContext) error {
	books, err := c.DB.FetchAll()
	if err != nil {
		return ctx.InternalServerError(err)
	}

	return ctx.OK(books)
}

// Show runs the show action.
func (c *BookController) Show(ctx *app.ShowBookContext) error {
	book, err := c.DB.Fetch(ctx.BookID)
	if err != nil {
		return ctx.InternalServerError(err)
	}

	return ctx.OK(book)
}

// Update runs the update action.
func (c *BookController) Update(ctx *app.UpdateBookContext) error {
	book, err := c.DB.Fetch(ctx.BookID)
	if err != nil {
		return ctx.InternalServerError(err)
	}

	if ctx.Payload.Title != nil {
		book.Title = ctx.Payload.Title
	}
	if ctx.Payload.Author != nil {
		book.Author = ctx.Payload.Author
	}
	if ctx.Payload.Publisher != nil {
		book.Publisher = ctx.Payload.Publisher
	}
	if ctx.Payload.PublishDate != nil {
		book.PublishDate = ctx.Payload.PublishDate
	}
	book.Rating = ctx.Payload.Rating
	if ctx.Payload.Status != nil {
		book.Status = ctx.Payload.Status
	}

	err = c.DB.Update(book)
	if err != nil {
		return ctx.InternalServerError(err)
	}

	return ctx.NoContent()
}
