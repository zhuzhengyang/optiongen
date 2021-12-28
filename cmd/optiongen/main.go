package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/timestee/optiongen"
)

// Globals
var (
	// Flags
	verbose              = flag.Bool("v", false, "Verbose, Deprecated use debug")
	debug                = flag.Bool("debug", false, "debug")
	optionWithStructName = flag.Bool("option_with_struct_name", false, "should the option func with struct name?")
	newFuncName          = flag.String("new_func", "", "new func name")
	xconf                = flag.Bool("xconf", false, "gen xconf tag")
	emptyCompositeNil    = flag.Bool("empty_composite_nil", false, "empty composite nil")
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
	log.SetPrefix(optiongen.OptionGen + ": ")
	flag.Usage = usage
	flag.Parse()

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("unable to get working directory: %v", err)
	}
	optiongen.EnableDebug = *debug
	optiongen.Verbose = *verbose
	optiongen.TagForXConf = *xconf
	optiongen.EmptyCompositenNil = *emptyCompositeNil
	optiongen.ParseDir(wd, *optionWithStructName, *newFuncName)
}
