package dto

import "github.com/Serares/curly-octo-enigma/domain/repo/db"

type QuestionDTO struct {
	Question db.Question `json:"question"`
	Answers  []db.Answer `json:"answers"`
	User_ID  string      `json:"user_id"`
}
