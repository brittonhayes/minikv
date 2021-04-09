FROM golang:1.16-alpine AS build
WORKDIR /minikv
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -ldflags="-w -s" -a -installsuffix cgo -o app ./cmd/server

FROM scratch
COPY --from=build /minikv/app /bin/minikv
ENV ENVIRONMENT="docker"

VOLUME /data

ENTRYPOINT ["/bin/minikv"]

EXPOSE 8080
