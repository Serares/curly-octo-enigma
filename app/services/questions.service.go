package services

import (
	"log/slog"

	"github.com/Serares/curly-octo-enigma/app/client"
	"github.com/Serares/curly-octo-enigma/domain/repo/dto"
)

type QuestionsService struct {
	Logger *slog.Logger
	Client *client.APIClient
}

func NewQuestionsService(logger *slog.Logger, client *client.APIClient) *QuestionsService {
	return &QuestionsService{
		Logger: logger,
		Client: client,
	}
}

func (s *QuestionsService) ListQuestions(token string) ([]dto.QuestionDTO, error) {
	questions, err := s.Client.ListQuestions(token)
	if err != nil {
		s.Logger.Error("Failed to list questions", slog.String("error", err.Error()))
		return nil, err
	}

	return questions, nil
}

func (s *QuestionsService) CreateQuestion(questionParams *dto.QuestionDTO, token string) error {
	return s.Client.CreateQuestion(questionParams, token)
}
