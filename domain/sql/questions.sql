-- name: CreateQuestion :exec
insert into questions (
        id,
        created_at,
        updated_at,
        user_sub,
        user_email,
        user_name,
        upvotes,
        downvotes,
        title,
        body
    )
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
-- name: GetQuestion :one
SELECT *
FROM questions
WHERE id = ?
LIMIT 1;
-- name: ListQuestions :many
SELECT *
FROM questions
ORDER BY created_at DESC
LIMIT ? OFFSET ?;
-- name: UpdateQuestion :exec
UPDATE questions
SET updated_at = ?,
    title = ?,
    body = ?
WHERE id = ?
    AND user_sub = ?;
-- name: DeleteQuestion :exec
DELETE FROM questions
WHERE id = ?;
-- name: UpvoteQuestion :exec
UPDATE questions
SET upvotes = upvotes + 1
WHERE id = ?;
-- name: DownvoteQuestion :exec
UPDATE questions
SET downvotes = downvotes + 1
WHERE id = ?;
-- name: GetQuestionWithCountAnswers :one
SELECT q.id,
    COUNT(a.id) AS number_of_answers
FROM questions q
    LEFT JOIN answers a ON q.id = a.question_id
WHERE q.id = ?
GROUP BY q.id;
-- name: GetQuestionWithAnswers :many
SELECT *
FROM questions q
    LEFT JOIN answers a ON q.id = a.question_id
WHERE q.id = ?
ORDER BY a.created_at DESC;