(function() {

function recipeToString(recipe) {
  let out = recipe.name + recipe.description;
  for (const ingredient of recipe.ingredients) out += ingredient;
  for (const step of recipe.steps) out += step;
  return out;
}

function listRecipes(container, filter) {
  const map = {};
  for (const recipe of recipes) map[recipe.name] = recipe;
  set('filter', filter);
  filter = filter.toLowerCase();
  for (const recipe of container.getElementsByTagName('li')) {
    const link = recipe.getElementsByTagName('a')[0].innerHTML;
    if (recipeToString(map[link]).toLowerCase().includes(filter)) {
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
