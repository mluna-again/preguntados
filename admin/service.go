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
	for i, q := range qd {
		answers, err := db.GetAnswers(ctx, q.ID)

		if err != nil {
			return qd, err
		}

		for _, a := range answers {
			var ad AnswerData
			ad.fromAnswer(a)
			qd[i].Answers = append(qd[i].Answers, ad)
		}
	}

	return qd, nil
}
