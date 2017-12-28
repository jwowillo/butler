# `butler`

`butler` is a small and fast website which lists searchable recipes served from
http://www.recipe-butler.com.

## Installing

Run `make` to make docs and all commands. Run `make doc` to only make
documentation. Run `make butler_gen|server` to make the corresponding command.

## Running

Instructions for running `butler_gen|server` can be found after installing the
commands by running `butler_gen|server --help`.

`run` does both `butler_gen` and `butler_server` and accepts the union of the
flags from both.
`deploy` installs and starts the server on a host and is run with the host to
deploy to in the '--host' flag and the working directory to run the server from
in the '--working-directory' flag.

## Documentation

Documentation is located in directory 'doc' in Markdown files for the project
requirements and design. PDFs generated from these are also included.

Online API documentation is located at godoc.org/github.com/jwowillo/butler.
