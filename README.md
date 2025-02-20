Demonstration repository for the issue: https://github.com/vektah/gqlparser/issues/347

## Steps to reproduce

1. Clone this repository
2. Run `go run server.go`
3. Run `make failing-request` . Observe that no default value was set

At the same time `make success-request` works as expected and default value is set.