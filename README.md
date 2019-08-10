# uno-cli - A cli client for [dumbergerl's uno server](https://github.com/DumbergerL/uno-server)

This cli client can be used to play uno on the Uno bot challenge's server. It's useful for developing uno bot strategies.

# Usage

## Prerequisites
- Install and setup [Go](https://golang.org/dl/) if you haven't yet.
- You will need an active internet connection when running/building for the first time since uno-cli needs to download some libraries.

## Playing
Either run
`go build main.go`
and then execute resulting binary or just run
`go run main.go`.

By default uno-cli will connect to `http://localhost:3000`, this can be changed with the `-host` and `-port` flags.

E.g. to connect to `http://10.0.0.1:1234` run `go run main.go -host=10.0.0.1 -port=1234`.

## Uno Bot Challenge
The Uno Bot Challenge is a small competition, where one has to develop a Uno playing bot within two hours. In the end the bots compete against each other.

**bots that participated**

* [my bot](https://github.com/staubichsauger/jabberwoky) (Language: Go - of course ;-) )
* [dumbergerl](https://github.com/DumbergerL/uno-bot) (Language: JavaScript / Node.js)
* revilo196 (Language: Go)
