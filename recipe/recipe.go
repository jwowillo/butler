// Package recipe defines a Recipe and has a function to get all Recipes from a
// directory.
package recipe

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

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
	Ingredients []Ingredient
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

// TODO: Have ingredientYAML.

// recipeYAML describes the structure of a Recipe's YAML file.
type recipeYAML struct {
	Name        string           `yaml:"name"`
	Description string           `yaml:"description"`
	Notes       []string         `yaml:"notes"`
	Ingredients []ingredientYAML `yaml:"ingredients"`
	Steps       []string         `yaml:"steps"`
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
		is := make([]Ingredient, 0, len(ry.Ingredients))
		for _, iy := range ry.Ingredients {
			i, err := ingredientYAMLToIngredient(iy)
			if err != nil {
				return err
			}
			is = append(is, i)
		}
		path = path[0 : len(path)-len(ext)]
		*rs = append(*rs, Recipe{
			Path:        "/" + path + "/",
			Name:        ry.Name,
			Description: ry.Description,
			Notes:       ry.Notes,
			Steps:       ry.Steps,
			Ingredients: is,
		})
		return nil
	}
}

var (
	// ErrNoUnit ...
	ErrNoUnit = errors.New("must have unit if amount is present")
	// ErrUndefined ...
	ErrUndefined = errors.New("amount can't be undefined")
)

func ingredientYAMLToIngredient(iy ingredientYAML) (Ingredient, error) {
	if iy.Amount == "" {
		if iy.Unit != "" {
			return Ingredient{}, ErrNoUnit
		}
		return Ingredient{Item: iy.Item}, nil
	}
	var f Fraction
	if strings.Contains(iy.Amount, "/") {
		parts := strings.Split(iy.Amount, "/")
		num, err := strconv.Atoi(parts[0])
		if err != nil {
			return Ingredient{}, err
		}
		den, err := strconv.Atoi(parts[1])
		if err != nil {
			return Ingredient{}, err
		}
		f.Numerator, f.Denominator = num, den
	} else {
		num, err := strconv.Atoi(iy.Amount)
		if err != nil {
			return Ingredient{}, err
		}
		f.Numerator, f.Denominator = num, 1
	}
	if f.IsUndefined() {
		return Ingredient{}, ErrUndefined
	}
	return Ingredient{
		Amount: f,
		Unit:   iy.Unit,
		Item:   iy.Item,
	}, nil
}

type ingredientYAML struct {
	Amount string `yaml:"amount"`
	Unit   string `yaml:"unit"`
	Item   string `yaml:"item"`
}

// Ingredient ...
type Ingredient struct {
	Amount     Fraction
	Unit, Item string
}

func article(x string) string {
	switch x[0] {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return "an"
	}
	return "a"
}

func plural(x string) string {
	if x[len(x)-1] == 'o' {
		return x + "es"
	}
	if x[len(x)-1] == 'y' {
		return x[:len(x)-1] + "ies"
	}
	return x + "s"
}

// String ...
func (i Ingredient) String() string {
	if i.Amount.IsUndefined() && i.Unit == "" {
		return i.Item
	}
	var out string
	if i.Unit != "" {
		if i.Amount.IsWhole() && i.Amount.Numerator == 1 {
			out = fmt.Sprintf(
				"%s %s of %s",
				i.Amount, i.Unit, i.Item,
			)
		} else if i.Amount.IsWhole() {
			out = fmt.Sprintf(
				"%s %s of %s",
				i.Amount, plural(i.Unit), i.Item,
			)
		} else {
			out = fmt.Sprintf(
				"%s of %s %s of %s",
				i.Amount,
				article(i.Unit),
				i.Unit,
				i.Item,
			)
		}
	} else {
		if i.Amount.IsWhole() && i.Amount.Numerator == 1 {
			out = fmt.Sprintf("%s %s", i.Amount, i.Item)
		} else if i.Amount.IsWhole() {
			out = fmt.Sprintf("%s %s", i.Amount, i.Item)
		} else {
			out = fmt.Sprintf(
				"%s of %s %s",
				i.Amount, article(i.Item), i.Item,
			)
		}
	}
	return out
}

// Fraction ...
type Fraction struct {
	Numerator, Denominator int
}

// IsWhole ...
func (f Fraction) IsWhole() bool {
	return f.Simplified().Denominator == 1
}

// IsUndefined ...
func (f Fraction) IsUndefined() bool {
	return f.Denominator == 0
}

// Simplified ...
func (f Fraction) Simplified() Fraction {
	a, b := f.Numerator, f.Denominator
	for b != 0 {
		a, b = b, a%b
	}
	return Fraction{
		Numerator:   f.Numerator / a,
		Denominator: f.Denominator / a,
	}
}

// String ...
func (f Fraction) String() string {
	f = f.Simplified()
	if f.Denominator == 1 {
		return fmt.Sprintf("%d", f.Numerator)
	}
	return fmt.Sprintf("%d/%d", f.Numerator, f.Denominator)
}
