.PHONY: build run compile waypoint

build:
	@echo Building docker image
	@docker build -f build/Dockerfile -t minikv .

run: build
	@docker run -d --rm minikv -p 8080:8080

compile:
	@echo Compiling minikv
	@go build -o bin/minikv ./cmd/server

waypoint:
	@echo Deploying to Waypoint
	waypoint up
