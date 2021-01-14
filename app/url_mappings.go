package app

import (
	"github.com/bhanupratap1810/bookstore_users-api/controllers"
	"github.com/bhanupratap1810/bookstore_users-api/controllers/ping"
	"github.com/bhanupratap1810/bookstore_users-api/controllers/users"
)

func mapUrls(service controllers.Service){
	router.GET("/ping", ping.Ping)

	router.POST("users", users.CreateUserHandler(service))
	router.GET("users/:user_id", users.GetUserHandler(service))
	router.PUT("/users/:user_id", users.UpdateUserHandler(service))
	router.PATCH("/users/:user_id", users.UpdateUserHandler(service))
	router.DELETE("/users/:user_id", users.DeleteUserHandler(service))
}