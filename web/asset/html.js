// ul from list.
//
// HTMLElements will be appended to each li and anything else is set as the
// innerHTML.
function ul(list) {
  const u = document.createElement('ul');
  for (const item of list) {
    const li = document.createElement('li');
    if (item instanceof HTMLElement) li.appendChild(item);
    else li.innerHTML = item;
    u.appendChild(li);
  }
  return u;
}

// h2 with name.
function h2(name) {
  const h = document.createElement('h2');
  h.innerHTML = name;
  return h;
}

// h3 with name.
function h3(name) {
  const h = document.createElement('h3');
  h.innerHTML = name;
  return h;
}

// clear the container.
function clear(container) {
  if (!container) return;
  while (container.firstChild) container.removeChild(container.firstChild);
}

// remove the node from the page.
function remove(node) {
  if (!node || !node.parentNode) return;
  node.parentNode.removeChild(node);
}

// prepend the item beore the container.
function prepend(container, item) {
  if (!container) return;
  container.parentNode.insertBefore(item, container);
}

// checkBox which performs the checkedAction when checked and the
// uncheckedAction when unchecked.
function checkBox(checkedAction, uncheckedAction) {
  const input = document.createElement('input');
  input.type = 'checkbox';
  input.addEventListener('change', function() {
    if (this.checked) checkedAction();
    else uncheckedAction();
  })
  return input;
}

// a with target href and innerHTML name.
function a(href, name) {
  const l = document.createElement('a');
  l.href = href;
  l.innerHTML = name;
  return l;
}
