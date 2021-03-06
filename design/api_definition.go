package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("Books", func() {
	Title("Books")
	Description("Book CRUD API")
	Contact(func() {
		Name("Jared Warren")
		Email("jlwarren1@gmail.com")
		URL("http://jlwarren1.com")
	})
	Docs(func() {
		Description("Books Service")
		URL("http://jlwarren1.com")
	})
	Host("localhost:8080")
	Scheme("http")

	ResponseTemplate(Created, func(pattern string) {
		Status(201)
		Headers(func() {
			Header("Location", String, "href to created resource", func() {
				Pattern(pattern)
			})
		})
	})
})
