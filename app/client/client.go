package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/Serares/curly-octo-enigma/domain/repo/dto"
)

var (
	NotAuthorized = errors.New("Not Authorized")
	NotFound      = errors.New("Not Found")
)

type APIEndpoints struct {
	CreateQuestion   func() string
	ListQuestions    func() string
	UpVoteQuestion   func(questionId string) string
	DownVoteQuestion func(questionId string) string
	GetQuestionById  func(qestionId string) string
	AddAnswer        func(questionId string) string
	DeleteAnswer     func(answerId string) string
	UpVoteAnswer     func(answerId string) string
	DownVoteAnswer   func(answerId string) string
}

var endpoints = APIEndpoints{
	CreateQuestion:   func() string { return "/questions" },
	ListQuestions:    func() string { return "/questions" },
	GetQuestionById:  func(questionId string) string { return fmt.Sprintf("/questions/%s", questionId) },
	UpVoteQuestion:   func(questionId string) string { return fmt.Sprintf("/questions/upvote/%s", questionId) },
	DownVoteQuestion: func(questionId string) string { return fmt.Sprintf("/questions/downvote/%s", questionId) },
	AddAnswer:        func(questionId string) string { return fmt.Sprintf("/answers/%s", questionId) },
	DeleteAnswer:     func(answerId string) string { return fmt.Sprintf("/answers/%s", answerId) },
	UpVoteAnswer: func(answerId string) string {
		return fmt.Sprintf("/answers/%s", answerId)
	},
	DownVoteAnswer: func(answerId string) string {
		return fmt.Sprintf("/answers%s", answerId)
	},
}

type CreateQuestionRequest struct {
	Body  string `json:"body"`
	Title string `json:"title"`
}

type CreateAnswerRequest struct {
	Body  string `json:"body"`
	Title string `json:"title"`
}

// used to interact with gw api
type APIClient struct {
	Logger  *slog.Logger
	Client  *http.Client
	BaseUrl string
}

func NewApiClient(log *slog.Logger) *APIClient {
	baseApiUrl := os.Getenv("BASE_API_URL")

	return &APIClient{
		Logger: log.WithGroup("Api Client"),
		Client: &http.Client{
			Transport: http.DefaultTransport,
			Timeout:   30 * time.Second,
		},
		BaseUrl: baseApiUrl,
	}
}

func (ssrc *APIClient) sendRequest(
	url,
	method,
	contentType,
	token string,
	expStatus int,
	body io.Reader,
) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return []byte{}, err
	}
	if contentType == "" {
		contentType = "applicaiton/json"
	}
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", token) // 🤔 this has got to be the raw token
	r, err := ssrc.Client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer r.Body.Close()
	msg, err := io.ReadAll(r.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("cannot read body: %w", err)
	}
	if r.StatusCode != expStatus {
		err = fmt.Errorf("error invalid response")
		if r.StatusCode == http.StatusNotFound {
			err = fmt.Errorf("not found")
		}
		return []byte{}, fmt.Errorf("%w: %s", err, msg)
	}

	return msg, nil
}

func (c *APIClient) CreateQuestion(questionParams *CreateQuestionRequest, token string) error {
	// encode the question dto
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(questionParams); err != nil {
		c.Logger.Error("error encoding question dto", "error", err)
		return err
	}

	_, err := c.sendRequest(
		fmt.Sprintf("%s%s", c.BaseUrl, endpoints.CreateQuestion()),
		http.MethodPost,
		"application/json",
		token,
		http.StatusOK,
		&body,
	)

	if err != nil {
		c.Logger.Error("error creating question", "error", err)
		return err
	}

	return nil
}

func (c *APIClient) ListQuestions(token string) ([]dto.QuestionDTO, error) {
	var questionsResponse []dto.QuestionDTO
	resp, err := c.sendRequest(
		fmt.Sprintf("%s%s", c.BaseUrl, endpoints.ListQuestions()),
		http.MethodGet,
		"application/json",
		token,
		http.StatusOK,
		nil,
	)

	if err != nil {
		c.Logger.Error("error requesting the quesitons", "error", err)
		return nil, err
	}

	err = json.Unmarshal(resp, &questionsResponse)
	if err != nil {
		c.Logger.Error("error listing questions", "error", err)
		return nil, err
	}

	return questionsResponse, nil
}

func (c *APIClient) DeleteQuestionById(questionId, token string) error {
	_, err := c.sendRequest(
		fmt.Sprintf("%s%s", c.BaseUrl, endpoints.GetQuestionById(questionId)),
		http.MethodDelete,
		"application/json",
		token,
		http.StatusOK,
		nil,
	)

	if err != nil {
		c.Logger.Error("error deleting question", "error", err)
		return err
	}

	return nil
}

func (c *APIClient) UpvoteQuestion(questionId, token string) error {
	_, err := c.sendRequest(
		fmt.Sprintf("%s%s", c.BaseUrl, endpoints.UpVoteQuestion(questionId)),
		http.MethodPut,
		"application/json",
		token,
		http.StatusOK,
		nil,
	)

	if err != nil {
		c.Logger.Error("error upvoting question", "error", err)
		return err
	}

	return nil
}

func (c *APIClient) AddAnswer(questionId string, answerParams CreateAnswerRequest, token string) error {
	// encode the question dto
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(answerParams); err != nil {
		c.Logger.Error("error encoding answer dto", "error", err)
		return err
	}

	_, err := c.sendRequest(
		fmt.Sprintf("%s%s", c.BaseUrl, endpoints.AddAnswer(questionId)),
		http.MethodPost,
		"application/json",
		token,
		http.StatusOK,
		&body,
	)

	if err != nil {
		c.Logger.Error("error adding answer", "error", err)
		return err
	}

	return nil
}

func (c *APIClient) UpvoteAnswer(questionId, answerId, token string) error {
	_, err := c.sendRequest(
		fmt.Sprintf("%s%s", c.BaseUrl, endpoints.UpVoteAnswer(answerId)),
		http.MethodPut,
		"application/json",
		token,
		http.StatusOK,
		nil,
	)

	if err != nil {
		c.Logger.Error("error upvoting answer", "error", err)
		return err
	}

	return nil
}
