// Package recipe defines a Recipe and has a function to get all Recipes from a
// directory.
package recipe

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Recipe has all the data which describes a recipe.
type Recipe struct {
	// Path the Recipe will be located at.
	Path string
	// Name of the Recipe.
	Name string
	// Description sentence of the Recipe.
	Description string
	// Ingredients to make the Recipe.
	Ingredients []string
	// Steps to make the Recipe.
	Steps []string
	// Notes which will be helpful when making the Recipe.
	Notes []string
}

// List of all Recipes stored in YAML files in the directory.
//
// The name of the YAML file becomes the path.
//
// The expected structure of each YAML file is:
//   name: <NAME>
//   description: <DESCRIPTION>
//   ingredients:
//       - <INGREDIENT>
//       - <INGREDIENT>
//   steps:
//       - <STEP>
//       - <STEP>
//   notes:
//       - <NOTE>
//       - <NOTE>
//
// Returns an error if any if the YAML files couldn't be read into Recipes.
func List(dir string) ([]Recipe, error) {
	var rs []Recipe
	if err := filepath.Walk(dir, recipesWalk(&rs)); err != nil {
		return nil, err
	}
	return rs, nil
}

// recipeYAML describes the structure of a Recipe's YAML file.
type recipeYAML struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Notes       []string `yaml:"notes"`
	Ingredients []string `yaml:"ingredients"`
	Steps       []string `yaml:"steps"`
}

// recipesWalk returns a filepath.WalkFunc that looks for YAML files to convert
// to Recipes and store them in the list.
//
// The filepath.WalFunc returns any error that occurs.
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
