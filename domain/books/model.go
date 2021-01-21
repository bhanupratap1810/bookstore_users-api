package books

import (
	"github.com/bhanupratap1810/bookstore_users-api/utils/errors"
	"strings"
)

type Book struct {
	BookId          int64  `json:"book_id"`
	BookName   string `json:"book_name"`
	BookAuthor  string `json:"book_author"`
	BookType string `json:"book_type"`
	Status string `json:"status"`
}

type Books []Book

func (book *Book) Validate() *errors.RestErr {
	book.BookName = strings.TrimSpace(book.BookName)
	book.BookAuthor = strings.TrimSpace(book.BookAuthor)
	book.BookType = strings.TrimSpace(book.BookType)
	book.Status = strings.TrimSpace(strings.ToLower(book.Status))
	if book.Status == ""{
		book.Status = "not issued"
	}
	return nil
}