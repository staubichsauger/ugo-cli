package main

import (
	"flag"
	"github.com/staubichsauger/ugo-cli/cli"
	"log"
	"net/url"
	"strconv"
	"strings"
)

func main() {
	host := flag.String("host", "localhost", "ugo server host")
	port := flag.Int("port", 3000, "ugo server port")
	flag.Parse()

	url, err := url.Parse("http://" + *host + ":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatal("Invalid hostname and port supplied: ", err)
	}

	player := cli.Client{
		Url: *url,
	}

	err = player.Login()
	if err != nil {
		log.Fatal("error logging in: " + err.Error())
	}

	stop := make(chan error)

	go player.Play(stop)

	for err := range stop {
		log.Print(err)
		if strings.Contains(err.Error(), "Status: 500") {
			return
		}
	}
}
