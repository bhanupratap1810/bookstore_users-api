package users

import (
	"encoding/json"
	"fmt"
)

type PublicUser struct {
	Id          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	State     string `json:"state"`
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	State      string `json:"state"`
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}

func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:          user.Id,
			DateCreated: user.DateCreated,
			State:      user.State,
		}
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error in marshaling the user")
		return nil
	}
	var privateUser PrivateUser
	err = json.Unmarshal(userJson, &privateUser)
	if err != nil {
		fmt.Println("Error in unmarshalling the user")
		return nil
	}
	return privateUser
}