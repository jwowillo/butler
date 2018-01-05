(function() {

function recipeToString(recipe) {
  let out = recipe.name + recipe.description;
  for (const ingredient of recipe.ingredients) out += ingredient;
  for (const step of recipe.steps) out += step;
  return out;
}

function listRecipes(container, meal, filter) {
  set('filter', filter);
  filter = filter.toLowerCase();
  clear(container);
  const filtered = [];
  for (const recipe of recipes) {
    if (!recipeToString(recipe).toLowerCase().includes(filter)) continue;
    filtered.push(a(recipe.path, recipe.name));
  }
  container.appendChild(ul(filtered));
  addCheckBoxes(meal, container, getChecked());
}

const input = document.getElementById('filter');
const meal = document.getElementById('box');
const results = document.getElementById('results');

input.addEventListener('keyup', (event) => listRecipes(results, meal, input.value));

listRecipes(results, meal, get('filter'));

})();
