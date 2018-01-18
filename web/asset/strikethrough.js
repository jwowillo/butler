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

function strikethroughList(container) {
  checkboxList(
    container,
    checked => strike(container, checked),
    checked => strike(container, checked)
  )
}

