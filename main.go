package main

import (
	"log"

	"github.com/alexshd/total-coder-w17/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
