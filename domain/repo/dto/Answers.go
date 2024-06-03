package dto

import "time"

type AnswerDTO struct {
	Id         string    `json:"id"`
	QuestionId string    `json:"questionId"`
	Content    string    `json:"content"`
	Upvotes    int64     `json:"upvotes"`
	Downvotes  int64     `json:"downvotes"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
