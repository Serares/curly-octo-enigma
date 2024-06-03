package questions

import (
	"context"

	"github.com/Serares/curly-octo-enigma/domain/repo/dto"
)

type QuestionsRepostiroy interface {
	// questions
	FindAllQuestions(ctx context.Context) ([]dto.QuestionDTO, error)
	FindQuestionById(ctx context.Context, id string) (dto.QuestionDTO, error)
	CreateQuestion(ctx context.Context, question *dto.QuestionDTO) (dto.QuestionDTO, error)
	UpdateQuestion(ctx context.Context, question *dto.QuestionDTO) (dto.QuestionDTO, error)
	DeleteQuestionById(ctx context.Context, id string) error
	// answers
	AddAnswer(ctx context.Context, questionId string, answerParams dto.AnswerDTO) error
	RemoveAnswer(ctx context.Context, questionId string, answerId string) error
	UpdatedAnswer(ctx context.Context, questionId string, answerParams dto.AnswerDTO) error
}
