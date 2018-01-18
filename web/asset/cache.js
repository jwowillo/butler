// get object at key in local storage.
//
// Classes will need to be reconstructed since they are demoted to JSON objects.
function get(key) {
  return JSON.parse(localStorage.getItem(window.location.href+key));
}

// set object at key to value in local storage.
function set(key, value) {
  localStorage.setItem(window.location.href+key, JSON.stringify(value));
}
