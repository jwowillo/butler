# butler

butler is a small and fast website which lists searchable recipes.

## Recipes

The default recipes the website is generated from are kept in the `recipe`
package. They are YAML files containing the content of the recipe. The command
`butler_gen` does allow for other directories containing recipes to be passed as
flags as long as the recipes they contain are properly formatted.

## Documentation

Documentation is kept in the 'doc' directory with 'requirements.pdf' and
'design.pdf'. These are generated from corresponding Markdown files. API
documentation is also hosted at godoc.org/github.com/jwowillo/butler.

## Tests

The `recipe`, `source`, and `gen` packages contain tests named with '_test.go'
suffixes. Each can be ran with `go test` in the corresponding package.

## Commands

Commands are located in the `cmd` package if they are library commands or in the
project root if they are commands useful for debugging. Library commands are
named `butler_<X>` where variable `X` is a unique name for the command. The
commands are:

* `butler_gen`: Generates the butler static files. Accepts 'url', 'in', and
  'out' flags. 'url' is the URL the generated site will be hosted on, 'in' is
  the directory containing the recipe files, 'out' is the directory to place the
  generated files into.
* `butler_server`: Serves the butler static files. Accepts 'url' and 'directory'
  flags. 'url' is the URL the site will be hosted on and 'directory' is the
  directory the butler static files are kept in.
* `rebuild`: Rebuilds the commands and the static files. Useful for debugging.
* `run_debug`: Runs the server with default flags on a local port. Useful for
  debugging.

## Building

A Makefile contains all build targets. These include:

* `generate`: Runs `go generate` in the source directory to generate Go files
  with static files embedded in them.
* `butler_server`: Builds the `butler_server` command.
* `butler_gen`: Builds the `butler_gen` command.
* `doc`: Generates the PDF documentation from the Markdown files.

