package admin

import (
	"context"

	"github.com/mluna-again/pregunta2/models"
)

func allQuestions(ctx context.Context, db *models.Queries) ([]QuestionData, error) {
	var qds []QuestionData

	questions, err := db.GetQuestions(ctx)

	if err != nil {
		return qds, err
	}

	for _, q := range questions {
		var qd QuestionData
		qd.fromQuestion(q)
		qds = append(qds, qd)
	}

	return qds, nil
}
