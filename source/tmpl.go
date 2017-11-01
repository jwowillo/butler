package source

import (
	"bytes"
	"html/template"
	"io"
	"path/filepath"

	"github.com/jwowillo/butler/gen"
	"github.com/jwowillo/butler/recipe"
)

// sourcesTmpl returns all the template gen.Sources. This is a combination of
// home page templates and recipe page templates.
//
// There is one recipe template for every recipe.Recipe. The gen.Sources which
// have JS file paths are placed into the home page template.
func sourcesTmpl(rs []recipe.Recipe, ss []gen.Source) []gen.Source {
	return append(sourcesRecipes(rs), sourceHome(rs, ss))
}

// sourceHome returns the home page gen.Source. The recipe.Recipes are made into
// links placed in the home page and all the gen.Sources with JS file paths are
// placed in the home page.
func sourceHome(rs []recipe.Recipe, ss []gen.Source) gen.Source {
	return func() (*gen.Page, error) {
		t, err := makeTmpl("index", rs, ss)
		if err != nil {
			return nil, err
		}
		return gen.NewPage("index.html", t), nil
	}
}

func sourcesRecipes(rs []recipe.Recipe) []gen.Source {
	ss := make([]gen.Source, 0, len(rs))
	for _, r := range rs {
		x := r
		ss = append(ss, func() (*gen.Page, error) {
			t, err := makeTmpl("recipe", x, nil)
			if err != nil {
				return nil, err
			}
			p := filepath.Join(x.Path, "index.html")
			return gen.NewPage(p, t), nil
		})
	}
	return ss
}

// makeTmpl renders a template into an io.Reader and a nil error or nil and a
// non nil error if the template couldn't be rendered. name is the path of the
// template to be rendered without the file extension. x is extra values meant
// to be injected into the template. ss are gen.Sources which represent JS files
// that are injected into the templates.
func makeTmpl(name string, x interface{}, ss []gen.Source) (io.Reader, error) {
	ss = filter(".js", ss)
	ps := make([]*gen.Page, 0, len(ss))
	for _, s := range ss {
		p, err := s()
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
	base := "tmpl"
	tmpl := template.New(name)
	bs, err := Asset(filepath.Join(base, "base.html"))
	if err != nil {
		return nil, err
	}
	tmpl, err = tmpl.Parse(string(bs))
	if err != nil {
		return nil, err
	}
	bs, err = Asset(filepath.Join(base, name+".html"))
	if err != nil {
		return nil, err
	}
	tmpl, err = tmpl.Parse(string(bs))
	if err != nil {
		return nil, err
	}
	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, struct {
		X     interface{}
		Pages []*gen.Page
	}{X: x, Pages: ps}); err != nil {
		return nil, err
	}
	return buf, nil
}
