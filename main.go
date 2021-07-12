package main

import (
	"ignacio83/discover-prime-numbers-go/ui"
	"log"
	"os"
)

const logPrefix = "discover-prime-numbers-go: "

func init() {
	log.SetPrefix(logPrefix)
	log.SetFlags(0)
}

func main() {
	inputFactory := ui.DiscoverPrimeNumbersCommandFromUserInputFactory{Writer: os.Stdout, Reader: os.Stdin}
	command, err := inputFactory.Build()
	if err != nil {
		log.Fatal(err)
	}
	command.Execute()
}