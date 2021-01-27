package config

import (
	"github.com/bhanupratap1810/bookstore_users-api/constants"
	"os"
)

type Configuration struct {
	MysqlUsersUsername string
	MysqlUsersPassword string
	MysqlUsersHost     string
	MysqlUsersSchema   string
}

//pick from json file
func LoadConfig() *Configuration {

	return &Configuration{
		MysqlUsersUsername: os.Getenv(constants.MysqlUsersUsername),
		MysqlUsersPassword: os.Getenv(constants.MysqlUsersPassword),
		MysqlUsersHost:     os.Getenv(constants.MysqlUsersHost),
		MysqlUsersSchema:   os.Getenv(constants.MysqlUsersSchema),
	}
}
