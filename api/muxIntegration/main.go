package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/Serares/curly-octo-enigma/api/muxIntegration/mux"
	"github.com/joho/godotenv"
)

func main() {
	// used for local testing
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("error loading the .env file")
	}

	port := os.Getenv("PORT")
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	m := mux.Mux(log)
	server := &http.Server{
		Addr:         fmt.Sprintf("localhost:%s", port),
		Handler:      m,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}
	fmt.Printf("Listening on %v\n", server.Addr)
	server.ListenAndServe()
}
