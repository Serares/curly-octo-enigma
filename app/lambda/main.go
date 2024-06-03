package main

import (
	"os"

	"log/slog"

	"github.com/Serares/curly-octo-enigma/app/routes"
	"github.com/akrylysov/algnhsa"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	m := routes.Mux(log)

	algnhsa.ListenAndServe(m, nil)
}
