// Package source exposes a function called All which returns all gen.Sources
// meant to be passed to gen.Build to generate the butler gensite.
package source

import (
	"path/filepath"

	"github.com/jwowillo/butler/gen"
	"github.com/jwowillo/butler/recipe"
)

//go:generate go-bindata --pkg $GOPACKAGE --prefix data data/asset data/tmpl

// All gen.Sources for the butler site at the URL with all of the
// recipe.Recipes.
func All(url string, rs []recipe.Recipe) []gen.Source {
	ss := sourcesAsset()
	ss = append(ss, sourceJS(rs))
	ss = append(ss, sourcesTmpl(rs, ss)...)
	return append(ss, sourceSitemap(url, ss))
}

// filter gen.Sources based on their file path extensions.
func filter(ext string, ss []gen.Source) []gen.Source {
	var fss []gen.Source
	for _, s := range ss {
		p, err := s()
		if err != nil {
			continue
		}
		if filepath.Ext(p.Path) == ext {
			fss = append(fss, s)
		}
	}
	return fss
}
