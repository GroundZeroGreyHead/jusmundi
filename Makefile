# Define the application name
APP_NAME := JUSMUNDI
MAIN_FILE := main.go
OUTPUT_BINARY := main
# Docker settings
DOCKER_IMAGE := $(APP_NAME):latest

# Define phony targets to avoid conflicts with file names
.PHONY: build run clean test dev build-and-run build-docker run-docker

# Define the target for building the application
build:
	go build -o $(OUTPUT_BINARY) $(MAIN_FILE)

# Define the target for running the application
run: build
	./$(OUTPUT_BINARY)

# Define the target for cleaning up any compiled binary
clean:
	rm -f $(OUTPUT_BINARY)

# Define the target for running tests
test:
	go test -v ./...

# Define the target for running the application in development mode
dev:
	go run $(MAIN_FILE)

# Define the target for building and running the application in a single step
build-and-run: build run

# Target to build the Docker image
build-docker:
	docker build -t $(DOCKER_IMAGE) .

# Target to run the application in a Docker container
run-docker: build-docker
	docker run $(DOCKER_IMAGE)
