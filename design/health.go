package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("health", func() {
	// Google AppEngine URL
	BasePath("/_ah")

	Action("health", func() {
		Routing(
			GET("/health"),
		)
		Description("Health check API")
		Response(OK, "text/plain")
	})
})
