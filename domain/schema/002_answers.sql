-- +goose Up
CREATE TABLE IF NOT EXISTS answers (
    id TEXT PRIMARY KEY NOT NULL,
    question_id TEXT NOT NULL,
    content TEXT NOT NULL,
    upvotes INTEGER NOT NULL,
    downvotes INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (question_id) REFERENCES questions(id) ON DELETE CASCADE
);
-- +goose Down
DROP TABLE answers;