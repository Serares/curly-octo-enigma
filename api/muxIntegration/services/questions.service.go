package services

import (
	"context"
	"log/slog"

	"github.com/Serares/curly-octo-enigma/api/muxIntegration/utils"
	"github.com/Serares/curly-octo-enigma/domain/repo/db"
	"github.com/Serares/curly-octo-enigma/domain/repo/dto"
	"github.com/Serares/curly-octo-enigma/domain/repo/questions"
	"github.com/google/uuid"
)

type QuestionsService struct {
	Logger      *slog.Logger
	CredService *CredentialsService
	Questions   questions.QuestionsRepostiroy
}

func NewQuestionsService(
	logger *slog.Logger,
	credService *CredentialsService,
	questions questions.QuestionsRepostiroy,
) *QuestionsService {
	return &QuestionsService{
		Logger:      logger.WithGroup("MuxQuestions Service"),
		CredService: credService,
		Questions:   questions,
	}
}

func (s *QuestionsService) ListAllQuestions() ([]dto.QuestionDTO, error) {
	return s.Questions.ListAllQuestions(context.Background())
}

func (s *QuestionsService) GetSingleQuestion(token, questionId string) (dto.QuestionDTO, error) {
	claims, err := s.CredService.GetTokenClaims(context.Background(), token)
	if err != nil {
		s.Logger.Error("failed to get the token claims")
		return dto.QuestionDTO{}, err
	}

	questionDto, err := s.Questions.FindQuestionById(context.Background(), questionId)
	if err != nil {
		s.Logger.Error("failed to get the question by id", "error", err)
		return dto.QuestionDTO{}, err
	}

	questionDto.IsAuthor = utils.CheckIfAuthor(&claims, questionDto.Question.UserSub)
	return questionDto, nil
}

func (s *QuestionsService) CreateQuestion(token string, questionParams dto.QuestionDTO) error {
	claims, err := s.CredService.GetTokenClaims(context.Background(), token)
	if err != nil {
		s.Logger.Error("failed to get the token claims")
		return err
	}
	dto := dto.QuestionDTO{
		Question: db.Question{
			ID:        uuid.New().String(),
			UserSub:   claims.Sub,
			Title:     questionParams.Question.Title,
			Body:      questionParams.Question.Body,
			UserName:  claims.Name,
			UserEmail: claims.Email,
		},
	}
	return s.Questions.CreateQuestion(context.Background(), &dto)
}

// should have a guard on token and user id
func (s *QuestionsService) DeleteQuestionById(questionId string) error {
	return s.Questions.DeleteQuestionById(context.Background(), questionId)
}

func (s *QuestionsService) CreateAnswer(token string, questionId string, answerParams dto.AnswerDTO) error {
	claims, err := s.CredService.GetTokenClaims(context.Background(), token)
	if err != nil {
		s.Logger.Error("failed to get the token claims")
		return err
	}
	dto := dto.AnswerDTO{
		Id:         uuid.New().String(),
		QuestionId: questionId, // this is passed as param also lol
		UserSub:    claims.Sub,
		UserEmail:  claims.Email,
		Content:    answerParams.Content,
	}
	return s.Questions.AddAnswer(context.Background(), questionId, &dto)
}
