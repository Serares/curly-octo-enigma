package handlers

import (
	"log/slog"
	"net/http"
	"strings"

	middelware "github.com/Serares/curly-octo-enigma/api/muxIntegration/middleware"
	"github.com/Serares/curly-octo-enigma/api/muxIntegration/services"
	"github.com/Serares/curly-octo-enigma/utils"
)

type QuestionsHandler struct {
	Logger           *slog.Logger
	QuestionsService *services.QuestionsService
}

func NewQuestionsHandler(
	log *slog.Logger,
	questionsServie *services.QuestionsService,
) *QuestionsHandler {
	return &QuestionsHandler{
		Logger:           log.WithGroup("Mux questions handler"),
		QuestionsService: questionsServie,
	}
}

func (h *QuestionsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	path := strings.Split(r.URL.Path, "/") // for /api/v1 -> {"", "api", "v1"}
	token := middelware.CheckAuthHeader(r)
	if token == "" {
		h.Logger.Error("Token not found in the Authorization header")
	}
	switch r.Method {
	case http.MethodGet:
		if path[3] == "questions" {
			if slug == "" {
				// return all questions
				questions, err := h.QuestionsService.ListAllQuestions()
				if err != nil {
					utils.ReplyError(w, r, http.StatusInternalServerError, "Error trying to get a list of questions")
				}
				utils.ReplySuccess(w, r, http.StatusOK, questions)
				return
			} else {
				// return question by question id
			}
		}
	}
}
