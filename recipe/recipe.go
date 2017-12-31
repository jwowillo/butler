// Package recipe defines a Recipe and has a function to get a list of all
// Recipes in a directory.
package recipe

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

// recipeYAML describes the structure of a Recipe's YAML.
type recipeYAML struct {
	Name        string           `yaml:"name"`
	Description string           `yaml:"description"`
	Notes       []string         `yaml:"notes"`
	Ingredients []ingredientYAML `yaml:"ingredients"`
	Steps       []string         `yaml:"steps"`
}

// recipeToIngredient converts the recipeYAML to a Recipe.
//
// Returns an error if the recipeYAML couldn't be converted to a Recipe.
func yamlToRecipe(ry recipeYAML, path string) (Recipe, error) {
	is := make([]Ingredient, 0, len(ry.Ingredients))
	for _, iy := range ry.Ingredients {
		i, err := yamlToIngredient(iy)
		if err != nil {
			return Recipe{}, err
		}
		is = append(is, i)
	}
	return Recipe{
		Path:        "/" + path + "/",
		Name:        ry.Name,
		Description: ry.Description,
		Notes:       ry.Notes,
		Steps:       ry.Steps,
		Ingredients: is,
	}, nil
}
