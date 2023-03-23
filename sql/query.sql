-- name: GetQuestions :many
SELECT * FROM questions;

-- name: GetQuestionById :one
SELECT * FROM questions WHERE id = $1;

-- name: GetAnswers :many
SELECT * FROM answers WHERE question_id = ANY(@ids::bigint[])
ORDER BY created_at DESC;

-- name: InsertQuestion :one
INSERT INTO questions (body) VALUES ($1) RETURNING *;

-- name: InsertAnswers :copyfrom
INSERT INTO answers (body, question_id, is_correct) VALUES ($1, $2, $3);
