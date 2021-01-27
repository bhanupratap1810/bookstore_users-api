package app

import (
	"github.com/bhanupratap1810/bookstore_users-api/config"
	"github.com/bhanupratap1810/bookstore_users-api/controllers"
	"github.com/bhanupratap1810/bookstore_users-api/datasources/mysql"
	"github.com/bhanupratap1810/bookstore_users-api/domain/book_issue"
	"github.com/bhanupratap1810/bookstore_users-api/domain/books"
	"github.com/bhanupratap1810/bookstore_users-api/domain/users"
	"github.com/bhanupratap1810/bookstore_users-api/services"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {

	//load config

	appConfig := config.LoadConfig()

	//initialize all the services which includes dao layer service layer any other third type client

	//Database service

	dbService, err := mysql.NewDbService(appConfig.MysqlUsersUsername, appConfig.MysqlUsersPassword, appConfig.MysqlUsersHost, appConfig.MysqlUsersSchema)
	if err != nil {
		panic("Database Connection Error")
		return
	}

	oAuthService := services.NewJWTService()

	userDaoService := users.NewUserDaoMysqlService(*dbService)
	userService := services.NewUserServiceImpl(userDaoService)
	UserService := controllers.NewUserService(userService)
	//mapUrls(controllers.Service{UserServiceImpl: userService})
	mapUsersUrls(UserService, oAuthService)

	bookDaoService := books.NewBookDaoMysqlService(*dbService)
	bookService := services.NewBookServiceImpl(bookDaoService)
	BookService := controllers.NewBookService(bookService)
	mapBooksUrls(BookService,oAuthService)

	bookIssueDaoService := book_issue.NewBookIssueDaoMysqlService(*dbService)
	bookIssueService := services.NewBookIssueServiceImpl(bookIssueDaoService)
	BookIssueService := controllers.NewBookIssueService(userService, bookService, bookIssueService,)
	mapBookIssueUrls(BookIssueService, oAuthService)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
