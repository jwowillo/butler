// Package main exposes a command which builds the static files which make up
// the butler site into an dirput directory from an input directory containing
// recipe files and a URL the site will be served from.
package main

import (
	"flag"
	"log"

	"github.com/jwowillo/butler/page"
	"github.com/jwowillo/butler/recipe"
	"github.com/jwowillo/gen"
)

// main builds the butler static files from input files and a URL the site will
// be served from.
func main() {
	rs, err := recipe.List(recipes)
	if err != nil {
		log.Fatal(err)
	}
	ps, err := page.List(web, rs)
	if err != nil {
		log.Fatal(err)
	}
	if debug {
		err = gen.WriteOnly(dir, ps)
	} else {
		ts := []gen.Transform{gen.Minify, gen.Gzip}
		err = gen.Write(dir, ts, ps)
	}
	if err != nil {
		log.Fatal(err)
	}
}

var (
	// web directory.
	web string
	// recipes directory.
	recipes string
	// dir directory to write files to.
	dir string
	// debug is true if written files shouldn't be gzipped or minified.
	debug bool
)

func stringVar(s *string, f, h string) {
	flag.StringVar(s, f, "", h)
}

func boolVar(b *bool, f, h string) {
	flag.BoolVar(b, f, false, h)
}

// init parses command line aruments into received variables.
func init() {
	stringVar(&web, "web", "directory with web files")
	stringVar(&recipes, "recipes", "directory with recipe files")
	stringVar(&dir, "directory", "directory to build to")
	boolVar(&debug, "debug", "files won't be gzipped and minified if set")
	flag.Parse()
	if web == "" {
		log.Fatal("must pass directory with web files")
	}
	if recipes == "" {
		log.Fatal("must pass directory with recipe files")
	}
	if dir == "" {
		log.Fatal("must pass directory to build to")
	}
}
