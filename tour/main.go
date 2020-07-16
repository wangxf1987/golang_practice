package main

import (
	"golang_practice/tour/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute error: %v", err)
	}
}
