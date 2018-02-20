package main

import (
	"github.com/goadesign/goa"
	"github.com/jaredwarren/redeam/app"
	"github.com/jaredwarren/redeam/db"
)

// HealthController implements the health resource.
type HealthController struct {
	*goa.Controller
	DB *db.BookStore
}

// NewHealthController creates a health controller.
func NewHealthController(service *goa.Service, db *db.BookStore) *HealthController {
	return &HealthController{
		Controller: service.NewController("HealthController"),
		DB:         db,
	}
}

// Health runs the health action.
func (c *HealthController) Health(ctx *app.HealthHealthContext) error {
	return ctx.OK([]byte("ok"))
}
