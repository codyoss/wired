# wired

Wired is meant to be an example project that shows off how to combine a handful of Go "basics", along with some more
advanced topics, to build a "real" application.

## Concepts used in this repo

- building a http.Client and calling a RESTful api
- making a service layer for business logic
- caching results of service calls
- how to write a basic test
- how to write a table driven test
- how to use build tags to break out integration tests
- serving JSON responses over HTTP
- having a multi-binary project
- make a CLI
- auto wiring dependencies with compile time code generation
- functional programming; passing behavior
- fan-out work with goroutines

## Build

The below commands build binaries and copy them back to the root directory.

### api

`(cd cmd/api/ && go build && cp api ../..)`

### cli

`(cd cmd/swctl/ && go build && cp swctl ../..)`

## Run tests

This command below turns off test caching with `-count=1` and enables running the integration tests as well with `-tags=integration`.
`go test -count=1 -tags=integration ./...`

## TODO

- [ ] make file
- [ ] basic configuration
- [ ] better test coverage
