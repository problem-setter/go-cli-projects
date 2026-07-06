package main

import (
	"flag"
	"fmt"

	validator "github.com/problem-setter/task-tracker/internal"
)

func main() {
	flag.Usage = validator.PrintHelp
	flag.Parse()

	args := flag.Args()

	if err := validator.Validation(args); err != nil {
		fmt.Printf("%v\n", err)
	}
}
