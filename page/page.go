package page

import (
	"path/filepath"

	"github.com/jwowillo/butler/recipe"
	"github.com/jwowillo/gen"
)

// List ...
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
