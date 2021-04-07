# minikv

> A mini key value store built with Go + Gorilla Mux + BoltDB

## ⚡ Usage 

### 🐋 Docker (Recommended) 

Build the alpine Docker container and start the KV store.

```shell
make run
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
