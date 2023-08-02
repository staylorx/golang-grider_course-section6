package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}
type triangle struct {
	base, height float64
}

func main() {
	t := triangle{base: 10.0, height: 10.0}
	printArea(t)

	s := square{sideLength: 10.0}
	printArea(s)

}

func printArea(s shape) {
	log.Infof("Area = %v", s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
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
