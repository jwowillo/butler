package recipe_test

import (
	"fmt"
	"testing"

	"github.com/jwowillo/butler/recipe"
)

func TestIngredientString(t *testing.T) {
	i := recipe.Ingredient{Item: "cheese"}
	fmt.Println(i)
	i = recipe.Ingredient{
		Amount: recipe.Fraction{Numerator: 3, Denominator: 4},
		Unit:   "cup",
		Item:   "cheese",
	}
	fmt.Println(i)
	i = recipe.Ingredient{
		Amount: recipe.Fraction{Numerator: 10, Denominator: 2},
		Unit:   "cup",
		Item:   "cheese",
	}
	fmt.Println(i)
	i = recipe.Ingredient{
		Amount: recipe.Fraction{Numerator: 10, Denominator: 4},
		Unit:   "cup",
		Item:   "cheese",
	}
	fmt.Println(i)
	i = recipe.Ingredient{
		Amount: recipe.Fraction{Numerator: 1, Denominator: 1},
		Item:   "envelope",
	}
	fmt.Println(i)
	i = recipe.Ingredient{
		Amount: recipe.Fraction{Numerator: 1, Denominator: 2},
		Item:   "envelope",
	}
	fmt.Println(i)
}
