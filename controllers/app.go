package controllers

import "github.com/bhanupratap1810/bookstore_users-api/services"

type Service struct {
	UserServiceImpl      services.UserService
	BookServiceImpl      services.BookService
	BookIssueServiceImpl services.BookIssueService
	//reward service
	//payment service
}

func NewUserService(userServiceImpl services.UserService) Service {
	return Service{UserServiceImpl: userServiceImpl}
}

func NewBookService(bookServiceImpl services.BookService) Service {
	return Service{BookServiceImpl: bookServiceImpl,
	}
}

func NewBookIssueService(userServiceImpl services.UserService, bookServiceImpl services.BookService,
	bookIssueServiceImpl services.BookIssueService) Service {
	return Service{
		UserServiceImpl: userServiceImpl,
		BookServiceImpl: bookServiceImpl,
		BookIssueServiceImpl: bookIssueServiceImpl,
	}
}
