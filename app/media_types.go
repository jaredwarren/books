// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "Books": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/jaredwarren/redeam/design
// --out=$(GOPATH)/src/github.com/jaredwarren/redeam
// --version=v1.3.1

package app

import (
	"github.com/goadesign/goa"
	"time"
)

// A Book (default view)
//
// Identifier: application/vnd.book+json; view=default
type Book struct {
	// Author(s) of the book
	Author *string `form:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// Date of creation
	CreatedAt time.Time `form:"created_at" json:"created_at" xml:"created_at"`
	// Book ID
	ID int `form:"id" json:"id" xml:"id"`
	// Date of publication
	PublishDate *time.Time `form:"publish_date,omitempty" json:"publish_date,omitempty" xml:"publish_date,omitempty"`
	// Publisher of the book
	Publisher *string `form:"publisher,omitempty" json:"publisher,omitempty" xml:"publisher,omitempty"`
	Rating    int     `form:"rating" json:"rating" xml:"rating"`
	Status    *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Book title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Date of last change
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" xml:"updated_at"`
}

// Validate validates the Book media type instance.
func (mt *Book) Validate() (err error) {

	if mt.Rating < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, mt.Rating, 1, true))
	}
	if mt.Rating > 3 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, mt.Rating, 3, false))
	}
	if mt.Status != nil {
		if !(*mt.Status == "CheckedIn" || *mt.Status == "CheckedOut") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.status`, *mt.Status, []interface{}{"CheckedIn", "CheckedOut"}))
		}
	}
	return
}

// BookCollection is the media type for an array of Book (default view)
//
// Identifier: application/vnd.book+json; type=collection; view=default
type BookCollection []*Book

// Validate validates the BookCollection media type instance.
func (mt BookCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}
