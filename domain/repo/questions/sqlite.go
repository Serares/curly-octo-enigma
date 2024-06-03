package questions

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/Serares/curly-octo-enigma/domain/repo/db"
	"github.com/Serares/curly-octo-enigma/domain/repo/dto"
	"github.com/Serares/curly-octo-enigma/domain/repo/utils"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

type SqliteQuestionsRepository struct {
	db           *db.Queries
	dbConnection *sql.DB
}

func NewSqliteQuestionsRepository() (*SqliteQuestionsRepository, error) {
	dbUrl, err := utils.CreateSqliteUrl()
	if err != nil {
		return nil, fmt.Errorf("error creating the connection string for database %w", err)
	}
	dbConn, err := sql.Open("libsql", dbUrl)
	if err != nil {
		return nil, err
	}
	dbConn.SetConnMaxIdleTime(30 * time.Minute)
	if err := dbConn.Ping(); err != nil {
		return nil, fmt.Errorf("error reaching the database %w", err)
	}
	dbQueries := db.New(dbConn)

	return &SqliteQuestionsRepository{
		db:           dbQueries,
		dbConnection: dbConn,
	}, nil
}

func (s *SqliteQuestionsRepository) FindAllQuestions(ctx context.Context) ([]dto.QuestionDTO, error) {
	return nil, errors.New("Not implemented")
}

func (s *SqliteQuestionsRepository) FindQuestionById(ctx context.Context, id string) (dto.QuestionDTO, error) {
	var wg sync.WaitGroup
	var asyncErrors = make([]error, 2)
	var question db.Question
	var answers []db.Answer
	wg.Add(2)

	go func() {
		defer wg.Done()
		var err error
		// get question
		question, err = s.db.GetQuestion(ctx, id)
		if err != nil {
			asyncErrors[0] = err
		}
	}()

	go func() {
		defer wg.Done()
		var err error
		// get answers
		answers, err = s.db.GetAnswersByQuestionID(ctx, id)
		if err != nil {
			asyncErrors[1] = err
		}
	}()
	wg.Wait()
	err := errors.Join(asyncErrors...)
	if err != nil {
		return dto.QuestionDTO{}, err
	}

	// map to dto
	dto := dto.QuestionDTO{
		Question: question,
		Answers:  make([]db.Answer, 0),
		User_ID:  question.UserSub,
	}
	if len(answers) > 0 {
		dto.Answers = append(dto.Answers, answers...)
	}
	return dto, nil
}

func (s *SqliteQuestionsRepository) CreateQuestion(question *dto.QuestionDTO) error {
	return errors.New("Not implemented")
}

func (s *SqliteQuestionsRepository) UpdateQuestion(question *dto.QuestionDTO) error {
	return errors.New("Not implemented")
}

func (s *SqliteQuestionsRepository) DeleteQuestionById(id string) error {
	return errors.New("Not implemented")
}

// answers

func (s *SqliteQuestionsRepository) AddAnswer(questionId string, answerParams dto.AnswerDTO) error {
	return errors.New("Not implemented")
}

func (s *SqliteQuestionsRepository) RemoveAnswer(questionId string, answerId string) error {
	return errors.New("Not implemented")
}

func (s *SqliteQuestionsRepository) UpdateAnswer(questionId string, answerParams dto.AnswerDTO) error {
	return errors.New("Not implemented")
}
