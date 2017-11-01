// Package gen allows sources of files to be built into a directory lazily.
// These sources only require a path and an io.Reader to fetch the file content
// from.
package gen

// TODO: Keep source since it simplifies collation of stuff at source
// implementation level and lets them all be logged here, when they're
// instantiated. Also simplifies logic.

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"mime"
	"os"
	"path/filepath"
	"regexp"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
	"github.com/tdewolff/minify/xml"
)

// Page is an io.Reader containing the Page's content with a path describing
// where the Page exists at.
type Page struct {
	Path string
	io.Reader
}

// NewPage creates a Page out of a path the Page exists at and the io.Reader
// whcih contains the Page's content.
func NewPage(p string, r io.Reader) *Page {
	return &Page{Path: p, Reader: r}
}

// Source is a constructor which either returns a Page and no error or nil and
// an error if the Source couldn't construct the page.
type Source func() (*Page, error)

// Generate writes all the Sources to the out directory after minifying then
// gzipping them.
//
// Returns an error if the Sources couldn't be written, meaning there could have
// been errors minifying, gzipping, constructing, or writing.
func Generate(out string, ss []Source) error {
	t := compose(gzipTransform, minifyTransform)
	ss = apply(t, ss)
	return write(out, ss...)
}

// write all the Sources to the directory named out by constructing all of them
// and writing them as files relative to out.
//
// Returns an error if any Sources couldn't be constructed or written to files.
func write(out string, ss ...Source) error {
	for _, s := range ss {
		p, err := s()
		if err != nil {
			return err
		}
		full := filepath.Join(out, p.Path)
		if err := os.MkdirAll(
			filepath.Dir(full),
			os.ModePerm,
		); err != nil {
			return err
		}
		f, err := os.Create(full)
		if err != nil {
			return err
		}
		defer f.Close()
		if _, err := io.Copy(f, p); err != nil {
			return err
		}
	}
	return nil
}

// transform accepts one Source and returns a Source.
type transform func(Source) Source

// apply the transform to all of the Sources creating a list of Sources the
// same length as the original.
func apply(t transform, ss []Source) []Source {
	as := make([]Source, 0, len(ss))
	for _, s := range ss {
		as = append(as, t(s))
	}
	return as
}

// compose transforms t1 and t2 into a new transform that is the equivalent of
// applying t1 on the result of applying t2 to a Source.
func compose(t1, t2 transform) transform {
	return func(s Source) Source {
		return t1(t2(s))
	}
}

// minifyTransform transforms the Source into a minified version of the Source
// if it has an CSS, HTML, JavaScript, or XML file extension.
//
// The new Source will return an error if the initial Source couldn't be
// constructed or minified.
//
// The initial Source will be returned if it has an unsupported file extension.
func minifyTransform(s Source) Source {
	return func() (*Page, error) {
		p, err := s()
		if err != nil {
			return nil, err
		}
		m := minify.New()
		m.AddFunc("text/css", css.Minify)
		m.AddFunc("text/html", html.Minify)
		m.AddFunc("application/javascript", js.Minify)
		m.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)
		ext := filepath.Ext(p.Path)
		if ext != ".xml" &&
			ext != ".css" &&
			ext != ".js" &&
			ext != ".html" {
			return p, nil
		}
		t := mime.TypeByExtension(ext)
		buf := &bytes.Buffer{}
		if err := m.Minify(t, buf, p); err != nil {
			return nil, err
		}
		return NewPage(p.Path, buf), nil
	}
}

// gzipTransform transforms the Source into a gzipped version of the Source.
//
// Returns an error if the initial Source couldn't be isntantiated, if the
// constructed Source couldn't be read, or if the Source couldn't be gzipped.
func gzipTransform(s Source) Source {
	return func() (*Page, error) {
		p, err := s()
		if err != nil {
			return nil, err
		}
		bs, err := ioutil.ReadAll(p)
		if err != nil {
			return nil, err
		}
		buf := &bytes.Buffer{}
		zw := gzip.NewWriter(buf)
		if _, err := zw.Write(bs); err != nil {
			return nil, err
		}
		if err := zw.Close(); err != nil {
			return nil, err
		}
		return NewPage(p.Path, buf), nil
	}
}
