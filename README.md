# testing-infrastructure

Repository created while learning how to test infrastructure

## Docs

* [Terratest examples](https://terratest.gruntwork.io/examples/)
* [Terratest quickstart](https://terratest.gruntwork.io/docs/getting-started/quick-start/)

## Basic example - how to run tests

```bash
cd basic-example/infra
terraform apply -var-file varfile.tfvars

cd basic-example/test
go mod init github.com/sebastianczech/testing-infrastructure/tree/main/basic-example
go mod tidy
go test -v -timeout 30m
```

## Importing existing resources

```bash
cd import-resources

docker run --name hashicorp-learn --detach --publish 8080:80 nginx:latest
docker ps --filter="name=hashicorp-learn"

terraform import docker_container.web $(docker inspect --format="{{.ID}}" hashicorp-learn)
terraform plan
terraform show
terraform show -no-color >> docker.tf
terraform apply -auto-approve

terraform state list
terraform state rm docker_container.web
```