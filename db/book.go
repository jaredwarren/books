package db

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"time"

	"github.com/boltdb/bolt"
	"github.com/jaredwarren/redeam/app"
)

// BookStore wrapper for db
type BookStore struct {
	db *bolt.DB
}

// NewBookStore create a new book store
func NewBookStore(db *bolt.DB) *BookStore {
	return &BookStore{
		db: db,
	}
}

// Fetch book by id
func (s *BookStore) Fetch(bookID int) (*app.Book, error) {
	book := &app.Book{}
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Book"))

		v := b.Get(Itob(bookID))
		if v == nil {
			return errors.New("key not found")
		}
		err := json.Unmarshal(v, book)
		if err != nil {
			return err
		}

		return nil
	})
	return book, err
}

// FetchAll return list of books
func (s *BookStore) FetchAll() (app.BookCollection, error) {
	books := app.BookCollection{}

	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Book"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			book := &app.Book{}
			err := json.Unmarshal(v, book)
			if err != nil {
				return err
			}

			books = append(books, book)
		}

		return nil
	})

	return books, err
}

// Insert book, returns inserted book id
func (s *BookStore) Insert(book *app.Book) error {
	err := s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Book"))
		id, _ := b.NextSequence()

		book.ID = int(id)
		book.CreatedAt = time.Now()
		book.UpdatedAt = book.CreatedAt

		buf, err := json.Marshal(book)
		if err != nil {
			return err
		}

		// Persist bytes to users bucket.
		return b.Put(Itob(book.ID), buf)
	})

	return err
}

// Update book
func (s *BookStore) Update(book *app.Book) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Book"))

		book.UpdatedAt = book.CreatedAt

		buf, err := json.Marshal(book)
		if err != nil {
			return err
		}

		// Persist bytes to users bucket.
		return b.Put(Itob(book.ID), buf)
	})
}

// Delete book
func (s *BookStore) Delete(bookID int) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Book"))
		return b.Delete(Itob(bookID))
	})
}

// Itob returns an 8-byte big endian representation of v.
func Itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
