# ugo-cli - A cli client for [dumbergerl's server](https://github.com/DumbergerL/uno-server)

This cli client can be used to manually play on the DubmergerL's server. It's useful for developing bot strategies.

# Usage

## Prerequisites
- Install and setup [Go](https://golang.org/dl/) if you haven't yet.
- You will need an active internet connection when running/building for the first time since ugo-cli needs to download some libraries.

## Playing
Either run
`go build main.go`
and then execute resulting binary or just run
`go run main.go`.

By default ugo-cli will connect to `http://localhost:3000`, this can be changed with the `-host` and `-port` flags.

E.g. to connect to `http://10.0.0.1:1234` run `go run main.go -host=10.0.0.1 -port=1234`.
