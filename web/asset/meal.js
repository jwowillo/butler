(function() {

// BUG: This code is entirely too messy.

// BUG: Adding fractions doesn't work.
  // Make a fraction class that starts as undefined.
  // On first amount, set amount to the fraction.
  // Add fractions continously.
  // Use fractions tostring.

// BUG: Plurals don't combine.
  // Once items can be singular and their plural can be generated, can just keep
  // item representation in this code and itll be more simple.

// BUG: Bookmarks aren't implemented yet.

function addCheckBoxes(mealContainer, recipeContainer) {
  const map = {};
  for (const recipe of recipes) {
    map[recipe.name] = recipe;
  }
  let checked = [];
  for (const item of recipeContainer.getElementsByTagName('li')) {
    const link = item.firstChild;
    const input = document.createElement('input');
    input.type = 'checkbox';
    item.insertBefore(input, link);
    input.addEventListener('change', function() {
      if (this.checked) {
        checked.push(map[link.innerHTML]);
      } else {
        checked = checked.filter((recipe) => recipe.name != link.innerHTML);
      }
      makeMeal(mealContainer, recipeContainer, checked);
    });
  }
}

function isNumber(x) {
  if (x.indexOf('/') > -1) {
    return true;
  }
  return !isNaN(parseFloat(x)) && !isNaN(x - 0)
}

function makeMeal(mealContainer, recipeContainer, checked) {
  if (checked.length == 0) {
    recipeContainer.parentNode.removeChild(mealContainer);
  }
  if (checked.length == 1) {
    recipeContainer.parentNode.insertBefore(mealContainer, recipeContainer);
  }
  while (mealContainer.firstChild) {
    mealContainer.removeChild(mealContainer.firstChild);
  }
  const header = document.createElement('h3');
  header.innerHTML = 'Meal';
  mealContainer.appendChild(header);
  const ingredientsHeader = document.createElement('h2');
  ingredientsHeader.innerHTML = 'Ingredients:';
  const ingredientsList = document.createElement('ul');

  const used = new Set();
  for (const i in checked) {
    for (const ingredientA of checked[i].ingredients) {

      const partsA = ingredientA.split(' ');
      let amountA = -1, restA = '';
      if (!isNumber(partsA[0])) {
        restA = partsA[0];
      } else {
        amountA = eval(partsA[0]);
        restA = partsA.slice(1).join(' ');
      }

      if (used.has(restA)) continue;
      used.add(restA);

      for (const j in checked) {
        if (i == j) continue;
        for (const ingredientB of checked[j].ingredients) {
          const partsB = ingredientB.split(' ');
          let amountB = '', restB = '';
          if (!isNumber(partsB[0])) {
            restB = partsB[0];
          } else {
            amountB = eval(partsB[0]);
            restB = partsB.slice(1).join(' ');
          }
          if (amountA != -1 && amountB != -1 && restA == restB) {
            amountA += amountB;
          }
        }
      }

      const ingredient = document.createElement('li');
      if (amountA == -1) {
        ingredient.innerHTML = restA;
      } else {
        ingredient.innerHTML = amountA + ' ' + restA;
      }
      ingredientsList.appendChild(ingredient);
    }
  }

  mealContainer.appendChild(ingredientsHeader);
  mealContainer.appendChild(ingredientsList);
  const stepsHeader = document.createElement('h2');
  stepsHeader.innerHTML = 'Steps:';
  mealContainer.appendChild(stepsHeader);
  const steps = document.createElement('ul');
  for (const recipe of checked) {
    const link = document.createElement('a');
    link.href = recipe.path;
    const recipeTag = document.createElement('li');
    recipeTag.innerHTML = recipe.name;
    link.appendChild(recipeTag);
    steps.appendChild(link);
  }
  mealContainer.appendChild(steps);
}

const meal = document.createElement('div');
meal.id = 'box';
const results = document.getElementById('results');

addCheckBoxes(meal, results);

})();
