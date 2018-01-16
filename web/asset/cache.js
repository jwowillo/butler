// get object at key in local storage.
//
// Classes will need to be reconstructed since they are demoted to JSON objects.
function get(key) {
  return JSON.parse(localStorage.getItem(window.location.href+key));
}

// set object at key to value in local storage.
function set(key, value) {
  localStorage.setItem(window.location.href+key, JSON.stringify(value));
}

function getRecipes(key) {
  const recipes = get(key);
  if (!(recipes instanceof Array)) return [];
  for (const recipe of recipes) {
    for (const i in recipe.ingredients) {
      recipe.ingredients[i].amount = new Fraction(
        recipe.ingredients[i].amount.numerator,
        recipe.ingredients[i].amount.denominator
      );
    }
  }
  return recipes;
}

function setRecipes(key, value) {
  set(key, value);
}
