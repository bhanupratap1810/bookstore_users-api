package services

import (
	"github.com/bhanupratap1810/bookstore_users-api/domain/book_issue"
	"github.com/bhanupratap1810/bookstore_users-api/utils/errors"
)

type BookIssueService interface {
	CreateBookIssue(book_issue.BookIssue) (*book_issue.BookIssue, *errors.RestErr)
	GetBookIssue() (book_issue.BookIssues, *errors.RestErr)
	GetBookIssueByBook(int64) (*book_issue.BookIssue, *errors.RestErr)
	GetBookIssueByUser(int64) (*book_issue.BookIssue, *errors.RestErr)
	DeleteBook(int64) *errors.RestErr
}

type BookIssueServiceImpl struct {
	bookIssueDaoService book_issue.BookIssueDaoService
}

func NewBookIssueServiceImpl(bookIssueDaoService book_issue.BookIssueDaoService) BookIssueService {
	return &BookIssueServiceImpl{bookIssueDaoService: bookIssueDaoService}
}

func (bi *BookIssueServiceImpl) CreateBookIssue(bookIssue book_issue.BookIssue) (*book_issue.BookIssue, *errors.RestErr) {
	//if err := bookIssue.Validate(); err != nil {
	//	return nil, err
	//}
	if err := bi.bookIssueDaoService.Save(&bookIssue); err != nil {
		return nil, err
	}
	return &bookIssue, nil
}
func (bi *BookIssueServiceImpl) GetBookIssue() (book_issue.BookIssues, *errors.RestErr) {
	return bi.bookIssueDaoService.GetAll(&book_issue.BookIssue{})
}
func (bi *BookIssueServiceImpl) GetBookIssueByBook(bookId int64) (*book_issue.BookIssue, *errors.RestErr) {
	result := &book_issue.BookIssue{BookId: bookId}
	err := bi.bookIssueDaoService.GetByBookId(result)
	if err != nil {
		//log
		return nil, err
	}
	return result, nil
}
func (bi *BookIssueServiceImpl) GetBookIssueByUser(userId int64) (*book_issue.BookIssue, *errors.RestErr) {
	result := &book_issue.BookIssue{UserId: userId}
	err := bi.bookIssueDaoService.GetByUserId(result)
	if err != nil {
		//log
		return nil, err
	}
	return result, nil
}
func (bi *BookIssueServiceImpl) DeleteBook(userId int64) *errors.RestErr {
	bookIssue := &book_issue.BookIssue{BookId: userId}
	return bi.bookIssueDaoService.Delete(bookIssue)
}
