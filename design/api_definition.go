package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("Books", func() {
	Title("Redeam.com Books")
	Description("Redeam.com Book Test")
	Contact(func() {
		Name("Jared Warren")
		Email("jlwarren1@gmail.com")
		URL("jlwarren1.com")
	})
	Docs(func() {
		Description("Books Service")
		URL("http://jlwarren1.com")
	})
	Host("localhost:8443")
	Scheme("https")
})
