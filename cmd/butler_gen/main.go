// Package main has a command that generates butler static files based on
// content in input directories.
package main

import (
	"flag"
	"log"

	"github.com/jwowillo/butler/page"
	"github.com/jwowillo/butler/recipe"
	"github.com/jwowillo/gen"
)

// main builds the butler static files from input directories and gzips and
// minifies the files if the debug flag isn't set.
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
	// dir to write files to.
	dir string
	// debug is true if written files shouldn't be gzipped or minified.
	debug bool
)

// stringFlag that stores the value of the flag f in the string s with help
// message h.
func stringFlag(s *string, f, h string) {
	flag.StringVar(s, f, "", h)
}

//  boolFlag that stores the value of the flag f in the bool b with help message
//  h.
func boolFlag(b *bool, f, h string) {
	flag.BoolVar(b, f, false, h)
}

// init parses flags into variables.
func init() {
	stringFlag(&web, "web", "directory with web files")
	stringFlag(&recipes, "recipes", "directory with recipe files")
	stringFlag(&dir, "directory", "directory to build to")
	boolFlag(&debug, "debug", "files won't be gzipped and minified if set")
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
