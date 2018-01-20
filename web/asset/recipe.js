(function() {

const ingredientsContainer = document.getElementById('ingredients');
const stepsContainer = document.getElementById('steps');

if (ingredientsContainer) {
  strikethroughList(ingredientsContainer.children[1]);
}

if (stepsContainer) {
  strikethroughList(stepsContainer.children[1]);
}

})()
