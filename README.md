# testing-infrastructure

Repository created while learning how to test infrastructure using materials:
* [Terratest examples](https://terratest.gruntwork.io/examples/)
* [Terratest quickstart](https://terratest.gruntwork.io/docs/getting-started/quick-start/)

## Commands

```bash
go mod init github.com/sebastianczech/testing-infrastructure/tree/main/basic-example
go mod tidy
go test -v -timeout 30m
```