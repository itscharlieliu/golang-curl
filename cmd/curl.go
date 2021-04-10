package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	flag.Parse()

	args := flag.Args()

	resp, err := http.Get(args[1])

	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
