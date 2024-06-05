#!/bin/bash
DOCKER=docker
DOCKER_PATH=$(command -v $DOCKER)

if [ -z "$DOCKER_PATH" ]; then
    echo "Docker not found. Checking for podman."

    DOCKER=podman
    PODMAN_PATH=$(command -v podman)
    if [ -z "$PODMAN_PATH" ]; then
        echo "Podman not found. No suitable container engine."
        exit 127
    fi
fi

# Build the Docker images for the Go modules
$DOCKER build -t api_module_image ./api/muxIntegration
$DOCKER build -t app_module_image ./app

sleep 5

# Run the Go module containers
$DOCKER run -d \
    --name go_api_container \
    --env-file ./api/.env.local \
    api_module_image

$DOCKER run -d \
    --name go_app_container \
    --env-file ./app/.env.local \
    app_module_image
