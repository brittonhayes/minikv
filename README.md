# minikv

> A mini key value store built with Go + Gorilla Mux + BoltDB

## ⚡ Usage 

### 🐋 Docker (Recommended) 

Fetch the Docker container and start the KV store.

```shell
docker run -d --rm bjhayes/minikv:latest -p 8080:8080
```

### 🇼 Hashicorp Waypoint

Build the docker image, tag it, and deploy it to your local environment, Kubernetes, or AWS. 
Just tweak the [waypoint.hcl](./waypoint.hcl) to your needs!

```shell
make waypoint
```

### 🐹 Go 

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
