-- name: GetQuestions :many
SELECT * FROM questions;

-- name: GetAnswers :many
SELECT * FROM answers WHERE question_id = $1
ORDER BY created_at DESC
LIMIT 4;
