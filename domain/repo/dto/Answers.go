package dto

import "time"

type AnswerDTO struct {
	Id         string
	QuestionId string
	Content    string
	Upvotes    int64
	Downvotes  int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
