package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http/httptest"
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/Serares/curly-octo-enigma/app/routes"
	"github.com/joho/godotenv"
)

func setupServer(t *testing.T) (string, func()) {
	t.Helper()
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("error loading the .env file")
	}
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	m := routes.Mux(log)

	ts := httptest.NewServer(m)
	return ts.URL, func() {
		log.Info("tearing down server")
		ts.Close()
	}
}

// 😉 debug server
func TestServer(t *testing.T) {
	url, cleanup := setupServer(t)

	fmt.Println("Test Server Started on: ", url)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	defer cleanup()
}
