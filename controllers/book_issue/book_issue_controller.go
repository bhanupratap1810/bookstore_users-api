package book_issue

import (
	"github.com/bhanupratap1810/bookstore_users-api/controllers"
	"github.com/bhanupratap1810/bookstore_users-api/controllers/middlewares"
	"github.com/bhanupratap1810/bookstore_users-api/domain/book_issue"
	"github.com/bhanupratap1810/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateBookIssueHandler(service controllers.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		//service.UserServiceImpl.GetUser()

		userid, role, err1 := middlewares.GetUserIdAndRoleFromContext(c)
		if err1 != nil {
			restErr:=errors.NewUnauthorizedError("unauthenticated")
			c.JSON(restErr.Status, restErr)
			return
		}

		var bookIssue book_issue.BookIssue
		if err := c.ShouldBindJSON(&bookIssue); err != nil {
			restErr := errors.NewBadRequestError("invalid json body")
			c.JSON(restErr.Status, restErr)
			return
		}

		if userid==0 || role== "unemployed"{
			restErr:=errors.NewBadRequestError("not permitted")
			c.JSON(restErr.Status, restErr)
			return
		}

		if role!="admin" && userid!=bookIssue.UserId{
			restErr:=errors.NewForbiddenError("only admin or the same user can issue a book")
			c.JSON(restErr.Status, restErr)
			return
		}
		b, err := service.BookIssueServiceImpl.CreateBookIssue(bookIssue)
		if err != nil {
			c.JSON(err.Status, err)
			return
		}
		c.JSON(http.StatusCreated, b)
	}
}

func GetBookIssueHandler(service controllers.Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		_, role, err1 := middlewares.GetUserIdAndRoleFromContext(c)
		if err1 != nil {
			restErr:=errors.NewUnauthorizedError("unauthenticated")
			c.JSON(restErr.Status, restErr)
			return
		}

		if role!="admin" {
			restErr:=errors.NewForbiddenError("only admin can view all the book issue details")
			c.JSON(restErr.Status, restErr)
			return
		}
		BookIssue, err := service.BookIssueServiceImpl.GetBookIssue()
		if err != nil {
			c.JSON(err.Status, err)
			return
		}
		c.JSON(http.StatusOK, BookIssue)
	}
}

func GetBookIssueByBookHandler(service controllers.Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		_, role, err1 := middlewares.GetUserIdAndRoleFromContext(c)
		if err1 != nil {
			restErr:=errors.NewUnauthorizedError("unauthenticated")
			c.JSON(restErr.Status, restErr)
			return
		}

		if role!="admin" {
			restErr:=errors.NewForbiddenError("only admin can check the book issue details")
			c.JSON(restErr.Status, restErr)
			return
		}
		bookId, idErr := middlewares.GetId(c.Param("id"))
		if idErr != nil {
			c.JSON(idErr.Status, idErr)
			return
		}

		bookIssue, getErr := service.BookIssueServiceImpl.GetBookIssueByBook(bookId)
		if getErr != nil {
			c.JSON(getErr.Status, getErr)
			return
		}
		c.JSON(http.StatusOK, bookIssue)
	}
}

func SearchBookIssueByUserHandler(service controllers.Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		userid, role, err1 := middlewares.GetUserIdAndRoleFromContext(c)
		if err1 != nil {
			restErr:=errors.NewUnauthorizedError("unauthenticated")
			c.JSON(restErr.Status, restErr)
			return
		}

		userId, idErr := middlewares.GetId(c.Param("id"))
		if idErr != nil {
			c.JSON(idErr.Status, idErr)
			return
		}

		if userid==0 || role== "unemployed"{
			restErr:=errors.NewBadRequestError("not permitted")
			c.JSON(restErr.Status, restErr)
			return
		}

		if role!="admin" && userid!=userId{
			restErr:=errors.NewForbiddenError("only admin or the same user can check their issue details")
			c.JSON(restErr.Status, restErr)
			return
		}

		bookIssue, getErr := service.BookIssueServiceImpl.GetBookIssueByUser(userId)
		if getErr != nil {
			c.JSON(getErr.Status, getErr)
			return
		}
		c.JSON(http.StatusOK, bookIssue)
	}
}

func DeleteBookHandler(service controllers.Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		_, role, err1 := middlewares.GetUserIdAndRoleFromContext(c)
		if err1 != nil {
			restErr:=errors.NewUnauthorizedError("unauthenticated")
			c.JSON(restErr.Status, restErr)
			return
		}

		bookId, idErr := middlewares.GetId(c.Param("id"))
		if idErr != nil {
			c.JSON(idErr.Status, idErr)
			return
		}

		if role!="admin" {
			restErr:=errors.NewForbiddenError("only admin can return a book")
			c.JSON(restErr.Status, restErr)
			return
		}

		if err := service.BookServiceImpl.DeleteBook(bookId); err != nil {
			c.JSON(err.Status, err)
			return
		}
		c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
	}
}
