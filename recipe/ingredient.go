package recipe

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	// ErrUnitWithNoAmount is returned if an Ingredient has a unit and no
	// amount.
	ErrUnitWithNoAmount = errors.New("can't have unit with no amount")
	// ErrUndefined is returned if an Ingredient has an undefined amount.
	ErrUndefined = errors.New("amount can't be undefined")
)

// Ingredient is a component in a Recipe.
type Ingredient struct {
	// Amount of the Ingredient.
	Amount Fraction
	// Unit the amount is in.
	Unit string
	// Item the Ingredient actually is.
	Item string
}

// String representation of the Ingredient.
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

// ingredientYAML describes the structure of a Ingredient's YAML.
type ingredientYAML struct {
	Amount string `yaml:"amount"`
	Unit   string `yaml:"unit"`
	Item   string `yaml:"item"`
}

// yamlToIngredient converts the ingredientYAML to an Ingredient.
//
// Returns ErrUnitWithNoAmount and ErrUndefined if either are met.
func yamlToIngredient(iy ingredientYAML) (Ingredient, error) {
	if iy.Amount == "" {
		if iy.Unit != "" {
			return Ingredient{}, ErrUnitWithNoAmount
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

// article x should be preceeded by.
func article(x string) string {
	switch x[0] {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return "an"
	}
	return "a"
}

// plural form of x.
func plural(x string) string {
	if x[len(x)-1] == 'o' {
		return x + "es"
	}
	if x[len(x)-1] == 'y' {
		return x[:len(x)-1] + "ies"
	}
	return x + "s"
}
