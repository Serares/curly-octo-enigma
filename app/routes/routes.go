package routes

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/Serares/curly-octo-enigma/domain/repo/questions"
	"github.com/aws/aws-sdk-go-v2/config"
)

func Mux(log *slog.Logger) *http.ServeMux {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Error("error trying to initialize the cfg")
		os.Exit(1)
	}
	questionsRepo, err := questions.NewSqliteQuestionsRepository()
	if err != nil {
		log.Error("error trying to initialize the questions repo")
		os.Exit(1)
	}
}
