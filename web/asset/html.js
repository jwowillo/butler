// ul from list.
//
// HTMLElements will be appended to each li and anything else is set as the
// innerHTML.
function ul(list) {
  const u = document.createElement('ul');
  for (const item of list) {
    const li = document.createElement('li');
    if (item instanceof HTMLElement) {
      li.appendChild(item);
    } else {
      const span = document.createElement('span');
      span.innerHTML = item;
      li.appendChild(span);
    }
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

// checkboxList creates a checkboxList over the list elements in the container
// that does checked when an item is checked and unchecked when an item is
// unchecked.
//
// checked and unchecked are passed the items that are checked.
function checkboxList(container, checked, unchecked) {
  const getChecked = () => get(container.id) || [];
  const setChecked = items => set(container.id, items);
  for (const item of container.getElementsByTagName('li')) {
    const first = item.firstChild;
    const input = checkbox(
      function() {
        let items = getChecked();
        items.push(first.innerHTML);
        setChecked(items);
        checked(items);
      },
      function() {
        let items = getChecked();
        items = items.filter((item) => item != first.innerHTML);
        setChecked(items);
        unchecked(items);
      }
    );
    prepend(first, input);
  }
  const items = new Set();
  for (const item of getChecked()) items.add(item);
  for (const item of container.getElementsByTagName('li')) {
    const input = item.children[0];
    const node = item.children[1];
    if (items.has(node.innerHTML)) {
      input.checked = true
    };
  }
  checked(Array.from(items));
}

// button with text that does cmd when clicked.
function button(text, cmd) {
  const b = document.createElement('button');
  b.innerHTML = text;
  b.addEventListener('click', cmd);
  b.className = text.toLowerCase();
  return b;
}

// checkbox which performs the checkedAction when checked and the
// uncheckedAction when unchecked.
function checkbox(checked, unchecked) {
  const input = document.createElement('input');
  input.type = 'checkbox';
  input.addEventListener('change', function() {
    if (this.checked) checked();
    else unchecked();
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
