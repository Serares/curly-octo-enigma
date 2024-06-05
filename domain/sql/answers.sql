-- name: CreateAnswer :exec
INSERT INTO answers (
        id,
        question_id,
        user_sub,
        user_email,
        content,
        upvotes,
        downvotes,
        created_at,
        updated_at
    )
VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?);
-- name: DeleteAnswer :exec
DELETE FROM answers
WHERE id = ?;
-- name: GetAnswer :one
SELECT *
FROM answers
WHERE id = ?
LIMIT 1;
-- name: ListAnswers :many
SELECT *
FROM answers
ORDER BY created_at DESC;
-- name: UpvoteAnswer :exec
UPDATE answers
SET upvotes = upvotes + 1
WHERE id = ?;
-- name: DownvoteAnswer :exec
UPDATE answers
SET downvotes = downvotes + 1
WHERE id = ?;
-- name: GetAnswersByQuestionID :many
SELECT *
FROM answers
WHERE question_id = ?
ORDER BY created_at DESC;