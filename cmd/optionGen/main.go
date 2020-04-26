package main

import (
	"github.com/timestee/optionGen"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix(optionGen.OptionGen + ": ")
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("unable to get working directory: %v", err)
	}
	optionGen.ParseDir(wd)
}
