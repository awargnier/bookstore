package handlers

import (
	"example/web-service-gin/models"
	"testing"
)

func TestGetBook(t *testing.T) {
	var books []models.Book

	book := models.Book{Title: "vingt-mille lieux sous les mer", Author: "jules verne", Price: 2}

	books = append(books, *GetBooksById(book.ID))

	if book.ID != books.ID {
		t.Error("Couldn't find correct book from bookstore")
	}

}
