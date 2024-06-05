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
		fmt.Println("the db url:", dbUrl)
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

// returns a general list of questions
func (s *SqliteQuestionsRepository) ListAllQuestions(ctx context.Context) ([]dto.QuestionDTO, error) {
	dbQuestions, err := s.db.ListQuestionsWithCounts(ctx)
	if err != nil {
		return nil, err
	}

	dtoQuestions := make([]dto.QuestionDTO, 0)

	if len(dbQuestions) < 1 {
		return nil, errors.New("no questions found")
	}

	for _, q := range dbQuestions {
		dto := dto.QuestionDTO{
			Question: db.Question{
				ID:        q.ID,
				CreatedAt: q.CreatedAt,
				UpdatedAt: time.Time{},
				UserSub:   q.UserSub,
				UserName:  "",
				UserEmail: "",
				Upvotes:   q.Upvotes,
				Downvotes: q.Downvotes,
				Title:     q.Title,
				Body:      "",
			},
			Answers:      []db.Answer{},
			AnswersCount: q.Count,
			User_ID:      q.UserSub,
			IsAuthor:     false, // this has to be populated in service
		}
		dtoQuestions = append(dtoQuestions, dto)
	}

	return dtoQuestions, nil
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

func (s *SqliteQuestionsRepository) CreateQuestion(ctx context.Context, question *dto.QuestionDTO) error {
	params := db.CreateQuestionParams{
		ID:        question.Question.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserSub:   question.User_ID,
		Title:     question.Question.Title,
		Body:      question.Question.Body,
		UserEmail: question.Question.UserEmail,
		UserName:  question.Question.UserName,
		Upvotes:   0,
		Downvotes: 0,
	}

	return s.db.CreateQuestion(ctx, params)
}

func (s *SqliteQuestionsRepository) UpdateQuestion(ctx context.Context, question *dto.QuestionDTO) error {
	return errors.New("not implemented")
}

func (s *SqliteQuestionsRepository) DeleteQuestionById(ctx context.Context, id string) error {
	return s.db.DeleteQuestion(ctx, id)
}

// answers
func (s *SqliteQuestionsRepository) AddAnswer(ctx context.Context, questionId string, answerParams *dto.AnswerDTO) error {
	dbParams := db.CreateAnswerParams{
		ID:         answerParams.Id,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		QuestionID: questionId,
		Content:    answerParams.Content,
		Upvotes:    0,
		Downvotes:  0,
	}
	return s.db.CreateAnswer(context.Background(), dbParams)
}

func (s *SqliteQuestionsRepository) RemoveAnswer(ctx context.Context, answerId string) error {
	return s.db.DeleteAnswer(ctx, answerId)
}

func (s *SqliteQuestionsRepository) UpdateAnswer(ctx context.Context, answerParams dto.AnswerDTO) error {
	return errors.New("not implemented")
}

func (s *SqliteQuestionsRepository) DownvoteAnswer(ctx context.Context, answerId string) error {
	return errors.New("not implemented")
}

func (s *SqliteQuestionsRepository) UpvoteAnswer(ctx context.Context, asnwerId string) error {
	return errors.New("not implemented")
}
