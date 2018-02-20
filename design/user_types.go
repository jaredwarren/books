package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// BookPayload ...
var BookPayload = Type("BookPayload", func() {
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
})
