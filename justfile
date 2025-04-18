# Displays available recipes by running `just -l`.
setup:
  #!/usr/bin/env bash
  just -l

# Run Go tests (short mode)
test-unit:
  go test -short $(go list ./... \
    | grep -v '^github\.com/ethereum/go-ethereum/cmd/')


# Run go build to make sure Nibiru can import this dependency
test-build:
  go build ./...
