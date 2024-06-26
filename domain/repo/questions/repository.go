package questions

import (
	"context"

	"github.com/Serares/curly-octo-enigma/domain/repo/dto"
)

type QuestionsRepostiroy interface {
	// questions
	ListAllQuestions(ctx context.Context) ([]dto.QuestionDTO, error)
	FindQuestionById(ctx context.Context, id string) (dto.QuestionDTO, error)
	CreateQuestion(ctx context.Context, question *dto.QuestionDTO) error
	UpdateQuestion(ctx context.Context, question *dto.QuestionDTO) error
	DeleteQuestionById(ctx context.Context, id string) error
	// answers
	AddAnswer(ctx context.Context, questionId string, answerParams *dto.AnswerDTO) error
	RemoveAnswer(ctx context.Context, answerId string) error
	UpdateAnswer(ctx context.Context, answerParams dto.AnswerDTO) error
	UpvoteAnswer(ctx context.Context, asnwerId string) error
	DownvoteAnswer(ctx context.Context, answerId string) error
}
