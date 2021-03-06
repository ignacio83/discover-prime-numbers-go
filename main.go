package main

import (
	"fmt"
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
	inputFactory := ui.DiscoverPrimeNumbersFromUserInputFactory{Writer: os.Stdout, Reader: os.Stdin}
	command, err := inputFactory.Build(func(primeNumber uint64) {
		fmt.Printf("%d\n", primeNumber)
	})
	if err != nil {
		log.Fatal(err)
	}
	errExecute := command.Execute()
	if errExecute != nil {
		log.Fatal(errExecute)
	}
}
