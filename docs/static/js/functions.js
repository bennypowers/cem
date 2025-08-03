import { doc } from './variables.js';

// Object validation
export function isObj(obj) {
  return (obj && typeof obj === 'object' && obj !== null) ? true : false;
}

// DOM creation and manipulation
export function createEl(element = 'div') {
  return document.createElement(element);
}

export function emptyEl(el) {
  while(el.firstChild)
  el.removeChild(el.firstChild);
}

export function elem(selector, parent = document){
  let elem = isObj(parent) ? parent.querySelector(selector) : false;
  return elem ? elem : false;
}

export function elems(selector, parent = document) {
  return isObj(parent) ? parent.querySelectorAll(selector) : [];
}

// Class manipulation
export function pushClass(el, targetClass) {
  if (isObj(el) && targetClass) {
    let elClass = el.classList;
    elClass.contains(targetClass) ? false : elClass.add(targetClass);
  }
}

export function deleteClass(el, targetClass) {
  if (isObj(el) && targetClass) {
    let elClass = el.classList;
    elClass.contains(targetClass) ? elClass.remove(targetClass) : false;
  }
}

export function modifyClass(el, targetClass) {
  if (isObj(el) && targetClass) {
    const elClass = el.classList;
    elClass.contains(targetClass) ? elClass.remove(targetClass) : elClass.add(targetClass);
  }
}

export function containsClass(el, targetClass) {
  if (isObj(el) && targetClass && el !== document ) {
    return el.classList.contains(targetClass) ? true : false;
  }
}

// Element relationships
export function isChild(node, parentClass) {
  let objectsAreValid = isObj(node) && parentClass && typeof parentClass == 'string';
  return (objectsAreValid && node.closest(parentClass)) ? true : false;
}

// Attribute handling
export function elemAttribute(elem, attr, value = null) {
  if (value) {
    elem.setAttribute(attr, value);
  } else {
    value = elem.getAttribute(attr);
    return value ? value : false;
  }
}

// String utilities
export function deleteChars(str, subs) {
  let newStr = str;
  if (Array.isArray(subs)) {
    for (let i = 0; i < subs.length; i++) {
      newStr = newStr.replace(subs[i], '');
    }
  } else {
    newStr = newStr.replace(subs, '');
  }
  return newStr;
}

export function isBlank(str) {
  return (!str || str.trim().length === 0);
}

// Element matching
export function isMatch(element, selectors) {
  if(isObj(element)) {
    if(selectors.isArray) {
      let matching = selectors.map(function(selector){
        return element.matches(selector)
      })
      return matching.includes(true);
    }
    return element.matches(selectors)
  }
}

// Math utilities
export function closestInt(goal, collection) {
  return collection.reduce(function (prev, curr) {
    return (Math.abs(curr - goal) < Math.abs(prev - goal) ? curr : prev);
  });
}

// Class checking
export function hasClasses(el) {
  if(isObj(el)) {
    const classes = el.classList;
    return classes.length
  }
}

// Element wrapping
export function wrapEl(el, wrapper) {
  el.parentNode.insertBefore(wrapper, el);
  wrapper.appendChild(el);
}

export function wrapText(text, context, wrapper = 'mark') {
  let open = `<${wrapper}>`;
  let close = `</${wrapper}>`;
  let escapedOpen = `%3C${wrapper}%3E`;
  let escapedClose = `%3C/${wrapper}%3E`;
  function wrap(context) {
    let c = context.innerHTML;
    let pattern = new RegExp(text, "gi");
    let matches = text.length ? c.match(pattern) : null;

    if(matches) {
      matches.forEach(function(matchStr){
        c = c.replaceAll(matchStr, `${open}${matchStr}${close}`);
        context.innerHTML = c;
      });

      const images = elems('img', context);

      if(images) {
        images.forEach(image => {
          image.src = image.src.replaceAll(open, '').replaceAll(close, '').replaceAll(escapedOpen, '').replaceAll(escapedClose, '');
        });
      }
    }
  }

  const contents = ["h1", "h2", "h3", "h4", "h5", "h6", "p", "code", "td"];

  contents.forEach(function(c){
    const cs = elems(c, context);
    if(cs.length) {
      cs.forEach(function(cx, index){
        if(cx.children.length >= 1) {
          Array.from(cx.children).forEach(function(child){
            wrap(child);
          })
          wrap(cx);
        } else {
          wrap(cx);
        }
        // sanitize urls and ids
      });
    }
  });

  const hyperLinks = elems('a');
  if(hyperLinks) {
    hyperLinks.forEach(function(link){
      link.href = link.href.replaceAll(encodeURI(open), "").replaceAll(encodeURI(close), "");
    });
  }
}

// Type conversion
export function parseBoolean(string = "") {
  string = string.trim().toLowerCase();
  switch (string) {
    case 'true':
      return true;
    case 'false':
      return false;
    default:
      return undefined;
  }
}

// Icon loading
export function loadSvg(icon, parent) {
  parent.innerHTML = `
  <svg class="icon_${icon}">
    <use xlink:href="#${icon}"></use>
  </svg>`;
}

// Clipboard utilities
export function copyToClipboard(str) {
  let copy, selection, selected;
  copy = createEl('textarea');
  copy.value = str;
  copy.setAttribute('readonly', '');
  copy.style.position = 'absolute';
  copy.style.left = '-9999px';
  selection = document.getSelection();
  doc.appendChild(copy);
  // check if there is any selected content
  selected = selection.rangeCount > 0 ? selection.getRangeAt(0) : false;
  copy.select();
  document.execCommand('copy');
  doc.removeChild(copy);
  if (selected) { // if a selection existed before copying
    selection.removeAllRanges(); // unselect existing selection
    selection.addRange(selected); // restore the original selection
  }
}