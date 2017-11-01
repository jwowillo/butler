// Package main exposes a runnable file server that can have its listening
// directory and serving port configured. This is useful to serve the generated
// static files which make up the butler site.
package main

import (
	"flag"
	"log"
	"net/http"
)

// gzipHeader wraps the http.Handler to return gzip responses.
func gzipHeader(h http.Handler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Encoding", "gzip")
		h.ServeHTTP(w, r)
	}
}

// main binds a file server serving from a received directory, logs both the
// recieved port and directory, then listens for requests.
func main() {
	http.HandleFunc("/", gzipHeader(http.FileServer(http.Dir(dir))))
	log.Printf("listening on %s and serving files from %s\n", port, dir)
	http.ListenAndServe(port, nil)
}

var (
	// port to serve from.
	port string
	// dir to serve.
	dir string
)

// init parses command line arguments into received variables.
func init() {
	flag.StringVar(&port, "port", "", "port to serve from")
	flag.StringVar(&port, "p", "", "port to serve from")
	flag.StringVar(&dir, "directory", "", "directory with static files")
	flag.StringVar(&dir, "d", "", "directory with static files")
	flag.Parse()
	if port == "" {
		log.Fatal("must pass port to serve from")
	}
	if dir == "" {
		log.Fatal("must pass directory with static files")
	}
}
