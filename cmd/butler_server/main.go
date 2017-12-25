// Package main exposes a runnable file server that can have its listening
// directory and serving port configured. This is useful to serve the generated
// static files which make up the butler site.
package main

// TODO: Update style

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
	if debug {
		http.HandleFunc("/", http.FileServer(http.Dir(dir)).ServeHTTP)
		log.Printf("listening on %s and serving files from %s\n", port, dir)
	} else {
		http.HandleFunc("/", gzipHeader(http.FileServer(http.Dir(dir))))
		log.Printf("listening on %s and serving gzipped files from %s\n", port, dir)
	}
	http.ListenAndServe(port, nil)
}

var (
	// port to serve from.
	port string
	// dir to serve.
	dir string
	// debug is true if directory to serve is gzipped.
	debug bool
)

func boolVar(b *bool, f, h string) {
	flag.BoolVar(b, f, false, h)
}

// init parses command line arguments into received variables.
func init() {
	flag.StringVar(&port, "port", "", "port to serve from")
	flag.StringVar(&dir, "directory", "", "directory with static files")
	boolVar(&debug, "debug", "serves non gzipped files")
	flag.Parse()
	if port == "" {
		log.Fatal("must pass port to serve from")
	}
	if dir == "" {
		log.Fatal("must pass directory with static files")
	}
}
