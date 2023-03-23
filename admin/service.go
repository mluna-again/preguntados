package admin

import (
	"context"
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
