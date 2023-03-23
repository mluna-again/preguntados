package admin

import (
	"context"

	"github.com/mluna-again/pregunta2/models"
)

func allQuestions(ctx context.Context) ([]QuestionData, error) {
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

func withAnswers(ctx context.Context, qd []QuestionData) ([]QuestionData, error) {
	mapppedAnswers := make(map[int64][]AnswerData)
	var ids []int64
	for _, q := range qd {
		ids = append(ids, q.ID)
	}

	answers, err := db.GetAnswers(ctx, ids)

	if err != nil {
		return qd, err
	}

	for _, a := range answers {
		ans := AnswerData{}
		ans.fromAnswer(a)
		mapppedAnswers[a.QuestionID] = append(mapppedAnswers[a.QuestionID], ans)
	}

	for i, q := range qd {
		qd[i].Answers = mapppedAnswers[q.ID]
	}

	return qd, nil
}

func createQuestion(ctx context.Context, q QuestionData) (QuestionData, error) {
	var question QuestionData

	tx, err := dbPool.Begin(ctx)
	if err != nil {
		return question, err
	}

	defer tx.Rollback(ctx)

	qtx := db.WithTx(tx)
	row, err := qtx.InsertQuestion(ctx, q.Body)
	if err != nil {
		return question, nil
	}

	question.fromQuestion(row)

	var answers []models.InsertAnswersParams
	for _, a := range q.Answers {
		ans := models.InsertAnswersParams{
			Body: a.Body,
			QuestionID: row.ID,
			IsCorrect: a.IsCorrect,
		}
		answers = append(answers, ans)
	}
	
	_, err = qtx.InsertAnswers(ctx, answers)
	if err != nil {
		return question, err
	}

	return question, nil
}

