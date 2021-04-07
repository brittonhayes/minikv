FROM golang:1.16-alpine as build
WORKDIR /minikv
COPY . .
RUN go build -o /bin/minikv ./cmd/server

FROM scratch as runtime
COPY --from=build /bin/minikv /bin/minikv
ENTRYPOINT ["/bin/minikv"]
