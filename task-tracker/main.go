package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()

	args := flag.Args()
	validation(args)

	fmt.Println(len(args))

	fmt.Println(args)
}
