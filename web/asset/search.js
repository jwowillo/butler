(function() {

// BUG: Bookmarks not implemented.

function recipeToString(recipe) {
  var out = recipe.name;
  out += recipe.description;
  for (const ingredient of recipe.ingredients) {
    out += ingredient;
  }
  for (const step of recipe.steps) {
    out += step;
  }
  return out;
}

function listRecipes(filter, container) {
  filter = filter.toLowerCase();
  while (container.firstChild) {
    container.removeChild(container.firstChild);
  }
  const ul = document.createElement('ul');
  for (const recipe of recipes) {
    if (!recipeToString(recipe).toLowerCase().includes(filter)) {
      continue;
    }
    const a = document.createElement('a');
    a.appendChild(document.createTextNode(recipe.name));
    a.href = recipe.path;
    const li = document.createElement('li');
    li.appendChild(a);
    ul.appendChild(li);
  }
  container.appendChild(ul);
}

const input = document.getElementById('filter');
const results = document.getElementById('results');

input.addEventListener('keyup', (event) => listRecipes(input.value, results));

})();
