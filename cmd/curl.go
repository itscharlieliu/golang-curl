package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"

	"github.com/itscharlieliu/golang-curl/pkg"
)

func main() {
	flag.Parse()

	args := flag.Args()

	fmt.Println(args[0])

	resp, err := http.Get(args[0])

	pkg.CheckErr(err)

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	pkg.CheckErr(err)

	fmt.Println(body)
}
