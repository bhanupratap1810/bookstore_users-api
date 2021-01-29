# bookstore_users-api
Users API

1. Clone the github repo. in src

2. Run the mysql server.

3. Run the Query:
SCHEMA `users_db` DEFAULT CHARACTER SET utf8 COLLATE utf8_spanish_ci ;

CREATE TABLE `users_db`.`users` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
  `first_name` VARCHAR(45) NULL,
  `last_name` VARCHAR(45) NULL,
  `email` VARCHAR(45) NOT NULL,
  `date_created` VARCHAR(45) NULL,
  PRIMARY KEY (`id`));

4. Set the env variables using:

export mysql_users_username=<username>
  
export mysql_users_password=<password>
  
export mysql_users_host=<address on which mysql server is running>
  
export mysql_users_schema=users_db
  
5. go get -u github.com/gin-gonic/gin
6. go mod init
7. go mod vendor
8. go run main.go
