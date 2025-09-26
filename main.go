package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http.HandleFunc: to handle the request
	// w : to write the response
	// r : to read the request
	// http.ResponseWriter : to write the response we don't use the pointer because it's interface and under the hood golang takes care of this and make it pointer
	// *http.Request : to read the request we use pointer to the request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w,: to send the response to the client
		// r.URL.Path: to get the path of the request
		// %s: to format the string
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	// Serves files from 'static/' directory; e.g. localhost/static/redis.png serves the image.
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// listen on 80 port, pass `nil` to say "Use the DefaultServeMux
	http.ListenAndServe(":80", nil)
}
