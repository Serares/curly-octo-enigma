package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"
)

type QuestionsHandler struct {
	Logger *slog.Logger
}

func NewQuestionsHandler(logger *slog.Logger) *QuestionsHandler {
	return &QuestionsHandler{
		Logger: logger,
	}
}

func (h *QuestionsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "Got the GET questions path %s", path)
		return
	}
}
