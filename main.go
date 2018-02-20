//go:generate goagen bootstrap -d github.com/jaredwarren/redeam/design

package main

import (
	"github.com/boltdb/bolt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jaredwarren/redeam/app"
	"github.com/jaredwarren/redeam/db"
)

func main() {
	// Create service
	service := goa.New("Books")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Open the books.db data file in current directory.
	bdb, err := bolt.Open("/data/books.db", 0777, nil)
	if err != nil {
		panic(err)
	}

	err = bdb.Update(func(tx *bolt.Tx) error {
		if _, err = tx.CreateBucketIfNotExists([]byte("Book")); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		bdb.Close()
		panic(err)
	}

	bookDb := db.NewBookStore(bdb)

	// Mount "book" controller
	c := NewBookController(service, bookDb)
	app.MountBookController(service, c)
	// Mount "health" controller
	c2 := NewHealthController(service, bookDb)
	app.MountHealthController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
