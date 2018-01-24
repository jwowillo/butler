# Makefile contains targets to assist in building butler utilities. Ones that
# don't build commands log what they are doing and a new line at the end. Ones
# that do build commands echo 'making', then what they're making, and a new
# line at the end.
.PHONY: doc

# all builds the site generator, the server, and docs.
all: butler_gen doc

# butler_gen builds the generator which creates the butler static files.
butler_gen:
	@echo 'making butler_gen'
	$(call go,butler_gen)
	@echo

# doc builds the project documentation.
doc:
	@echo 'making doc'
	$(call pdf,requirements)
	$(call pdf,design)
	@echo


# pdf is used to make PDFs from Latex files using Pandoc. The Latex files are
# expected to be found in the doc directory.
define pdf
	pandoc doc/$(1).md --latex-engine xelatex -o doc/$(1).pdf
endef

# go is used to install Go commands referred to by the name of the command. The
# commands are expected to be found in the cmd directory.
define go
	cd cmd/$(1) && go get && go install
endef
