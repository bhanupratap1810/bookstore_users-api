package constants


const (
	MysqlUsersUsername = "mysql_users_username"
	MysqlUsersPassword = "mysql_users_password"
	MysqlUsersHost     = "mysql_users_host"
	MysqlUsersSchema   = "mysql_users_schema"
)

const (
	QueryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);"
	QueryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	QueryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?"
	QueryDeleteUser = "DELETE FROM users WHERE id=?;"
)