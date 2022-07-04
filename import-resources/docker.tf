# resource "docker_image" "nginx" {
#   name = "nginx:latest"
# }

# resource "docker_container" "web" {
#   name  = "hashicorp-learn"

#   image = docker_image.nginx.latest

#   ports {
#     external = 8081
#     internal = 80
#   }
# }

resource "docker_container" "web" {
  name  = "hashicorp-learn"
  image = "sha256:cd4e03b35a8e938f7bf1257fbd19fc2d7becae39974c0328236fd7df49c0a92a"

  env  = []

  ports {
    external = 8081
    internal = 80
  }
}

