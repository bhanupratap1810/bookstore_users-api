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
		MysqlUsersUsername: os.Getenv(constants.Mysql_users_username),
		MysqlUsersPassword: os.Getenv(constants.Mysql_users_password),
		MysqlUsersHost:     os.Getenv(constants.Mysql_users_host),
		MysqlUsersSchema:   os.Getenv(constants.Mysql_users_schema),
	}
}
