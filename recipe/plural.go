package recipe

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

// l logs the errNoPlural for the word x.
func l(err error, x string) {
	log.Println(
		err.Error()+":",
		strings.ToLower(x),
		"needs to be added to dictionary",
	)
}

// SingularPhrase is the Ingredient's singular phrase if the amount is singular.
//
// This is the singular tense of the Ingredient ignoring the amount.
func SingularPhrase(i Ingredient) string {
	var out string
	if i.Unit == "" {
		out = strings.Title(i.Item)
	} else {
		p, err := plural(i.Item)
		if err != nil {
			l(err, i.Item)
		}
		out = fmt.Sprintf(
			"%s of %s",
			strings.Title(i.Unit), strings.Title(p),
		)
	}
	return out
}

// PluralPhrase is the Ingredient's plural phrase if the amount is a whole
// number.
//
// Will log if the plural couldn't be constructed.
func PluralPhrase(i Ingredient) string {
	var out string
	if i.Unit == "" {
		p, err := plural(i.Item)
		if err != nil {
			l(err, i.Item)
		}
		out = strings.Title(p)
	} else {
		pa, err := plural(i.Unit)
		if err != nil {
			l(err, i.Unit)
		}
		pb, err := plural(i.Item)
		if err != nil {
			l(err, i.Item)
		}
		out = fmt.Sprintf(
			"%s of %s",
			strings.Title(pa), strings.Title(pb),
		)
	}
	return out
}

// FractionalPhrase is the Ingredient's plural phrase if the amount isn't a
// whole number.
func FractionalPhrase(i Ingredient) string {
	var out string
	if i.Unit == "" {
		out = fmt.Sprintf(
			"of %s %s",
			article(i.Item), strings.Title(i.Item),
		)
	} else {
		p, err := plural(i.Item)
		if err != nil {
			l(err, i.Item)
		}
		out = fmt.Sprintf(
			"of %s %s of %s",
			article(i.Unit),
			strings.Title(i.Unit), strings.Title(p),
		)
	}
	return out
}

// errNoPlural is returned if a plural for the word isn't defined.
var errNoPlural = errors.New("no plural for word")

// plural for the word.
//
// Returns errNoPlural if no plural is defined.
func plural(x string) (string, error) {
	p, ok := plurals[strings.ToLower(x)]
	if !ok {
		return "", errNoPlural
	}
	return p, nil
}

// article to precede a word.
func article(x string) string {
	switch x[0] {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return "an"
	}
	return "a"
}

// plurals is a map of words to their defined plurals.
var plurals = map[string]string{
	"green bean":              "green beans",
	"salt":                    "salt",
	"pepper":                  "pepper",
	"cayenne":                 "cayenne",
	"bacon":                   "bacon",
	"cranberry":               "cranberries",
	"shallot":                 "shallots",
	"ham seasoning":           "ham seasoning",
	"cup":                     "cups",
	"slice":                   "slices",
	"pound":                   "pounds",
	"teaspoon":                "teaspoons",
	"sheet":                   "sheets",
	"clove":                   "cloves",
	"nutmeg":                  "nutmeg",
	"ground thyme":            "ground thyme",
	"ground sage":             "ground sage",
	"ground beef":             "ground beef",
	"pistachio":               "pistachios",
	"white onion":             "white onion",
	"filo dough":              "filo dough",
	"cucumber":                "cucumbers",
	"egg":                     "eggs",
	"garlic":                  "garlic",
	"mint":                    "mint",
	"lemon":                   "lemons",
	"potato":                  "potatoes",
	"butter":                  "butter",
	"white rice":              "white rice",
	"chicken broth":           "chicken broth",
	"tomato":                  "tomatoes",
	"corn kernel":             "corn kernels",
	"spanish onion":           "spanish onions",
	"poblano pepper":          "poblano peppers",
	"jalapeno":                "jalapenos",
	"red bell pepper":         "red bell peppers",
	"ribeye":                  "ribeyes",
	"neutral oil":             "neutral oil",
	"toast":                   "toast",
	"parsley":                 "parsley",
	"spicy sausage":           "spicy sausage",
	"shredded gruyere cheese": "shredded gruyere cheese",
	"olive oil":               "olive oil",
	"balsamic vinegar":        "balsamic vinegar",
	"dijon mustard":           "dijon mustard",
	"honey":                   "honey",
	"tablespoon":              "tablespoons",
	"thick slice":             "thick slices",
	"link":                    "links",
	"spinach":                 "spinach",
	"greek yogurt":            "greek yogurt",
	"dried cranberry":         "dried cranberries",
	"flour":                   "flour",
	"brown sugar":             "brown sugar",
	"cinnamon":                "cinnamon",
	"baking powder":           "baking powder",
	"milk":                    "milk",
	"stick":                   "sticks",
	"new york strip":          "new york strip",
}
