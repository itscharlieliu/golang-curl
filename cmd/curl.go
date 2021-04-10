package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	flag.Parse()

	args := flag.Args()

	fmt.Println(args[0])

	resp, err := http.Get(args[0])

	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Body)
}
