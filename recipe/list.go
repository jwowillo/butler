package recipe

import (
	"io/ioutil"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

// List of all Recipes stored in YAML files in the directory.
//
// The name of the YAML file becomes the path.
//
// The expected structure of each YAML file is:
//   name: <NAME>
//   description: <DESCRIPTION>
//   ingredients:
//       - ?amount: <AMOUNT>
//         ?unit: <UNIT>
//         ?item: <ITEM>
//       - ?amount: <AMOUNT>
//         ?unit: <UNIT>
//         ?item: <ITEM>
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

// recipesWalk returns a filepath.WalkFunc that looks for YAML files to convert
// to Recipes and store them in the list.
//
// The filepath.WalkFunc returns any error that occurs.
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
		r, err := yamlToRecipe(ry, path[:len(path)-len(ext)])
		if err != nil {
			return err
		}
		*rs = append(*rs, r)
		return nil
	}
}
