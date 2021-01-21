package controllers

import "github.com/bhanupratap1810/bookstore_users-api/services"

type Service struct {
	UserServiceImpl services.UserService
	BookServiceImpl services.BookService
	//reward service
	//payment service
}

func NewUserService(userServiceImpl services.UserService) Service {
	return Service{UserServiceImpl: userServiceImpl}
}

func NewBookService(bookServiceImpl services.BookService, userServiceImpl services.UserService) Service {
	return Service{BookServiceImpl: bookServiceImpl,
		UserServiceImpl: userServiceImpl,
	}
}







