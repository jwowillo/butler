(function() {

// BUG: Bookmarks aren't implemented yet.

function addCheckBoxes(mealContainer, recipeContainer) {
  const map = {};
  for (const recipe of recipes) map[recipe.name] = recipe;
  let checked = [];
  for (const item of recipeContainer.getElementsByTagName('li')) {
    const link = item.firstChild;
    const input = checkBox(
      function() {
        checked.push(map[link.innerHTML]);
        makeMeal(mealContainer, recipeContainer, checked);
      },
      function() {
        checked = checked.filter((recipe) => recipe.name != link.innerHTML);
        makeMeal(mealContainer, recipeContainer, checked);
      }
    )
    prepend(link, input);
  }
}

function makeMeal(mealContainer, recipeContainer, checked) {
  if (checked.length == 0) remove(mealContainer);
  if (checked.length == 1) prepend(recipeContainer, mealContainer);
  clear(mealContainer);
  mealContainer.appendChild(h3('Meal:'));
  mealContainer.appendChild(h2('Ingredients:'));
  mealContainer.appendChild(ul(ingredients(checked)));
  mealContainer.appendChild(h2('Steps:'));
  mealContainer.appendChild(ul(recipeLinks(checked)));
}

function recipeLinks(checked) {
  const ls = [];
  for (const recipe of checked) {
    const link = document.createElement('a');
    link.href = recipe.path;
    link.innerHTML = recipe.name;
    ls.push(link);
  }
  return ls;
}

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
