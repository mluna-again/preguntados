package admin

import (
	"time"

	"github.com/mluna-again/pregunta2/models"
)

type QuestionErrors struct {
	ID string `json:"id,omitempty"`
	Body string `json:"body,omitempty"`
	Answers []AnswerErrors `json:"answers,omitempty"`
	AnswersCount string `json:"answers_count,omitempty"`
	MoreThanOneCorrect string `json:"more_than_one_correct,omitempty"`
	DifferentQuestionIDs string `json:"different_question_ids,omitempty"`
}

type QuestionData struct {
	ID   int64    `json:"id"`
	Body string `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Answers []AnswerData `json:"answers"`
}

func (q *QuestionData) fromQuestion(question models.Question) {
	q.ID = question.ID
	q.Body = question.Body
	q.CreatedAt = question.CreatedAt.Time
	q.UpdatedAt = question.UpdatedAt.Time
}

type AnswerErrors struct {
	ID string `json:"id,omitempty"`
	Body string `json:"body,omitempty"`
}

type AnswerData struct {
	ID   int64    `json:"id"`
	QuestionID int64 `json:"question_id"`
	Body string `json:"body"`
	IsCorrect bool `json:"is_correct"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *AnswerData) fromAnswer(answer models.Answer) {
	a.ID = answer.ID
	a.QuestionID = answer.QuestionID
	a.Body = answer.Body
	a.CreatedAt = answer.CreatedAt.Time
	a.UpdatedAt = answer.UpdatedAt.Time
}
