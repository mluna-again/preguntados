-- name: GetQuestions :many
SELECT * FROM questions;

-- name: GetAnswers :many
SELECT * FROM answers WHERE question_id = ANY(@ids::bigint[])
ORDER BY created_at DESC;
