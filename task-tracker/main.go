package main

import (
	"flag"

	validator "github.com/problem-setter/task-tracker/internal"
)

func main() {

	flag.Usage = validator.PrintHelp
	flag.Parse()

	args := flag.Args()

	validator.Validation(args)

	// fmt.Println(len(args))
	// fmt.Println(args)
}
