.PHONY: build run compile waypoint

build:
	@echo Building docker image
	@docker build -f build/Dockerfile -t minikv .

run: build
	docker run --publish 8080:8080 --name minikv --rm minikv

compile:
	@echo Compiling minikv
	@go build -ldflags="-w -s" -o bin/minikv ./cmd/minikv

waypoint:
	@echo Deploying to Waypoint
	waypoint up
