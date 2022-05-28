package main

import (
	"flag"
	"os"

	"github.com/tmlbl/new"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path := flag.String("path", wd, "The path to a template")
	flag.Parse()
	inst, err := new.NewInstance(*path)
	if err != nil {
		panic(err)
	}
	err = inst.Prompt()
	if err != nil {
		panic(err)
	}
}
