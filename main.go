package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// MVP Algorithm

// get URL,

// check with DNS (or local resolver) the IP address of the webserver,
// connect to the IP,
// and HTTP Headers and
// HTTP request,

// get HTTP Response,

// save to file

func main() {

	var URL string
	if len(os.Args) > 1 {
		URL = os.Args[1]
	} else {
		URL = "https://go.dev/"
	}

	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	e := os.WriteFile("./index.html", body, 0644)
	if e != nil {
		panic(e)
	}

}
