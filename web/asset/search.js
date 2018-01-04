(function() {

// BUG: Bookmarks not implemented.

function recipeToString(recipe) {
  let out = recipe.name + recipe.description;
  for (const ingredient of recipe.ingredients) out += ingredient;
  for (const step of recipe.steps) out += step;
  return out;
}

function listRecipes(filter, container) {
  filter = filter.toLowerCase();
  clear(container);
  const recipes = [];
  for (const recipe of recipes) {
    if (!recipeToString(recipe).toLowerCase().includes(filter)) continue;
    recipes.push(a(recipe.path, recipe.name));
  }
  container.appendChild(ul(recipes));
}

const input = document.getElementById('filter');
const results = document.getElementById('results');

input.addEventListener('keyup', (event) => listRecipes(input.value, results));

})();
