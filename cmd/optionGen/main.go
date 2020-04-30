package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/timestee/optionGen"
)

// Globals
var (
	// Flags
	verbose              = flag.Bool("v", false, "Verbose - print lots of stuff")
	debug                = flag.Bool("debug", false, "debug")
	optionWithStructName = flag.Bool("option_with_struct_name", false, "should the option func with struct name?")
)

// usage prints the syntax and exists
func usage() {
	BaseName := path.Base(os.Args[0])
	fmt.Fprintf(os.Stderr,
		"Syntax: %s [flags] package_name parameter\n\n"+
			"Flags:\n\n",
		BaseName)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(optionGen.OptionGen + ": ")

	flag.Usage = usage
	flag.Parse()

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("unable to get working directory: %v", err)
	}
	optionGen.EnableDebug = *debug
	optionGen.Verbose = *verbose
	optionGen.ParseDir(wd, *optionWithStructName)
}
