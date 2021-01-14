package services

import (
	"github.com/bhanupratap1810/bookstore_users-api/domain/users"
	"github.com/bhanupratap1810/bookstore_users-api/utils/errors"
)

type UserService interface {
	GetUser(userId int64) (*users.User, *errors.RestErr)
	CreateUser(user users.User) (*users.User, *errors.RestErr)
	UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr)
	DeleteUser(userId int64) *errors.RestErr
}

type UserServiceImpl struct {
	UserDaoService users.UserDaoService
}

func NewUserServiceImpl(userDaoService users.UserDaoService) UserService {
	return &UserServiceImpl{UserDaoService: userDaoService}
}

func (u *UserServiceImpl) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	err := u.UserDaoService.Get(result)
	if err != nil {
		//log
		return nil, err
	}
	return result, nil
}

func (u *UserServiceImpl) CreateUser(user users.User) (*users.User, *errors.RestErr){
	//return nil, nil
		if err := user.Validate(); err != nil {
			return nil, err
		}
		if err := u.UserDaoService.Save(&user); err != nil {
			return nil, err
		}
		return &user, nil
}

func (u *UserServiceImpl) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr){
		//current, err := UserService.GetUser(u, user.Id)
		//if err != nil {
		//	return nil, err
		//}

	current := &users.User{Id: user.Id}
	err := u.UserDaoService.Get(current)
	if err != nil {
		//log
		return nil, err
	}

		if isPartial {
			if user.FirstName != "" {
				current.FirstName = user.FirstName
			}
			if user.LastName != "" {
				current.LastName = user.LastName
			}
			if user.Email != "" {
				current.Email = user.Email
			}
		} else {
			current.FirstName = user.FirstName
			current.LastName = user.LastName
			current.Email = user.Email

		}

		if err := u.UserDaoService.Update(&user); err != nil {
			return nil, err
		}
		return current, nil
}

func (u *UserServiceImpl) DeleteUser(userId int64) *errors.RestErr{
		user := &users.User{Id: userId}
		return u.UserDaoService.Delete(user)
}

//func GetUser(userId int64) (*users.User, *errors.RestErr) {
//	result := &users.User{Id: userId}
//	if err := result.Get(); err != nil {
//		return nil, err
//	}
//	return result, nil
//}
//
//func CreateUser(user users.User) (*users.User, *errors.RestErr) {
//	if err := user.Validate(); err != nil {
//		return nil, err
//	}
//	if err := user.Save(); err != nil {
//		return nil, err
//	}
//	return &user, nil
//}
//
//func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
//	current, err := GetUser(user.Id)
//	if err != nil {
//		return nil, err
//	}
//
//	if isPartial {
//		if user.FirstName != "" {
//			current.FirstName = user.FirstName
//		}
//		if user.LastName != "" {
//			current.LastName = user.LastName
//		}
//		if user.Email != "" {
//			current.Email = user.Email
//		}
//	} else {
//		current.FirstName = user.FirstName
//		current.LastName = user.LastName
//		current.Email = user.Email
//
//	}
//
//	if err := current.Update(); err != nil {
//		return nil, err
//	}
//	return current, nil
//}
//
//func DeleteUser(userId int64) *errors.RestErr {
//	user := &users.User{Id: userId}
//	return user.Delete()
//}
