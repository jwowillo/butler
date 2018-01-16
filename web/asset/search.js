(function() {

// recipeToString concatenates the recipe into a long string.
function recipeToString(recipe) {
  let out = recipe.name + recipe.description;
  for (const ingredient of recipe.ingredients) out += ingredient;
  for (const step of recipe.steps) out += step;
  return out;
}

// listRecipes lists all the recipes that match the filter in the container.
function listRecipes(container, filter) {
  set('filter', filter);
  if (filter == null) filter = '';
  filter = filter.toLowerCase();
  for (const recipe of container.getElementsByTagName('li')) {
    const link = recipe.getElementsByTagName('a')[0].innerHTML;
    if (recipeToString(recipes[link]).toLowerCase().includes(filter)) {
      recipe.style.display = 'block';
    } else {
      recipe.style.display = 'none';
    }
  }
}

const input = document.getElementById('filter');
const results = document.getElementById('results');

input.addEventListener('keyup', (event) => listRecipes(results, input.value));

input.value = get('filter');
listRecipes(results, get('filter'));

})();
