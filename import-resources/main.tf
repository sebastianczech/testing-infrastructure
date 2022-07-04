terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 2.16"
    }
  }
  required_version = ">= 0.14"
}

provider "docker" {
  host = "unix:///var/run/docker.sock"
}