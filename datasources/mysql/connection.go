package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DbService struct {
	Client *sql.DB
}

func NewDbService(username, password, host, schema string) (*DbService, error) {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		schema)
	client, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("err connecting to mysql : ", err)
		return nil, err
	}
	if err = client.Ping(); err != nil {
		//add print statements
		return nil, err
	}

	return &DbService{Client: client}, nil
}

//const (
//	mysql_users_username = "mysql_users_username"
//	mysql_users_password = "mysql_users_password"
//	mysql_users_host     = "mysql_users_host"
//	mysql_users_schema   = "mysql_users_schema"
//)

//var (
//	Client   *sql.DB
//	username = os.Getenv(mysql_users_username)
//	password = os.Getenv(mysql_users_password)
//	host     = os.Getenv(mysql_users_host)
//	schema   = os.Getenv(mysql_users_schema)
//)
//
//func init() {
//	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
//		username,
//		password,
//		host,
//		schema)
//	var err error
//	Client, err = sql.Open("mysql", dataSourceName)
//	if err != nil {
//		panic(err)
//	}
//	if err = Client.Ping(); err != nil {
//		panic(err)
//	}
//	log.Println("database successfully configured")
//}
