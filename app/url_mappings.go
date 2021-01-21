package app

import (
	"github.com/bhanupratap1810/bookstore_users-api/controllers"
	"github.com/bhanupratap1810/bookstore_users-api/controllers/books"
	"github.com/bhanupratap1810/bookstore_users-api/controllers/ping"
	"github.com/bhanupratap1810/bookstore_users-api/controllers/users"
)

func mapUsersUrls(service controllers.Service) {
	router.GET("/ping", ping.Ping)

	router.POST("users", users.CreateUserHandler(service))
	router.GET("users/:user_id", users.GetUserHandler(service))
	router.PUT("/users/:user_id", users.UpdateUserHandler(service))
	router.PATCH("/users/:user_id", users.UpdateUserHandler(service))
	router.DELETE("/users/:user_id", users.DeleteUserHandler(service))
	router.GET("/internal/users/search", users.SearchUserHandler(service))
	router.POST("/user/login", users.LoginUserHandler(service))
}

func mapBooksUrls(service controllers.Service){
	router.GET("books", books.GetBooksHandler(service))
	router.POST("books", books.CreateBookHandler(service))
	router.PUT("/books/:book_id", books.UpdateBookHandler(service))
	router.PATCH("/books/:book_id", books.UpdateBookHandler(service))
	router.DELETE("/books/:book_id", books.DeleteBookHandler(service))
	router.GET("/books/:book_id", books.SearchBookByIdHandler(service))
	router.GET("/user/books/:user_id", books.SearchBookByUserHandler(service))
}