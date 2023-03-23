package admin

import (
	"time"

	"github.com/mluna-again/pregunta2/models"
)

type QuestionData struct {
	ID   int    `json:"id"`
	Body string `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *QuestionData) fromQuestion(question models.Question) {
	q.ID = int(question.ID)
	q.Body = question.Body
	q.CreatedAt = question.CreatedAt.Time
	q.UpdatedAt = question.UpdatedAt.Time
}
