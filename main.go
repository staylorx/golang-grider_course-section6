package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", body)
}

func init() {
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		lvl = "debug"
	}

	ll, err := log.ParseLevel(lvl)
	if err != nil {
		ll = log.DebugLevel
	}

	log.SetLevel(ll)
}
