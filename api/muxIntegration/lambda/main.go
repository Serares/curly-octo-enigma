package main

import (
	"log/slog"
	"os"

	"github.com/Serares/curly-octo-enigma/api/muxIntegration/mux"
	"github.com/akrylysov/algnhsa"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	m := mux.Mux(log)
	algnhsa.ListenAndServe(m, nil)
}
