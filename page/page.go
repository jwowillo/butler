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

// recipeTemplate is the struct for a recipe.Recipe which is injected into
// templates.
type recipeTemplate struct {
	Path        string
	Name        string
	Description string
	Ingredients []ingredientTemplate
	Steps       []string
	Notes       []string
}

// newRecipeTemplate makes a recipeTemplate from a recipe.Recipe.
func newRecipeTemplate(r recipe.Recipe) recipeTemplate {
	is := make([]ingredientTemplate, 0, len(r.Ingredients))
	for _, i := range r.Ingredients {
		is = append(is, ingredientTemplate{
			Ingredient:       i,
			SingularPhrase:   recipe.SingularPhrase(i),
			PluralPhrase:     recipe.PluralPhrase(i),
			FractionalPhrase: recipe.FractionalPhrase(i),
		})
	}
	return recipeTemplate{
		Path:        r.Path,
		Name:        r.Name,
		Description: r.Description,
		Ingredients: is,
		Steps:       r.Steps,
		Notes:       r.Notes,
	}
}

// ingredientTemplate is the struct for a recipe.Ingredient which is injected
// into templates.
type ingredientTemplate struct {
	recipe.Ingredient
	SingularPhrase   string
	PluralPhrase     string
	FractionalPhrase string
}

// tmpls is a helper function to return a list of all the gen.Templates for
// butler or an error if the gen.Templates couldn't be created.
func tmpls(web string, rs []recipe.Recipe) ([]gen.Page, error) {
	rts := make([]recipeTemplate, 0, len(rs))
	for _, r := range rs {
		rts = append(rts, newRecipeTemplate(r))
	}
	paths := func(ps ...string) []string {
		out := make([]string, 0, len(ps))
		for _, p := range ps {
			out = append(out, filepath.Join(web, "tmpl", p))
		}
		return out
	}
	hp, err := gen.NewTemplate(
		paths("base.html", "index.html"),
		rts,
		"/index.html",
	)
	if err != nil {
		return nil, err
	}
	jsp, err := gen.NewTemplate(paths("recipes.js"), rts, "/recipes.js")
	if err != nil {
		return nil, err
	}
	rps := make([]gen.Page, 0, len(rts))
	for _, rt := range rts {
		rp, err := gen.NewTemplate(
			paths("base.html", "recipe.html"),
			rt,
			filepath.Join(rt.Path, "index.html"),
		)
		if err != nil {
			return nil, err
		}
		rps = append(rps, rp)
	}
	return append(rps, hp, jsp), nil
}
