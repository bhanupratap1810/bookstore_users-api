package book_issue

import (
	"github.com/bhanupratap1810/bookstore_users-api/constants"
	"github.com/bhanupratap1810/bookstore_users-api/datasources/mysql"
	"github.com/bhanupratap1810/bookstore_users-api/utils/errors"
	"github.com/bhanupratap1810/bookstore_users-api/utils/mysql_utils"
)

type BookIssueDaoService interface {
	Save(*BookIssue) *errors.RestErr
	GetAll(*BookIssue) ([]BookIssue, *errors.RestErr)
	GetByBookId(*BookIssue) *errors.RestErr
	GetByUserId(*BookIssue) *errors.RestErr
	Delete(*BookIssue) *errors.RestErr
}

type bookIssueDaoMysql struct {
	DbService mysql.DbService
}

func NewBookDaoMysqlService(dbService mysql.DbService) BookIssueDaoService {
	return &bookIssueDaoMysql{
		DbService: dbService,
	}
}

func (bi *bookIssueDaoMysql) Save(bookIssue *BookIssue) *errors.RestErr {
	stmt, err := bi.DbService.Client.Prepare(constants.QueryIssueBook)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(bookIssue.BookId, bookIssue.UserId)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}
	issueId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(saveErr)
	}
	bookIssue.IssueId = issueId
	return nil
}

func (bi *bookIssueDaoMysql) GetAll(*BookIssue) ([]BookIssue, *errors.RestErr) {
	return nil, nil
}

func (bi *bookIssueDaoMysql) GetByBookId(*BookIssue) *errors.RestErr {
	return nil
}

func (bi *bookIssueDaoMysql) GetByUserId(*BookIssue) *errors.RestErr {
	return nil
}

func (bi *bookIssueDaoMysql) Delete(*BookIssue) *errors.RestErr {
	return nil
}