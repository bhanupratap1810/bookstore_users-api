package app

import (
	"github.com/bhanupratap1810/bookstore_users-api/controllers"
	"github.com/bhanupratap1810/bookstore_users-api/controllers/book_issue"
	"github.com/bhanupratap1810/bookstore_users-api/controllers/books"
	"github.com/bhanupratap1810/bookstore_users-api/controllers/middlewares"
	"github.com/bhanupratap1810/bookstore_users-api/controllers/ping"
	"github.com/bhanupratap1810/bookstore_users-api/controllers/users"
	"github.com/bhanupratap1810/bookstore_users-api/services"
)

func mapUsersUrls(service controllers.Service, authService services.JWTService) {
	router.GET("/ping", ping.Ping)
	router.POST("/user/login", users.LoginUserHandler(service))
	usersRouterGroup := router.Group("users")
	{
		usersRouterGroup.POST("", users.CreateUserHandler(service))
		verifiedUser := usersRouterGroup.Use(middlewares.VerifyAndServe(authService))
		verifiedUser.GET("/:user_id", users.GetUserHandler(service))
		verifiedUser.PUT("/:user_id", users.UpdateUserHandler(service))
		verifiedUser.PATCH("/:user_id", users.UpdateUserHandler(service))
		verifiedUser.DELETE("/:user_id", users.DeleteUserHandler(service))

	}
	router.Use(middlewares.VerifyAndServe(authService)).
		GET("/internal/users/search", users.SearchUserHandler(service))
	//verifiedUser:=router.Use()
	//verifiedUser.POST("users", users.CreateUserHandler(service))
	//router.GET("users/:user_id", users.GetUserHandler(service))
	//router.PUT("/users/:user_id", users.UpdateUserHandler(service))
	//router.PATCH("/users/:user_id", users.UpdateUserHandler(service))
	//router.DELETE("/users/:user_id", users.DeleteUserHandler(service))
}

func mapBooksUrls(service controllers.Service, authService services.JWTService) {
	booksRouterGroup := router.Group("books")
	{
		verifiedUser := booksRouterGroup.Use(middlewares.VerifyAndServe(authService))
		verifiedUser.GET("", books.GetBooksHandler(service))
		verifiedUser.POST("", books.CreateBookHandler(service))
		verifiedUser.GET("/:book_id", books.SearchBookByIdHandler(service))
		verifiedUser.PUT("/:book_id", books.UpdateBookHandler(service))
		verifiedUser.PATCH("/:book_id", books.UpdateBookHandler(service))
		verifiedUser.DELETE("/:book_id", books.DeleteBookHandler(service))

	}
	//router.GET("books", books.GetBooksHandler(service))
	//router.POST("books", books.CreateBookHandler(service))
	//router.PUT("/books/:book_id", books.UpdateBookHandler(service))
	//router.PATCH("/books/:book_id", books.UpdateBookHandler(service))
	//router.DELETE("/books/:book_id", books.DeleteBookHandler(service))
	//router.GET("/books/:book_id", books.SearchBookByIdHandler(service))
	//router.GET("/user/books/:user_id", books.SearchBookByUserHandler(service))
}

func mapBookIssueUrls(service controllers.Service, authService services.JWTService) {
	bookIssueRouterGroup := router.Group("book_issue")
	{
		verifiedUser := bookIssueRouterGroup.Use(middlewares.VerifyAndServe(authService))
		verifiedUser.POST("", book_issue.CreateBookIssueHandler(service))
		verifiedUser.GET("",book_issue.GetBookIssueHandler(service))
		verifiedUser.GET("book/:id",book_issue.GetBookIssueByBookHandler(service))
		verifiedUser.GET("user/:id",book_issue.SearchBookIssueByUserHandler(service))
		verifiedUser.DELETE("delete/:id",book_issue.DeleteBookHandler(service))

	}
}
