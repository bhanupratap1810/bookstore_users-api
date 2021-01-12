package controllers

import "github.com/bhanupratap1810/bookstore_users-api/services"

type Service struct {
	UserServiceImpl services.UserService
	//reward service
	//payment service
}

func NewService(userServiceImpl services.UserService) *Service {
	return &Service{UserServiceImpl: userServiceImpl}
}
