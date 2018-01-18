(function() {

function recipes(checked) {
  return checked.map(item => RECIPES[item]);
}

// addCheckboxes that put meals in the mealContainer to the left of recipes in
// the recipeContainer.
function addCheckboxes(mealContainer, recipeContainer) {
  checkboxList(
    recipeContainer,
    checked => makeMeal(mealContainer, recipeContainer, recipes(checked)),
    checked => makeMeal(mealContainer, recipeContainer, recipes(checked))
  )
}

// makeMeal adds the checked meals to the mealContainer and inserts it before
// the recipeContainer.
function makeMeal(mealContainer, recipeContainer, checked) {
  clear(mealContainer);
  if (checked.length == 0) {
    remove(mealContainer);
    return;
  }
  mealContainer.appendChild(h3('Meal:'));
  mealContainer.appendChild(h2('Ingredients:'));
  const ingredientsUl = ul(ingredients(checked));
  ingredientsUl.id = 'ingredients';
  for (const item of checked) ingredientsUl.id += item.name;
  strikethroughList(ingredientsUl);
  mealContainer.appendChild(ingredientsUl);
  mealContainer.appendChild(h2('Steps:'));
  mealContainer.appendChild(ul(recipeLinks(checked)));
  prepend(recipeContainer, mealContainer);
}

// recipeLinks for checked recipes.
function recipeLinks(checked) {
  const ls = [];
  for (const recipe of checked) ls.push(a(recipe.path, recipe.name));
  return ls;
}

// ingredients for checked recipes.
function ingredients(checked) {
  const is = [];
  const used = new Set();
  for (const i in checked) {
    for (const ia of checked[i].ingredients) {
      if (used.has(ia.singularPhrase)) continue;
      used.add(ia.singularPhrase);
      for (const j in checked) {
        if (i == j) continue;
        for (const ib of checked[j].ingredients) {
          if (ia.singularPhrase == ib.singularPhrase &&
            !isUndefined(ia.amount) && !isUndefined(ib.amount)) {
            ia.amount = add(ia.amount, ib.amount);
          }
        }
      }
      if (isUndefined(ia.amount)) {
        is.push(ia.singularPhrase);
      } else if (isWhole(ia.amount) && ia.amount.numerator == 1) {
        is.push(ia.amount + ' ' + ia.singularPhrase);
      } else if (isWhole(ia.amount)) {
        is.push(ia.amount + ' ' + ia.pluralPhrase);
      } else {
        is.push(ia.amount + ' ' + ia.fractionalPhrase);
      }
    }
  }
  return is;
}

const mealContainer = document.createElement('div');
mealContainer.id = 'box';
const recipeContainer = document.getElementById('results');

addCheckboxes(mealContainer, recipeContainer);

})();
