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

func NewBookIssueDaoMysqlService(dbService mysql.DbService) BookIssueDaoService {
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
	stmt, err := bi.DbService.Client.Prepare(constants.QueryGetAll)
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

	results := make([]BookIssue, 0)
	for rows.Next() {
		var bookIssue BookIssue
		if err := rows.Scan(&bookIssue.IssueId, &bookIssue.BookId, &bookIssue.UserId) ; err != nil {
			//logger.Error("error when scan user row into user struct", err)
			return nil, errors.NewInternalServerError(err.Error())
		}
		results = append(results, bookIssue)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError("no users matching role")
	}
	return results, nil
}

func (bi *bookIssueDaoMysql) GetByBookId(bookIssue *BookIssue) *errors.RestErr {
	stmt, err := bi.DbService.Client.Prepare(constants.QueryGetIssueByBookId)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(bookIssue.BookId)
	if getErr := result.Scan(&bookIssue.IssueId, &bookIssue.BookId, &bookIssue.UserId); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}
	return nil
}

func (bi *bookIssueDaoMysql) GetByUserId(bookIssue *BookIssue) *errors.RestErr {
		stmt, err := bi.DbService.Client.Prepare(constants.QueryGetIssueByUserId)
		if err != nil {
			return errors.NewInternalServerError(err.Error())
		}
		defer stmt.Close()

		result := stmt.QueryRow(bookIssue.UserId)
		if getErr := result.Scan(&bookIssue.IssueId, &bookIssue.BookId, &bookIssue.UserId); getErr != nil {
			return mysql_utils.ParseError(getErr)
		}
		return nil
}

func (bi *bookIssueDaoMysql) Delete(bookIssue *BookIssue) *errors.RestErr {
		stmt, err := bi.DbService.Client.Prepare(constants.QueryDeleteIssue)
		if err != nil {
			return errors.NewInternalServerError(err.Error())
		}
		defer stmt.Close()

		if _, err = stmt.Exec(bookIssue.UserId); err != nil {
			return mysql_utils.ParseError(err)
		}
		return nil
}