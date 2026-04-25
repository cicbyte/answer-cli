package tui

import "github.com/cicbyte/answer-cli/internal/models"

type viewKind int

const (
	viewQuestionList viewKind = iota
	viewQuestionDetail
)

type questionsLoadedMsg struct {
	questions []models.QuestionListItem
	count     int64
	page      int
}

type questionDetailLoadedMsg struct {
	question *models.QuestionInfoResp
}

type answersLoadedMsg struct {
	answers []models.AnswerInfo
	count   int64
}

type commentsLoadedMsg struct {
	objectID string
	comments []models.CommentInfo
}

type searchResultsMsg struct {
	results []*models.SearchResult
	count   int64
	query   string
}

type authCheckedMsg struct {
	username string
}

type apiErrorMsg struct {
	err error
}
