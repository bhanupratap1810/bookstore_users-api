package services

import (
	"github.com/bhanupratap1810/bookstore_users-api/domain/users"
	"github.com/bhanupratap1810/bookstore_users-api/utils/crypto_utils"
	"github.com/bhanupratap1810/bookstore_users-api/utils/errors"
)

type UserService interface {
	GetUser(userId int64) (*users.User, *errors.RestErr)
	CreateUser(user users.User) (*users.User, *errors.RestErr)
	UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr)
	DeleteUser(userId int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
	LoginUser(users.LoginRequest) (string, *errors.RestErr)
}

type UserServiceImpl struct {
	userDaoService users.UserDaoService
}

func NewUserServiceImpl(userDaoService users.UserDaoService) UserService {
	return &UserServiceImpl{userDaoService: userDaoService}
	//var impl UserServiceImpl
	//impl.userDaoService=userDaoService
	//impl := UserServiceImpl{
	//	userDaoService: userDaoService,
	//}
	//impl:=UserServiceImpl{}
	//impl.userDaoService=userDaoService
}

func (u *UserServiceImpl) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	err := u.userDaoService.Get(result)
	if err != nil {
		//log
		return nil, err
	}
	return result, nil
}

func (u *UserServiceImpl) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	//return nil, nil
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := u.userDaoService.Save(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserServiceImpl) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := u.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	//current := &users.User{Id: user.Id}
	//err := u.userDaoService.Get(current)
	//if err != nil {
	//	//log
	//	return nil, err
	//}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Role != "" {
			current.Role = user.Role
		}
		if user.State != "" {
			current.State = user.State
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Role = user.Role
		current.State = user.State

	}
	//current in place of &user
	if err := u.userDaoService.Update(current); err != nil {
		return nil, err
	}
	return current, nil
}

func (u *UserServiceImpl) DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return u.userDaoService.Delete(user)
}

func (u *UserServiceImpl) SearchUser(role string) (users.Users, *errors.RestErr) {
	user := &users.User{Role: role}
	return u.userDaoService.FindByRole(user)
}

func (u *UserServiceImpl) LoginUser(request users.LoginRequest) (string, *errors.RestErr) {
	//var sharedKey = []byte("sercrethatmaycontainch@r$32chars")
	user := &users.User{
		Email:    request.Email,
		Password: crypto_utils.GetMd5(request.Password),
	}
	if err := u.userDaoService.FindByEmailAndPassword(user); err != nil {
		return "", err
	}

	token := string(NewJWTService().GenerateToken(user.Id, user.Role))
	//myClaims := map[string]interface{}{
	//	"user_id": user.Id,
	//	"user_role": user.Role,
	//}
	//token, _ := jwt.Sign(jwt.HS256, sharedKey, myClaims, jwt.MaxAge(15 * time.Minute))

	return token, nil
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
