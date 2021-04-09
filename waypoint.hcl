project = "minikv"
app "server" {
  build {
    use "docker" {
      dockerfile = "./build/Dockerfile"
      disable_entrypoint = true
    }

    registry {
      use "docker" {
        image = "bjhayes/minikv"
        tag   = "latest"
      }
    }
  }

  deploy {
    use "docker" {}
  }

  url {
    auto_hostname = false
  }
}
