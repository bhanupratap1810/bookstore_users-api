package book_issue

type BookIssue struct {
	IssueId int64 `json:"issue_id"`
	BookId  int64 `json:"book_id"`
	UserId  int64 `json:"user_id"`
}

type BookIssues []BookIssue

