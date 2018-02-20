package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Book ...
var Book = MediaType("application/vnd.book+json", func() {
	Description("A Book")
	Attributes(func() {
		Attribute("title", String, "Book title")
		Attribute("author", String, "Author(s) of the book")
		Attribute("publisher", String, "Publisher of the book")
		Attribute("publish_date", DateTime, "Date of publication")
		Attribute("rating", Integer, func() {
			Minimum(1)
			Maximum(3)
			Example(1)
			Default(1)
		})
		Attribute("status", func() {
			Enum("CheckedIn", "CheckedOut")
		})

		Attribute("created_at", DateTime, "Date of creation")
		Attribute("updated_at", DateTime, "Date of last change")

		Required("title", "author", "publisher", "publish_date", "rating", "status", "created_at", "created_by")
	})

	View("default", func() {
		Attribute("title")
		Attribute("author")
		Attribute("publisher")
		Attribute("publish_date")
		Attribute("rating")
		Attribute("status")
		Attribute("created_at")
		Attribute("created_by")
	})
})
