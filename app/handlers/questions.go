package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/Serares/curly-octo-enigma/app/client"
)

type QuestionsHandler struct {
	Logger *slog.Logger
	Client *client.APIClient
}

func NewQuestionsHandler(
	logger *slog.Logger,
	client *client.APIClient,
) *QuestionsHandler {
	return &QuestionsHandler{
		Logger: logger,
		Client: client,
	}
}

func (h *QuestionsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "Got the GET questions path %s", path)
		return
	case http.MethodPost:
		fmt.Fprintf(w, "Got the POST questions path %s", path)
	}
}
