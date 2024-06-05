package dto

import "github.com/Serares/curly-octo-enigma/domain/repo/db"

type QuestionDTO struct {
	Question     db.Question `json:"question"`
	Answers      []db.Answer `json:"answers"`
	AnswersCount int64       `json:"answersCount"`
	User_ID      string      `json:"user_id"`
	IsAuthor     bool        `json:"isAuthor"` // used to identify the author of the question
}
