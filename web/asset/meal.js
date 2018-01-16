(function() {

// getChecked returns checked recipes.
function getChecked() {
  return getRecipes('checked');
}

// setChecked sets checked recipes.
function setChecked(checked) {
  setRecipes('checked', checked);
}

// addCheckBoxes that put meals in the mealContainer to the left of recipes in
// the recipeContainer.
function addCheckBoxes(mealContainer, recipeContainer) {
  const checked = new Set();
  for (const recipe of getChecked()) checked.add(recipe.name);
  for (const item of recipeContainer.getElementsByTagName('li')) {
    const link = item.firstChild;
    const input = checkBox(
      function() {
        const checked = getChecked();
        checked.push(recipes[link.innerHTML]);
        setChecked(checked);
        makeMeal(mealContainer, recipeContainer, checked);
      },
      function() {
        let checked = getChecked();
        checked = checked.filter((recipe) => recipe.name != link.innerHTML);
        setChecked(checked);
        makeMeal(mealContainer, recipeContainer, checked);
      }
    )
    if (checked.has(link.innerHTML)) input.checked = true;
    prepend(link, input);
  }
  makeMeal(mealContainer, recipeContainer, getChecked());
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
  mealContainer.appendChild(ul(ingredients(checked)));
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

const meal = document.createElement('div');
meal.id = 'box';
const results = document.getElementById('results');

addCheckBoxes(meal, results);

})();
