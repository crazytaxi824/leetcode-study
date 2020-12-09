package main

import (
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile)
	log.SetOutput(os.Stdout)

	log.Println("hello world")
}
