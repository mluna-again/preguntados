package admin

func validateQuestion(q QuestionData) (QuestionErrors, bool) {
	hasErrors := false
	var errors QuestionErrors
	if q.Body == "" {
		errors.Body = "Body is required"
		hasErrors = true
	}

	if len(q.Answers) != 4 {
		errors.AnswersCount = "Must have 4 answers"
		hasErrors = true
	}

	for _, a := range q.Answers {
		answerErrors, hasAnswerErrors := validateAnswer(a)
		if hasAnswerErrors {
			errors.Answers = append(errors.Answers, answerErrors)
			hasErrors = true
		}
	}

	if !onlyOneCorrectAnswer(q.Answers) {
		errors.MoreThanOneCorrect = "Must have (only) one correct answer"
		hasErrors = true
	}

	return errors, hasErrors
}

func validateAnswer(a AnswerData) (AnswerErrors, bool) {
	hasErrors := false
	var errors AnswerErrors
	if a.Body == "" {
		errors.Body = "Body is required"
		hasErrors = true
	}

	return errors, hasErrors
}

func onlyOneCorrectAnswer(a []AnswerData) bool {
	count := 0
	for _, answer := range a {
		if answer.IsCorrect {
			count++
		}
	}

	return count == 1
}
