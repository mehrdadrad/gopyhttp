package main

import "C"

import (
	"github.com/mehrdadrad/gopyhttp/gohttp"
)

//export HTTPRequest
func HTTPRequest(url string) {
	var r []gohttp.Request

	r = append(r, gohttp.Request{URL: url})

	gohttp.Run(r)
}

func main() {}
