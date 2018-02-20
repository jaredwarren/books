package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/boltdb/bolt"
	"github.com/goadesign/goa"
	"github.com/jaredwarren/redeam/app"
	"github.com/jaredwarren/redeam/app/test"
	"github.com/jaredwarren/redeam/db"
)

var sdb *db.BookStore

func setupTestCase(t *testing.T) func(t *testing.T) {
	// setup
	var err error
	bdb, err := bolt.Open("/data/testbooks.db", 0777, nil)
	if err != nil {
		panic(err)
	}

	err = bdb.Update(func(tx *bolt.Tx) error {
		// Start test with fresh bdb
		tx.DeleteBucket([]byte("Book"))
		if _, err = tx.CreateBucketIfNotExists([]byte("Book")); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		bdb.Close()
		panic(err)
	}

	sdb = db.NewBookStore(bdb)

	return func(t *testing.T) {
		defer bdb.Close()
		os.Remove("/data/testbooks.db")
		os.Remove("/data/testbooks.db.lock")
	}
}

// TestCreate test successfully createing a book
func TestCreate(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	var (
		service = goa.New("book")
		ctrl    = NewBookController(service, sdb)
		ctx     = context.Background()

		title       = randSeq(10)
		author      = randSeq(10)
		publisher   = randSeq(10)
		publishDate = time.Now()
		rating      = 3
		status      = "CheckedIn"
	)
	payload := &app.BookPayload{
		Title:       &title,
		Author:      &author,
		Publisher:   &publisher,
		PublishDate: &publishDate,
		Rating:      rating,
		Status:      &status,
	}

	r := test.CreateBookCreated(t, ctx, service, ctrl, payload)
	loc := r.Header().Get("Location")
	if loc == "" {
		t.Fatalf("missing Location header")
	}
	_, resID := filepath.Split(loc)
	bookID, err := strconv.Atoi(resID)
	if err != nil {
		t.Fatal(err)
	}

	book, err := getTestBook(bookID)
	if err != nil {
		t.Fatal(err)
	}

	if *book.Title != title {
		t.Fatal("Store title didn't match")
	}
}

// TestDelete test successfully deleting a book
func TestDelete(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	initialBook, err := insertTestBook(randSeq(10))
	if err != nil {
		t.Fatal(err)
	}

	var (
		service = goa.New("recipe")
		ctrl    = NewBookController(service, sdb)
		ctx     = context.Background()
	)

	test.DeleteBookNoContent(t, ctx, service, ctrl, initialBook.ID)

	_, err = getTestBook(initialBook.ID)
	if err == nil {
		t.Fatalf("book still in db")
	}

	expectedError := fmt.Errorf("key not found")
	if err.Error() != expectedError.Error() {
		t.Fatalf("Key Found")
	}
}

// TestList test successfully deleting a book
func TestList(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	bk1, err := insertTestBook(randSeq(10))
	if err != nil {
		t.Fatal(err)
	}

	bk2, err := insertTestBook(randSeq(10))
	if err != nil {
		t.Fatal(err)
	}

	var (
		service = goa.New("recipe")
		ctrl    = NewBookController(service, sdb)
		ctx     = context.Background()
	)

	_, books := test.ListBookOK(t, ctx, service, ctrl)

	if *books[0].Title != *bk1.Title {
		t.Fatalf("Book 1 title dosn't match")
	}

	if *books[1].Title != *bk2.Title {
		t.Fatalf("Book 2 title dosn't match")
	}
}

// TestShow test successfully deleting a book
func TestShow(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	_, err := insertTestBook(randSeq(10))
	if err != nil {
		t.Fatal(err)
	}

	bk2, err := insertTestBook(randSeq(10))
	if err != nil {
		t.Fatal(err)
	}

	var (
		service = goa.New("recipe")
		ctrl    = NewBookController(service, sdb)
		ctx     = context.Background()
	)

	_, book := test.ShowBookOK(t, ctx, service, ctrl, bk2.ID)

	if *book.Title != *bk2.Title {
		t.Fatalf("Book 2 title dosn't match")
	}
}

// TestUpdate test successfully updating a book
func TestUpdate(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	initialBook, err := insertTestBook(randSeq(10))
	if err != nil {
		t.Fatal(err)
	}

	var (
		service = goa.New("book")
		ctrl    = NewBookController(service, sdb)
		ctx     = context.Background()

		newTitle       = randSeq(10)
		newAuthor      = randSeq(10)
		newPublisher   = randSeq(10)
		newPublishDate = time.Now()
		newRating      = 3
		status         = "CheckedIn"
	)

	payload := &app.BookPayload{
		Title:       &newTitle,
		Author:      &newAuthor,
		Publisher:   &newPublisher,
		PublishDate: &newPublishDate,
		Rating:      newRating,
		Status:      &status,
	}
	test.UpdateBookNoContent(t, ctx, service, ctrl, initialBook.ID, payload)

	updatedBook, err := getTestBook(initialBook.ID)
	if err != nil {
		t.Fatal(err)
	}

	if updatedBook.Title == initialBook.Title {
		t.Fatalf("title didn't update, %s, %s", *updatedBook.Title, *initialBook.Title)
	}

	if *updatedBook.Title != newTitle {
		t.Fatalf("updated title didn't match input")
	}
}

// TestUpdateMissing test trying to updating a missing book
func TestUpdateMissing(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	title := "missing"
	initialBook := &app.Book{
		ID:    9999,
		Title: &title,
	}

	var (
		service = goa.New("book")
		ctrl    = NewBookController(service, sdb)
		ctx     = context.Background()

		newTitle = randSeq(10)
	)

	payload := &app.BookPayload{
		Title: &newTitle,
	}
	test.UpdateBookInternalServerError(t, ctx, service, ctrl, initialBook.ID, payload)

	_, err := getTestBook(initialBook.ID)
	expectedError := fmt.Errorf("key not found")
	if err.Error() != expectedError.Error() {
		t.Fatal(err)
	}
}

// utility to make inserting a book easier for testing
func insertTestBook(title string) (*app.Book, error) {
	book := &app.Book{
		Title:  &title,
		Rating: 1,
	}
	err := sdb.Insert(book)
	return book, err
}

// utility to make retrieving a book easier for testing
func getTestBook(bookID int) (*app.Book, error) {
	return sdb.Fetch(bookID)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Generate a random string for testing
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
