package app

import (
	"github.com/bhanupratap1810/bookstore_users-api/config"
	"github.com/bhanupratap1810/bookstore_users-api/controllers"
	"github.com/bhanupratap1810/bookstore_users-api/datasources/mysql"
	"github.com/bhanupratap1810/bookstore_users-api/domain/users"
	"github.com/bhanupratap1810/bookstore_users-api/services"
	"github.com/gin-gonic/gin"
)

var(
	router=gin.Default()
)

func StartApplication(){



	//load config

	appConfig := config.LoadConfig()

	//initialize all the services which includes dao layer service layer any other third type client

	//Database service

	dbService, err := mysql.NewDbService(appConfig.MysqlUsersUsername, appConfig.MysqlUsersPassword, appConfig.MysqlUsersHost, appConfig.MysqlUsersSchema, )
	if err != nil {
		//handle
		panic(err)
	}

	userDaoService := users.NewUserDaoMysqlService(*dbService)
	userService := services.NewUserServiceImpl(userDaoService)

	service := controllers.NewService(userService)

	mapUrls(service)




	router.Run(":8080")
}