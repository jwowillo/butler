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

func (i Ingredient) String() string {
	if IsUndefined(i.Amount) {
		return SingularPhrase(i)
	} else if IsWhole(i.Amount) && i.Amount.Numerator == 1 {
		return fmt.Sprintf("1 %s", SingularPhrase(i))
	} else if IsWhole(i.Amount) {
		return fmt.Sprintf("%s %s", i.Amount, PluralPhrase(i))
	}
	return fmt.Sprintf("%s %s", i.Amount, FractionalPhrase(i))
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
	if IsUndefined(f) {
		return Ingredient{}, ErrUndefined
	}
	return Ingredient{
		Amount: f,
		Unit:   iy.Unit,
		Item:   iy.Item,
	}, nil
}
