// Package recipe provides a structure for what a Recipe is in JSON and YAML
// formats and a utility called AllInDir for fetching all the Recipes in a
// directory from YAML files.
package recipe

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type recipeYAML struct {
	Name        string        `yaml:"name"`
	Description string        `yaml:"description"`
	Notes       []string      `yaml:"notes"`
	Ingredients []interface{} `yaml:"ingredients"`
	Steps       []interface{} `yaml:"steps"`
}

// Recipe ...
type Recipe struct {
	Path        string
	Name        string
	Description string
	Notes       []string
	Ingredients NestableList
	Steps       NestableList
}

// ListFromDir returns all Recipes stored as YAMl files in the provided
// directory. An error is returned if the Recipes in the directory can't be read.
func ListFromDir(dir string) ([]Recipe, error) {
	var rs []Recipe
	if err := filepath.Walk(dir, recipesWalk(&rs)); err != nil {
		return nil, err
	}
	return rs, nil
}

// recipesWalk returns a filepath.WalkFunc that stores all found Recipes in the
// provided list. The filepath.WalkFunc returns errors if Recipes can't be read.
func recipesWalk(rs *[]Recipe) filepath.WalkFunc {
	return func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fi.IsDir() {
			return nil
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
		steps, err := fillList(ry.Steps)
		if err != nil {
			return err
		}
		ingredients, err := fillList(ry.Ingredients)
		if err != nil {
			return err
		}
		*rs = append(*rs, Recipe{
			Path:        "/" + path + "/",
			Name:        ry.Name,
			Description: ry.Description,
			Notes:       ry.Notes,
			Steps:       steps,
			Ingredients: ingredients,
		})
		return nil
	}
}

func fillList(tvs []interface{}) (NestableList, error) {
	var nl NestableList
	for _, tv := range tvs {
		switch tv := tv.(type) {
		case map[interface{}]interface{}:
			for k, vs := range tv {
				k, ok := k.(string)
				if !ok {
					return nil, errors.New("bad type")
				}
				vs, ok := vs.([]interface{})
				if !ok {
					return nil, errors.New("bad type")
				}
				i := NestableItem{Item: k}
				for _, v := range vs {
					v, ok := v.(string)
					if !ok {
						return nil, errors.New("bad type")
					}
					i.List = append(i.List, NestableItem{Item: v})
				}
				nl = append(nl, i)
			}
		case string:
			nl = append(nl, NestableItem{Item: tv})
		}
	}
	return nl, nil
}
