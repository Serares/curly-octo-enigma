package routes

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/Serares/curly-octo-enigma/app/client"
	"github.com/Serares/curly-octo-enigma/app/handlers"
	"github.com/Serares/curly-octo-enigma/app/middleware"
	"github.com/Serares/curly-octo-enigma/app/services"
	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

func Mux(log *slog.Logger) *http.ServeMux {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	redirectURL := os.Getenv("REDIRECT_URL")

	ctx := context.Background()

	oidcP, err := oidc.NewProvider(ctx, "https://accounts.google.com")
	if err != nil {
		log.Error("Failed to create OIDC provider", "error", err)
		os.Exit(1)
	}
	oauth2Cfg := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint:     oidcP.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	authService := services.NewAuthService(
		log,
		oauth2Cfg,
		oidcP,
	)

	apiClient := client.NewApiClient(log)

	qService := services.NewQuestionsService(
		log,
		apiClient,
	)

	m := http.NewServeMux()

	qHandler := handlers.NewQuestionsHandler(
		log,
		qService,
		authService,
	)
	authHandler := handlers.NewAuthHandler(log, authService)
	// views
	m.Handle("GET /login", authHandler)
	m.Handle("GET /callback", authHandler)
	// get all questions
	m.Handle("GET /questions",
		middleware.NewMiddleware(
			qHandler,
			middleware.WithSecure(false),
		))
	// single question by id
	m.Handle("GET /questions/{slug}",
		middleware.NewMiddleware(
			qHandler,
			middleware.WithSecure(false),
		))
	// add a question
	m.Handle("POST /questions",
		middleware.NewMiddleware(
			qHandler,
			middleware.WithSecure(false),
		))
	// create a answer for a question
	m.Handle("POST /answers/{questionId}",
		middleware.NewMiddleware(
			qHandler,
			middleware.WithSecure(false),
		))
	// m.Handle("GET /")
	m.Handle("GET /", middleware.NewMiddleware(
		qHandler,
		middleware.WithSecure(false),
	))
	return m
}
