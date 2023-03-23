package admin

import (
	"time"

	"github.com/mluna-again/pregunta2/models"
)

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

type AnswerData struct {
	ID   int64    `json:"id"`
	QuestionID int64 `json:"question_id"`
	Body string `json:"body"`
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
