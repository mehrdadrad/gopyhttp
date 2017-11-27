package main

import "C"

import (
	"github.com/mehrdadrad/pygohttp/gohttp"
)

//export HTTPRequest
func HTTPRequest(urls []string) {
	var r []gohttp.Request

	for _, url := range urls {
		r = append(r, gohttp.Request{URL: url})
	}

	gohttp.Run(r)
}

func main() {}
