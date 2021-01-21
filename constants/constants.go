package constants


const (
	MysqlUsersUsername = "mysql_users_username"
	MysqlUsersPassword = "mysql_users_password"
	MysqlUsersHost     = "mysql_users_host"
	MysqlUsersSchema   = "mysql_users_schema"
)

const (
	QueryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created, password, role, state) VALUES(?,?,?,?,?,?,?);"
	QueryGetUser    = "SELECT id, first_name, last_name, email, date_created, role, state FROM users WHERE id=?;"
	QueryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=?, password=?, role=?, state=? WHERE id=?;"
	QueryDeleteUser = "UPDATE users SET state=? WHERE id=?;"
	QueryFindByRole = "SELECT id, first_name, last_name, email, date_created, role, state FROM users WHERE role=? AND state=?;"
	QueryFindByEmailAndPassword="SELECT id, first_name, last_name, email, date_created, role, state FROM users WHERE email=? AND password=? AND state=?"
)

const (
	QueryInsertBook = "INSERT INTO books(book_name, book_author, book_type, status) VALUES(?,?,?,?);"
	QueryGetBookById    = "SELECT book_id, book_name, book_author, book_type, status FROM books WHERE book_id=?;"
	//QueryGetBookByUser = "SELECT book_id, book_name, book_author, book_type FROM books WHERE borrower_id=?;"
	QueryUpdateBook = "UPDATE books SET book_name=?, book_author=?, book_type=?, status=? WHERE book_id=?;"
	QueryDeleteBook = "UPDATE books SET status=? WHERE book_id=?;"
	QueryGetBook = "SELECT book_id, book_name, book_author, book_type, status FROM books;"
)

const (
	QueryIssueBook = "INSERT INTO book_issue(book_id,user_id) VALUES(?,?);"
	QueryGetAll="SELECT issue_id, book_id, user_id FROM book_issue;"
	QueryGetIssueByBookId="SELECT issue_id, book_id, user_id FROM book_issue WHERE book_id=?;"
	QueryGetIssueByUserId="SELECT issue_id, book_id, user_id FROM book_issue WHERE user_id=?;"
	QueryDeleteIssue = "DELETE FROM book_issue WHERE user_id=?;"
)