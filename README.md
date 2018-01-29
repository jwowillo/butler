# `butler`

`butler` is a small and fast website which lists searchable recipes served from
http://www.recipe-butler.com.

## Installing

Run `make` to make docs and all commands. Run `make butler_gen|doc` to only make
the corresponding command.

## Deploying

`update_dependencies` and `deploy` are two scripts are included to aid in
deployment to a remote host. `update_dependencies` installs dependencies on the
host and restarts the server while `deploy` clones the project to the host,
generates the site, and starts the server. Instructions for running each can be
found by running `update_dependencies|deploy --help` for the respective command.

## Running

`butler_gen` is the main command which generates the `butler` website.
Instructions for running it can be found after installing the command and
running `butler_gen --help`.

`run_gen|server` are wrappers with default arguments around `butler_gen` and
`gen_server` respectively.

## Documentation

Documentation is in directory 'doc' as Markdown files for the project
requirements and design. PDFs generated from the Markdown files are also
included.

API documentation is located at
https://www.godoc.org/github.com/jwowillo/butler.
