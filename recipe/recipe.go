// Package recipe provides a structure for what a Recipe is in JSON and YAML
// formats and a utility called AllInDir for fetching all the Recipes in a
// directory from YAML files.
package recipe

// TODO: Don't reimplement these errors, use handling in OS. Make sure that the
// right things are still returned in test though. If a subdir or nonYAML file
// has no permission error, it should be skipped without bubbling errors to the
// top. Document that the errors are those defined plus anything that can show
// up during the walk.

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Recipe ...
type Recipe struct {
	Path        string
	Name        string
	Description string
	Ingredients []string
	Steps       []string
	Notes       []string
}

// List returns all Recipes stored as YAML files in the provided
// directory. An error is returned if the Recipes in the directory can't be read.
func List(dir string) ([]Recipe, error) {
	var rs []Recipe
	if err := filepath.Walk(dir, recipesWalk(&rs)); err != nil {
		return nil, err
	}
	return rs, nil
}

type recipeYAML struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Notes       []string `yaml:"notes"`
	Ingredients []string `yaml:"ingredients"`
	Steps       []string `yaml:"steps"`
}

// recipesWalk returns a filepath.WalkFunc that stores all found Recipes in the
// provided list. The filepath.WalkFunc returns errors if Recipes can't be read.
func recipesWalk(rs *[]Recipe) filepath.WalkFunc {
	return func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		ext := filepath.Ext(path)
		if ext != ".yaml" {
			return nil
		}
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}
		ry := recipeYAML{}
		if err := yaml.Unmarshal(bs, &ry); err != nil {
			return err
		}
		path = path[0 : len(path)-len(ext)]
		*rs = append(*rs, Recipe{
			Path:        "/" + path + "/",
			Name:        ry.Name,
			Description: ry.Description,
			Notes:       ry.Notes,
			Steps:       ry.Steps,
			Ingredients: ry.Ingredients,
		})
		return nil
	}
}
