#!/bin/bash

GOOSE_COMMAND=$1

# // check if goose_command is empty and print a message before exiting the script
if [ -z "$GOOSE_COMMAND" ]; then
    echo "Please provide a Goose command as an argument."
    exit 1
fi
# Check if Goose is installed
if ! command -v goose &>/dev/null; then
    echo "Goose could not be found, please install it before proceeding."
    exit 1
fi

CONNECTION_STRING=""
if $IS_LOCAL; then
    CONNECTION_STRING="${TURSO_DB_PROTOCOL}://${TURSO_DB_HOST}:${TURSO_DB_PORT}"
else
    CONNECTION_STRING="${TURSO_DB_PROTOCOL}://${TURSO_DB_NAME}.${TURSO_DB_HOST}?authToken=${TURSO_DB_TOKEN}"
fi

# Run Goose migrations
echo "Running Goose migrations..."
echo "Connection string: $CONNECTION_STRING"
cd "domain/schema"
goose turso "${CONNECTION_STRING}" $GOOSE_COMMAND

echo "Migration completed."
