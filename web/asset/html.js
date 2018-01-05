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

function h2(name) {
  const h = document.createElement('h2');
  h.innerHTML = name;
  return h;
}

function h3(name) {
  const h = document.createElement('h3');
  h.innerHTML = name;
  return h;
}

function clear(container) {
  if (!container) return;
  while (container.firstChild) container.removeChild(container.firstChild);
}

function remove(node) {
  if (!node || !node.parentNode) return;
  node.parentNode.removeChild(node);
}

function prepend(container, item) {
  if (!container) return;
  container.parentNode.insertBefore(item, container);
}

function checkBox(checkedAction, uncheckedAction) {
  const input = document.createElement('input');
  input.type = 'checkbox';
  input.addEventListener('change', function() {
    if (this.checked) checkedAction();
    else uncheckedAction();
  })
  return input;
}

function a(href, name) {
  const l = document.createElement('a');
  l.href = href;
  l.innerHTML = name;
  return l;
}
