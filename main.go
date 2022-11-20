package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("doesnt-exist.txt")
	if err != nil {
		log.Println(err) // log implements a simple logging package
	}
}