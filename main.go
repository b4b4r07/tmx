package main

import (
	"log"
	"os"
)

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	tmux, err := NewTmux()
	if err != nil {
		log.Print(err)
		return 1
	}
	if err := tmux.Run(args); err != nil {
		log.Print(err)
		return 1
	}
	return 0
}
