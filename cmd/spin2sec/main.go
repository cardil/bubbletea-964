package main

import (
	"log"

	"github.com/cardil/repro-bubbletea-epoll-devnull-input/pkg/spinner"
)

func main() {
	// Code
	s := spinner.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
