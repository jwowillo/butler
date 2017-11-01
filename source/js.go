package source

import (
	"bytes"
	"encoding/json"

	"github.com/jwowillo/butler/gen"
	"github.com/jwowillo/butler/recipe"
)

// sourceJS returns a gen.Source which has a list of recipe.Recipes as JS
// objects.
func sourceJS(rs []recipe.Recipe) gen.Source {
	return func() (*gen.Page, error) {
		b := bytes.NewBufferString("const recipes = [\n")
		for i, r := range rs {
			bs, err := json.MarshalIndent(r, "  ", "  ")
			if err != nil {
				return nil, err
			}
			b.WriteString("  ")
			b.Write(bs)
			if i < len(rs)-1 {
				b.WriteString(",\n")
			}
		}
		b.WriteString("\n];")
		return gen.NewPage("/recipes.js", b), nil
	}
}
