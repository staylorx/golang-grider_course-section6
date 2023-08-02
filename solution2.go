package main

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

type logWriter struct{}

func main() {

	fileName := "example.txt"

	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}
	log.Info("Opening file ", fileName)

	f, err := os.Open(fileName)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	//io.Copy(os.Stdout, f)
	lw := logWriter{}
	io.Copy(lw, f)

}

// Custom logWriter that implements the Write inteface
func (logWriter) Write(bs []byte) (int, error) {
	log.Info(string(bs))
	return len(bs), nil
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
