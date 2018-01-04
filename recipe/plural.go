package recipe

import (
	"errors"
	"fmt"
	"strings"

	l "log"
)

func log(err error, x string) {
	l.Println(
		err,
		":",
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
			log(err, i.Item)
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
			log(err, i.Item)
		}
		out = strings.Title(p)
	} else {
		pa, err := plural(i.Unit)
		if err != nil {
			log(err, i.Item)
		}
		pb, err := plural(i.Item)
		if err != nil {
			log(err, i.Item)
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
			log(err, i.Item)
		}
		out = fmt.Sprintf(
			"of %s %s of %s",
			article(i.Unit),
			strings.Title(i.Unit), strings.Title(p),
		)
	}
	return out
}

var errNoPlural = errors.New("no plural for word")

func plural(x string) (string, error) {
	p, ok := plurals[strings.ToLower(x)]
	if !ok {
		return "", errNoPlural
	}
	return p, nil
}

func article(x string) string {
	switch x[0] {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return "an"
	}
	return "a"
}

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
}
