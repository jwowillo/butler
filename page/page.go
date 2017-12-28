// Package page has a function to return a list of all butler gen.Pages.
package page

import (
	"path/filepath"

	"github.com/jwowillo/butler/recipe"
	"github.com/jwowillo/gen"
)

// List of all butler gen.Pages with static files in the web directory and
// recipe.Recipes rs to be injected into the static files.
//
// The passed web directory must have a directory called 'tmpl' that contains
// template files 'base.html', 'index.html', 'recipes.js', and 'recipe.html'.
//
// Returns an error if any of the gen.Pages couldn't be created.
func List(web string, rs []recipe.Recipe) ([]gen.Page, error) {
	ps, err := tmpls(web, rs)
	if err != nil {
		return nil, err
	}
	as, err := gen.Assets(filepath.Join(web, "asset"))
	if err != nil {
		return nil, err
	}
	return append(ps, as...), nil
}

// tmpls is a helper function to return a list of all the gen.Templates for
// butler or an error if the gen.Templates couldn't be created.
func tmpls(web string, rs []recipe.Recipe) ([]gen.Page, error) {
	paths := func(ps ...string) []string {
		out := make([]string, 0, len(ps))
		for _, p := range ps {
			out = append(out, filepath.Join(web, "tmpl", p))
		}
		return out
	}
	hp, err := gen.NewTemplate(
		paths("base.html", "index.html"),
		rs,
		"index.html",
	)
	if err != nil {
		return nil, err
	}
	jsp, err := gen.NewTemplate(paths("recipes.js"), rs, "recipes.js")
	if err != nil {
		return nil, err
	}
	rps := make([]gen.Page, 0, len(rs))
	for _, r := range rs {
		rp, err := gen.NewTemplate(
			paths("base.html", "recipe.html"),
			r,
			filepath.Join(r.Path, "index.html"),
		)
		if err != nil {
			return nil, err
		}
		rps = append(rps, rp)
	}
	return append(rps, hp, jsp), nil
}
