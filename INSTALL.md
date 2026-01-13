# INSTALL

We support several installation methods.

## RUNTIME REQUIREMENTS

* [Go](https://go.dev/) 1.25.4+

# PRECOMPILED BINARIES

https://github.com/mcandre/factorio/releases

## Instructions

1. Download release archive.
2. Extract archive.
3. Select executables for your target platform.
4. Copy executabless to a convenient location, e.g. `$HOME/bin`.
5. Ensure location is registered in `$PATH`.

# DOCKER

## Requirements

* [Docker](https://www.docker.com/) 28.0.1+

## Instructions

```sh
docker pull n4jm4/factorio
```

# BUILD FROM SOURCE

## Requirements

* Ensure `GOBIN` is registered in `$PATH`. Validate like `go env GOBIN; echo "$PATH"`

## Instructions

```sh
go get -tool github.com/mcandre/factorio/src/cmd/factorio
go mod tidy
go mod vendor
go install tool
```

For more information on developing factorio itself, see [DEVELOPMENT.md](DEVELOPMENT.md).
