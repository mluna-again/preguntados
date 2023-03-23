package admin

func validateQuestionForCreate(q QuestionData) (QuestionErrors, bool) {
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

	ansErrCount := 0
	errors.Answers = make([]AnswerErrors, 4)
	for i, a := range q.Answers {
		answerErrors, hasAnswerErrors := validateAnswer(a)
		if hasAnswerErrors {
			errors.Answers[i] = answerErrors
			hasErrors = true
			ansErrCount++
		}
	}

	if !onlyOneCorrectAnswer(q.Answers) {
		errors.MoreThanOneCorrect = "Must have (only) one correct answer"
		hasErrors = true
	}

	if ansErrCount == 0 {
		errors.Answers = nil
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

func validateQuestionForUpdate(q QuestionData) (QuestionErrors, bool) {
	errors, hasErrors := validateQuestionForCreate(q)

	if hasErrors {
		return errors, true
	}

	valid := true
	if q.ID == 0 {
		errors.ID = "no id"
		valid = false
	}

	if errors.Answers == nil {
		errors.Answers = make([]AnswerErrors, 4)
	}
	ansErrCount := 0
	for i, a := range q.Answers {
		if a.QuestionID == 0 {
			ansErrCount++
			errors.Answers[i].ID = "no id"
			valid = false
		}

		if a.QuestionID != q.ID {
			valid = false
			errors.DifferentQuestionIDs = "not all answers have the same question_id"
		}
	}

	if ansErrCount == 0 {
		errors.Answers = nil
	}

	return errors, !valid
}
