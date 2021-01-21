package books

import (
	"github.com/bhanupratap1810/bookstore_users-api/constants"
	"github.com/bhanupratap1810/bookstore_users-api/datasources/mysql"
	"github.com/bhanupratap1810/bookstore_users-api/utils/errors"
	"github.com/bhanupratap1810/bookstore_users-api/utils/mysql_utils"
)

type BookDaoService interface {
	Get(*Book) ([]Book, *errors.RestErr)
	Save(*Book) *errors.RestErr
	Update(*Book) *errors.RestErr
	Delete(*Book) *errors.RestErr
	//FindByUserId(*Book) *errors.RestErr
	FindByBookId(*Book) *errors.RestErr
}

type bookDaoMysql struct {
	DbService mysql.DbService
}

func NewBookDaoMysqlService(dbService mysql.DbService) BookDaoService {
	return &bookDaoMysql{
		DbService: dbService,
	}
}

func (b *bookDaoMysql) Get(*Book) ([]Book, *errors.RestErr){
	stmt, err := b.DbService.Client.Prepare(constants.QueryGetBook)
	if err != nil {
		//logger.Error("error when trying to prepare find users by status statement", err)
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		//logger.Error("error when trying to find users by status", err)
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]Book, 0)
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.BookId, &book.BookName, &book.BookAuthor, &book.BookType, &book.Status) ; err != nil {
			//logger.Error("error when scan user row into user struct", err)
			return nil, errors.NewInternalServerError(err.Error())
		}
		results = append(results, book)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError("no users matching role")
	}
	return results, nil
}

func (b *bookDaoMysql) Save(book *Book) *errors.RestErr{
	stmt, err := b.DbService.Client.Prepare(constants.QueryInsertBook)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(book.BookName, book.BookAuthor, book.BookType, book.Status)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}
	bookId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(saveErr)
	}
	book.BookId = bookId
	return nil
}
func (b *bookDaoMysql) Update(book *Book) *errors.RestErr{
	stmt, err := b.DbService.Client.Prepare(constants.QueryUpdateBook)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(book.BookName, book.BookAuthor, book.BookType, book.Status, book.BookId)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}
func (b *bookDaoMysql) Delete(book *Book) *errors.RestErr{
	stmt, err := b.DbService.Client.Prepare(constants.QueryDeleteBook)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	book.Status="deleted"
	if _, err = stmt.Exec(book.Status, book.BookId); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}
//func (b *bookDaoMysql) FindByUserId(book *Book) *errors.RestErr{
//	stmt, err := b.DbService.Client.Prepare(constants.QueryGetBookByUser)
//	if err != nil {
//		return errors.NewInternalServerError(err.Error())
//	}
//	defer stmt.Close()
//
//	result := stmt.QueryRow(book.BorrowerId)
//	if getErr := result.Scan(&book.BookId, &book.BookName, &book.BookAuthor, &book.BookType); getErr != nil {
//		return mysql_utils.ParseError(getErr)
//	}
//	return nil
//}
func (b *bookDaoMysql) FindByBookId(book *Book) *errors.RestErr{
	stmt, err := b.DbService.Client.Prepare(constants.QueryGetBookById)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(book.BookId)
	if getErr := result.Scan(&book.BookId, &book.BookName, &book.BookAuthor, &book.BookType, &book.Status); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}
	return nil
}