package routes

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
)

func Mux(log *slog.Logger) *http.ServeMux {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Error("error trying to initialize the cfg")
		os.Exit(1)
	}

}
