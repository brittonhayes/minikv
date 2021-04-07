project = "minikv"
app "server" {
  build {
    use "docker" {
      dockerfile = "./build/Dockerfile"
    }
  }

  deploy {
    use "docker" {}
  }

  url {
    auto_hostname = false
  }
}
