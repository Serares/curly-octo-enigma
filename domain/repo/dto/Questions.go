package dto

import "github.com/Serares/curly-octo-enigma/domain/repo/db"

type QuestionDTO struct {
	Question db.Question
	Answers  []db.Answer
	User_ID  string
}
