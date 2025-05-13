package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

const separator string = "========================================"
const mainDir string = "www-downloads"
const defaultURL string = "https://httpbin.org/"

// TODO: add logging that can be turned on and off adhoc by the user
// TODO: user should be able to run `Chrome /www-downloads/www.something.com`
// where `www.something.com` was one of URLs, and browse the page offline

// TODO: (MAYBE!) Add a config file that the user can use to override defaults...

// https://pkg.go.dev/os@go1.24.3#MkdirAll
func createDir(dir string) {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func getUrlList(args []string) []string {
	var u []string
	if len(args) > 1 {
		u = args[1:]
	} else {
		u = []string{defaultURL}
	}
	return u
}

// https://gobyexample.com/url-parsing
func parseURL(s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return u
}

// generic slice printing util
func printSlice[S ~[]E, E any](s S) {
	fmt.Println(separator)
	fmt.Printf("Slice: %v Type: %T\nLength = %d Capacity = %d\n", s, s, len(s), cap(s))
	for i, v := range s {
		fmt.Printf("%v : %v\n", i, v)
	}
	fmt.Println(separator)
}

var client *http.Client = &http.Client{}

// https://pkg.go.dev/net/http@go1.24.3#Get
func sendHttpRequests(s []*url.URL) {

	for i, v := range s {
		log.Println(i)

		res, err := client.Get(v.String())
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		} else {
			log.Print(res.StatusCode)
		}
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("%s", body)
		newDir := "www." + v.Hostname()
		newPath := mainDir + "/" + newDir
		createDir(newPath)
		err = os.WriteFile(newPath+"/index.html", body, 0644)
		if err != nil {
			log.Fatal(err)
		}

	}
}

// https://go.dev/tour/moretypes/15
func makeParsedUrlList(url_list []string) []*url.URL {
	var s []*url.URL
	for i := range url_list {
		t := parseURL(url_list[i])
		s = append(s, t)
	}
	return s
}

// TODO: HTTP Header
// TODO: https://developer.mozilla.org/en-US/docs/Glossary/Request_header
// TODO: https://stackoverflow.com/questions/46021330/how-can-i-read-a-header-from-an-http-request-in-golang
// TODO: https://siongui.github.io/2018/03/06/go-print-http-response-header/
// TODO: https://pkg.go.dev/net/http?utm_source=godoc#Header

func main() {

	createDir(mainDir)

	url_list := getUrlList(os.Args)
	url_list_parsed := makeParsedUrlList(url_list)
	printSlice(url_list_parsed)
	sendHttpRequests(url_list_parsed)

}
