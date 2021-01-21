package books

import (
	"github.com/bhanupratap1810/bookstore_users-api/controllers"
	"github.com/bhanupratap1810/bookstore_users-api/domain/books"
	"github.com/bhanupratap1810/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getBookId(bookIdParam string) (int64, *errors.RestErr) {

	bookId, bookErr := strconv.ParseInt(bookIdParam, 10, 64)
	if bookErr != nil {
		return 0, errors.NewBadRequestError("book id should be a number")
	}
	return bookId, nil
}

//func getUserId(userIdParam string) (int64, *errors.RestErr) {
//	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
//	if userErr != nil {
//		return 0, errors.NewBadRequestError("user id should be a number")
//	}
//	return userId, nil
//}

func GetBooksHandler(service controllers.Service) gin.HandlerFunc {
	return func(c * gin.Context) {
		Books, err := service.BookServiceImpl.GetBooks()
		if err != nil {
			c.JSON(err.Status, err)
			return
		}
		c.JSON(http.StatusOK, Books)
	}
}

func CreateBookHandler(service controllers.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		//service.UserServiceImpl.GetUser()
		var book books.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			restErr := errors.NewBadRequestError("invalid json body")
			c.JSON(restErr.Status, restErr)
			return
		}
		b, err := service.BookServiceImpl.CreateBook(book)
		if err != nil {
			c.JSON(err.Status, err)
			return
		}
		c.JSON(http.StatusCreated, b)
	}
}

func UpdateBookHandler(service controllers.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookId, idErr := getBookId(c.Param("book_id"))
		if idErr != nil {
			c.JSON(idErr.Status, idErr)
			return
		}
		var book books.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			restErr := errors.NewBadRequestError("invalid json body")
			c.JSON(restErr.Status, restErr)
			return
		}

		book.BookId = bookId

		isPartial := c.Request.Method == http.MethodPatch

		result, err := service.BookServiceImpl.UpdateBook(isPartial, book)
		if err != nil {
			c.JSON(err.Status, err)
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func DeleteBookHandler(service controllers.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookId, idErr := getBookId(c.Param("book_id"))
		if idErr != nil {
			c.JSON(idErr.Status, idErr)
			return
		}
		if err := service.BookServiceImpl.DeleteBook(bookId); err != nil {
			c.JSON(err.Status, err)
			return
		}
		c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
	}
}

func SearchBookByIdHandler(service controllers.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		bookId, idErr := getBookId(c.Param("book_id"))
		if idErr != nil {
			c.JSON(idErr.Status, idErr)
			return
		}

		book, getErr := service.BookServiceImpl.SearchBookById(bookId)
		if getErr != nil {
			c.JSON(getErr.Status, getErr)
			return
		}
		c.JSON(http.StatusOK, book)
	}
}

//func SearchBookByUserHandler(service controllers.Service) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		userId, idErr := getUserId(c.Param("user_id"))
//		if idErr != nil {
//			c.JSON(idErr.Status, idErr)
//			return
//		}
//
//		book, getErr := service.BookServiceImpl.SearchBookByUser(userId)
//		if getErr != nil {
//			c.JSON(getErr.Status, getErr)
//			return
//		}
//		c.JSON(http.StatusOK, book)
//	}
//}