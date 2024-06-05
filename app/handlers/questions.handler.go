package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/Serares/curly-octo-enigma/app/client"
	"github.com/Serares/curly-octo-enigma/app/middleware"
	"github.com/Serares/curly-octo-enigma/app/services"
	"github.com/Serares/curly-octo-enigma/app/views"
)

type QuestionsHandler struct {
	Logger           *slog.Logger
	QuestionsService *services.QuestionsService
	AuthService      *services.AuthService
}

func NewQuestionsHandler(
	logger *slog.Logger,
	questionService *services.QuestionsService,
	authService *services.AuthService,
) *QuestionsHandler {
	return &QuestionsHandler{
		Logger:           logger.WithGroup("Questions Handler"),
		QuestionsService: questionService,
		AuthService:      authService,
	}
}

func (h *QuestionsHandler) HandleQuestions(
	w http.ResponseWriter,
	r *http.Request,
	path []string,
) {
	rawToken := middleware.CheckAuth(r)
	slug := r.PathValue("slug")
	switch r.Method {
	case http.MethodGet:
		if slug == "" {
			questions, err := h.QuestionsService.ListQuestions(rawToken)
			viewQuestions(w, r, views.QuestionsProps{
				Questions: questions,
				Error:     err.Error(),
			})
			return
		} else if slug != "" {
			// go to single question page
			fmt.Fprintf(w, "<div>Method not implemented yet</div>")
		}
		return
	case http.MethodPost:
		r.ParseForm()
		questionBody := r.FormValue("question_body")
		questionTitle := r.FormValue("question_title")

		err := h.QuestionsService.CreateQuestion(
			&client.CreateQuestionRequest{
				Body:  questionBody,
				Title: questionTitle,
			},
			rawToken,
		)
		if err != nil {
			viewQuestions(
				w,
				r,
				views.QuestionsProps{
					Error: err.Error(),
				},
			)
			return
		}
		allQuestions, err := h.QuestionsService.ListQuestions(
			rawToken,
		)

		viewQuestions(
			w,
			r,
			views.QuestionsProps{
				Questions: allQuestions,
				Error:     err.Error(),
			},
		)

		fmt.Fprintf(w, "Got the POST questions path %s", path)
	case http.MethodDelete:
		fmt.Fprintf(w, "Delete method not implemented yet")
	}
}

func (h *QuestionsHandler) HandleAnswers(
	w http.ResponseWriter,
	r *http.Request,
	path []string,
) {
	// slug := r.PathValue("slug")
	// rawToken := middleware.CheckAuth(r)

	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "got to answers handler %s", path)
		return
	case http.MethodPost:
		fmt.Fprintf(w, "got to answers handler %s", path)
	case http.MethodDelete:
		fmt.Fprintf(w, "answers delete method")
	}
}

func (h *QuestionsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	if path[1] == "questions" {
		h.HandleQuestions(w, r, path)
		return
	} else if path[1] == "answers" {
		h.HandleAnswers(w, r, path)
		return
	}

	fmt.Fprintf(w, "<div>Path not found</div>")
}

func viewQuestions(
	w http.ResponseWriter,
	r *http.Request,
	props views.QuestionsProps,
) {
	views.Questions(props).Render(r.Context(), w)
}
