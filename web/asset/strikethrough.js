// strike through the checked items in container.
function strike(container, checked) {
  const set = new Set(checked);
  for (const item of container.getElementsByTagName('span')) {
    if (set.has(item.innerHTML)) {
      item.style.textDecoration = 'line-through';
    } else {
      item.style.textDecoration = 'none';
    }
  }
}

// strikethroughList puts check-boxes next to all container list elements that
// strike through the list element when checked.
function strikethroughList(container) {
  checkboxList(
    container,
    checked => strike(container, checked),
    checked => strike(container, checked)
  );
  container.appendChild(button('Clear', function() {
    set(container.id, []);
    strike(container, []);
    for (const item of container.getElementsByTagName('li')) {
      const input = item.children[0];
      input.checked = false;
    }
  }));
}

