package users

import (
	"github.com/bhanupratap1810/bookstore_users-api/utils/crypto_utils"
	"github.com/bhanupratap1810/bookstore_users-api/utils/errors"
	"strings"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Password    string `json:"password"`
	Role 		string `json:"role"`
	State      	string `json:"state"`
}

type Users []User

func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" && len(user.Password)<8 {
		return errors.NewBadRequestError("invalid password")
	}
	user.Password = crypto_utils.GetMd5(user.Password)
	user.Role = strings.TrimSpace(strings.ToLower(user.Role))
	if user.Role == ""{
		user.Role = "employee"
	}
	user.State = strings.TrimSpace(strings.ToLower(user.State))
	if user.State == ""{
		user.State = "active"
	}
	return nil
}