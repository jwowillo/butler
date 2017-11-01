// Package main exposes a command which builds the static files which make up
// the butler site into an output directory from an input directory containing
// recipe files and a URL the site will be served from.
package main

import (
	"flag"
	"log"

	"github.com/jwowillo/butler/gen"
	"github.com/jwowillo/butler/recipe"
	"github.com/jwowillo/butler/source"
)

// main builds the butler static files from input files and a URL the site will
// be served from.
func main() {
	rs, err := recipe.ListFromDir(in)
	if err != nil {
		log.Fatal(err)
	}
	if err := gen.Generate(out, source.All(url, rs)); err != nil {
		log.Fatal(err)
	}
}

var (
	// url to serve from.
	url string
	// in directory containing recipe files.
	in string
	// out directory to place built static files into.
	out string
)

// init parses command line aruments into received variables.
func init() {
	flag.StringVar(&url, "url", "", "URL which will be served from")
	flag.StringVar(&url, "u", "", "URL which will be served from")
	flag.StringVar(&in, "in", "", "directory with recipe files")
	flag.StringVar(&in, "i", "", "directory with recipe files")
	flag.StringVar(&out, "out", "", "directory to build into")
	flag.StringVar(&out, "o", "", "directory to build into")
	flag.Parse()
	if url == "" {
		log.Fatal("must pass URL which will be served from")
	}
	if in == "" {
		log.Fatal("must pass directory with recipe files")
	}
	if out == "" {
		log.Fatal("must pass directory to build into")
	}
}
