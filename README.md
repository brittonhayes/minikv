# minikv

> A mini key value store built with Go + Gorilla Mux + BoltDB

## ⚡ Usage

### 🐋 Docker (Recommended)

**Size**: `6.5 MB` 🔬

Fetch the Docker container and start the KV store. The Docker image is only `~6.5MB` in size. Thanks to a scratch
container, and some helpful go build tags that remove excess information from the binary.

```shell
docker run -d --rm bjhayes/minikv:latest -p 8080:8080
```

```shell
# Using a custom port
docker run -d --rm bjhayes/minikv:latest -e PORT=":9090" -p 9090:9090
```

### 🇼 Hashicorp Waypoint

Build the docker image, tag it, and deploy it to your local environment, Kubernetes, or AWS. Just tweak
the [waypoint.hcl](./waypoint.hcl) to your needs!

```shell
make waypoint
```

### 🐹 Go

**Size**: `6.2 MB` 🔬

Compile and run the Go executable

```shell
make compile
./bin/minikv
```

## ✨ Try it out

```shell
# Ping the server to see if it's up
curl http://localhost:8080/status/health

# Add an entry
curl -d '{"name":"bob"}' http://localhost:8080/mykey

# Read an entry
curl http://localhost:8080/mykey
```
