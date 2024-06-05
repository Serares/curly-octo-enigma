package mux

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/Serares/curly-octo-enigma/api/muxIntegration/handlers"
	"github.com/Serares/curly-octo-enigma/api/muxIntegration/services"
	"github.com/Serares/curly-octo-enigma/domain/repo/questions"
	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

func Mux(log *slog.Logger) *http.ServeMux {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	ctx := context.Background()

	oidcP, err := oidc.NewProvider(ctx, "https://accounts.google.com")
	if err != nil {
		log.Error("Failed to create OIDC provider", "error", err)
		os.Exit(1)
	}
	oauth2Cfg := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     oidcP.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	qRepo, err := questions.NewSqliteQuestionsRepository()
	if err != nil {
		log.Error("error initializing the db connection", "error", err)
		os.Exit(1)
	}
	credService := services.NewCredentialsService(
		log,
		oauth2Cfg,
		oidcP,
	)

	qService := services.NewQuestionsService(
		log,
		credService,
		qRepo,
	)

	m := http.NewServeMux()

	qHandler := handlers.NewQuestionsHandler(
		log,
		qService,
	)

	// ⚠️ use middleware to guard paths
	m.Handle("GET /api/v1/questions", qHandler)
	m.Handle("GET /api/v1/questions/{slug}", qHandler)
	m.Handle("POST /api/v1/questions", qHandler)
	m.Handle("DELETE /api/v1/questions/{slug}", qHandler)

	m.Handle("GET /api/v1/answers", qHandler)
	// slug the id of the answer
	m.Handle("DELETE /api/v1/answers/{slug}", qHandler)
	// slug the id of the question
	m.Handle("POST /api/v1/answers/{slug}", qHandler)

	return m
}
