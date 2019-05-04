package main

import (
	"log"
	"os"

	"github.com/iwittkau/nions"
	"github.com/iwittkau/nions/rocket"
)

func main() {

	sys := os.Getenv("NIONS_SYSTEM")
	var err error
	switch sys {
	case "":
		panic("environment variable NIONS_SYSTEM should not be empty")
	case nions.SystemRocket:
		err = rocket.Exec()
	case nions.SystemSlack:
	}

	if err != nil {
		log.Fatal(err)
	}
}
