package main

import (
	"log"

	"github.com/stdyum/api-journal/internal"
)

func main() {
	log.Fatalf("error launching web server %s", internal.App().Error())
}
