// get object at key in local storage.
//
// Classes will need to be reconstructed since they are demoted to JSON objects.
function get(key) {
  return JSON.parse(localStorage.getItem(key));
}

// set object at key to value in local storage.
function set(key, value) {
  localStorage.setItem(key, JSON.stringify(value));
}
