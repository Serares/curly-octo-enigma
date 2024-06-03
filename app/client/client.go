package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/Serares/curly-octo-enigma/domain/repo/dto"
)

type APIEndpoints struct {
	CreateQuestion   func() string
	UpVoteQuestion   func(questionId string) string
	DownVoteQuestion func(questionId string) string
	GetQuestionById  func(qestionId string) string
	AddAnswer        func(questionId string) string
	UpVoteAnswer     func(questionId, answerId string) string
	DownVoteAnswer   func(questionId, answerId string) string
}

var endpoints = APIEndpoints{
	CreateQuestion:   func() string { return "/questions" },
	GetQuestionById:  func(questionId string) string { return fmt.Sprintf("/questions/%s", questionId) },
	UpVoteQuestion:   func(questionId string) string { return fmt.Sprintf("/questions/upvote/%s", questionId) },
	DownVoteQuestion: func(questionId string) string { return fmt.Sprintf("/questions/downvote/%s", questionId) },
	AddAnswer:        func(questionId string) string { return fmt.Sprintf("/questions/addAnswer/%s", questionId) },
	UpVoteAnswer: func(questionId, answerId string) string {
		return fmt.Sprintf("/answers?questionId=%s&answerId=%s", questionId, answerId)
	},
	DownVoteAnswer: func(questionId, answerId string) string {
		return fmt.Sprintf("/answers?questionId=%s&answerId=%s", questionId, answerId)
	},
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

func (c *APIClient) CreateQuestion(questionDto *dto.QuestionDTO, token string) error {
	// encode the question dto
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(questionDto); err != nil {
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
