// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package models

import (
	"context"
)

const getAnswers = `-- name: GetAnswers :many
SELECT id, question_id, body, is_correct, created_at, updated_at FROM answers WHERE question_id = $1
ORDER BY created_at DESC
LIMIT 4
`

func (q *Queries) GetAnswers(ctx context.Context, questionID int64) ([]Answer, error) {
	rows, err := q.db.Query(ctx, getAnswers, questionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Answer
	for rows.Next() {
		var i Answer
		if err := rows.Scan(
			&i.ID,
			&i.QuestionID,
			&i.Body,
			&i.IsCorrect,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getQuestions = `-- name: GetQuestions :many
SELECT id, body, created_at, updated_at FROM questions
`

func (q *Queries) GetQuestions(ctx context.Context) ([]Question, error) {
	rows, err := q.db.Query(ctx, getQuestions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Question
	for rows.Next() {
		var i Question
		if err := rows.Scan(
			&i.ID,
			&i.Body,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
