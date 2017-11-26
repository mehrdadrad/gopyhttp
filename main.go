package main

import (
	"github.com/mehrdadrad/gopyhttp/gohttp"
)

func main() {
	var r []gohttp.Request
	r = append(r, gohttp.Request{URL: "https://www.yahoo.com"})
	r = append(r, gohttp.Request{URL: "https://www.linux.com"})
	r = append(r, gohttp.Request{URL: "https://www.microsoft.com"})
	r = append(r, gohttp.Request{URL: "https://www.google.com"})
	r = append(r, gohttp.Request{URL: "https://www.facebook.com"})

	gohttp.Run(r)
}
