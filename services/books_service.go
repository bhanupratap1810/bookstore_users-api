package services

import (
	"github.com/bhanupratap1810/bookstore_users-api/domain/books"
	"github.com/bhanupratap1810/bookstore_users-api/utils/errors"
)

type BookService interface {
	GetBooks() (books.Books, *errors.RestErr)
	CreateBook(books.Book) (*books.Book, *errors.RestErr)
	UpdateBook(bool, books.Book) (*books.Book, *errors.RestErr)
	DeleteBook(int64) *errors.RestErr
	SearchBookById(int64) (*books.Book, *errors.RestErr)
	SearchBookByUser(int64) (*books.Book, *errors.RestErr)
}

type BookServiceImpl struct {
	bookDaoService books.BookDaoService
}

func NewBookServiceImpl(bookDaoService books.BookDaoService) BookService {
	return &BookServiceImpl{bookDaoService: bookDaoService}
}

func (b *BookServiceImpl) GetBooks() (books.Books, *errors.RestErr){
	return b.bookDaoService.Get(&books.Book{})
}

func (b *BookServiceImpl) CreateBook(book books.Book) (*books.Book, *errors.RestErr){
	if err := book.Validate(); err != nil {
		return nil, err
	}
	if err := b.bookDaoService.Save(&book); err != nil {
		return nil, err
	}
	return &book, nil
}

func (b *BookServiceImpl) UpdateBook(isPartial bool, book books.Book) (*books.Book, *errors.RestErr) {
	current, err := b.SearchBookById(book.BookId)
	if err != nil {
		return nil, err
	}
	if isPartial {
		if book.BookName != "" {
			current.BookName = book.BookName
		}
		if book.BookAuthor != "" {
			current.BookAuthor = book.BookAuthor
		}
		if book.BookType != "" {
			current.BookType = book.BookType
		}
		if book.BorrowerId != 0 {
			current.BorrowerId = book.BorrowerId
		}
		if book.Status != "" {
			current.Status = book.Status
		}
	} else {
		current.BookName = book.BookName
		current.BookAuthor = book.BookAuthor
		current.BookType = book.BookType
		current.BorrowerId = book.BorrowerId
		current.Status = book.Status

	}

	if err := b.bookDaoService.Update(current); err != nil {
		return nil, err
	}
	return current, nil
}
func (b *BookServiceImpl) DeleteBook(bookId int64) *errors.RestErr {
	book := &books.Book{BookId: bookId}
	return b.bookDaoService.Delete(book)
}
func (b *BookServiceImpl) SearchBookById(bookId int64) (*books.Book, *errors.RestErr) {
	result := &books.Book{BookId: bookId}
	err := b.bookDaoService.FindByBookId(result)
	if err != nil {
		//log
		return nil, err
	}
	return result, nil
}
func (b *BookServiceImpl) SearchBookByUser(userId int64) (*books.Book, *errors.RestErr){
	result := &books.Book{BorrowerId: userId}
	err := b.bookDaoService.FindByUserId(result)
	if err != nil {
		//log
		return nil, err
	}
	return result, nil
}