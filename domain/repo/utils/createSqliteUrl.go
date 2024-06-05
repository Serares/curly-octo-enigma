package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Serares/curly-octo-enigma/shared/constants"
)

func CreateSqliteUrl() (string, error) {
	host := os.Getenv(constants.ENV_KEY_TURSO_DB_HOST)
	port := os.Getenv(constants.ENV_KEY_TURSO_DB_PORT)
	protocol := os.Getenv(constants.ENV_KEY_TURSO_DB_PROTOCOL)
	dbLocal := os.Getenv(constants.ENV_KEY_IS_DB_LOCAL)
	dbName := os.Getenv(constants.ENV_KEY_TURSO_DB_NAME)
	authToken := os.Getenv(constants.ENV_KEY_TURSO_DB_TOKEN)
	dbFile := os.Getenv(constants.ENV_KEY_DB_FILE)

	if dbFile != "" {
		// if db  file path is not empty then run the sqlite db file
		relativePath := filepath.Join(strings.Split(dbFile, "/")...)
		absPath, err := filepath.Abs(relativePath)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("file://%s", absPath), nil
	}

	// return a local sqlite url
	if dbLocal == "true" {
		return fmt.Sprintf("%s://%s:%s", protocol, host, port), nil
	}

	// return a turso url
	return fmt.Sprintf("%s://%s.%s?authToken=%s", protocol, dbName, host, authToken), nil
}
