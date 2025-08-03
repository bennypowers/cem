import {
  hash,
  active
} from './variables.js';

import {
  translations
} from './config.js';

import {
  elem,
  elems,
  pushClass,
  deleteClass,
  containsClass,
  elemAttribute,
  createEl,
  wrapEl,
  copyToClipboard
} from './functions.js';

import { initializeCodeBlocks } from './code.js';

// Navigation toggle
function initializeNavigation() {
  const navToggle = document.getElementById('nav-toggle');
  if (navToggle) {
    navToggle.addEventListener('click', function() {
      this.parentElement.querySelector('nav').classList.toggle('show');
    });
  }
}

function featureHeading(){
  // show active heading at top.
  const link_class = "section_link";
  const title_class = "section_title";
  const parent = elem(".aside");
  if(parent) {
    let active_heading = elem(`.${link_class}.active`);
    active_heading = active_heading ? active_heading : elem(`.${title_class}.active}`);
    if (active_heading) {
      parent.scroll({
        top: active_heading.offsetTop,
        left: 0,
        // behavior: 'smooth'
      });
    }
  }
}

function activeHeading(index, list_links) {
  let links_to_modify = Object.create(null);
  links_to_modify.active = list_links.filter(function(link) {
    return containsClass(link, active);
  })[0];

  links_to_modify.new = list_links[index]

  if (links_to_modify.active != links_to_modify.new) {
    links_to_modify.active ? deleteClass(links_to_modify.active, active): false;
    pushClass(links_to_modify.new, active);
  }
}

function updateDate() {
  const date = new Date();
  const year = date.getFullYear().toString();
  const year_el = elem('.year');
  if (year_el) {
    year_el.textContent = year;
  }
}

function customizeSidebar() {
  const tocActive = 'toc_active';
  const aside = elem('aside');
  const tocs = elems('nav', aside);
  if(tocs) {
    tocs.forEach(function(toc){
      toc.id = "";
      pushClass(toc, 'toc');
      if(toc.children.length >= 1) {
        const toc_items = Array.from(toc.children[0].children);

        const previous_heading = toc.previousElementSibling;
        previous_heading.matches('.active') ? pushClass(toc, tocActive) : false;

        toc_items.forEach(function(item){
          pushClass(item, 'toc_item');
          pushClass(item.firstElementChild, 'toc_link');
        });
      }
    });

    const current_toc = elem(`.${tocActive}`);

    if(current_toc) {
      const page_internal_links = Array.from(elems('a', current_toc));

      const page_links = page_internal_links.map(function(link){
        return document.getElementById(decodeURIComponent(link.hash.replace('#','')));
      });

      window.addEventListener('scroll', function() {
        let position = window.scrollY + window.innerHeight/2;
        let active_index = 0;
        for (const [index, element] of page_links.entries()) {
          if(element && element.offsetTop < position && element.offsetTop > page_links[active_index].offsetTop) {
            active_index = index;
          }
        }
        activeHeading(active_index, page_internal_links);
      });
    }
  }
}

function markExternalLinks() {
  let links = elems('a');
  if(links) {
    Array.from(links).forEach(function(link){
      let target, rel, blank, noopener, attr1, attr2, url, is_external;
      try {
        url = new URL(link.href);
        // definition of same origin: RFC 6454, section 4 (https://tools.ietf.org/html/rfc6454#section-4)
        is_external = url.host !== location.host || url.protocol !== location.protocol || url.port !== location.port;
      } catch {}
      if(is_external) {
        target = 'target';
        rel = 'rel';
        blank = '_blank';
        noopener = 'noopener';
        attr1 = elemAttribute(link, target);
        attr2 = elemAttribute(link, noopener);

        attr1 ? false : elemAttribute(link, target, blank);
        attr2 ? false : elemAttribute(link, rel, noopener);
      }
    });
  }
}

function sanitizeURL(url) {
  // removes any existing id on url
  const position_of_hash = url.indexOf(hash);
  if(position_of_hash > -1 ) {
    const id = url.substr(position_of_hash, url.length - 1);
    url = url.replace(id, '');
  }
  return url
}

function copyFeedback(parent) {
  const copy_txt = document.createElement('div');
  const yanked = 'link_yanked';
  copy_txt.classList.add(yanked);
  copy_txt.innerText = translations.copied_text;
  if(!elem(`.${yanked}`, parent)) {
    const icon = parent.getElementsByTagName('svg')[0];
    if (icon) {
      const original_src = icon.src;
      icon.src = '/icons/check.svg';
      parent.appendChild(copy_txt);
      setTimeout(function() {
        parent.removeChild(copy_txt)
        icon.src = original_src;
      }, 2250);
    }
  }
}

function copyHeadingLink() {
  let deeplink, deeplinks, new_link, parent, target;
  deeplink = 'link';
  deeplinks = elems(`.${deeplink}`);
  if(deeplinks) {
    document.addEventListener('click', function(event)
    {
      target = event.target;
      parent = target.parentNode;
      if (target && containsClass(target, deeplink) || containsClass(parent, deeplink)) {
        event.preventDefault();
        new_link = target.href != undefined ? target.href : target.parentNode.href;
        copyToClipboard(new_link);
        target.href != undefined ?  copyFeedback(target) : copyFeedback(target.parentNode);
      }
    });
  }
}

function makeTablesResponsive() {
  const tables = elems('table');
  if (tables) {
    tables.forEach(function(table){
      const table_wrapper = createEl();
      pushClass(table_wrapper, 'scrollable');
      wrapEl(table, table_wrapper);
    });
  }
}

function backToTop(){
  const toTop = elem("#toTop");
  if (toTop) {
    window.addEventListener("scroll", () => {
      const last_known_scroll_pos = window.scrollY;
      if(last_known_scroll_pos >= 200) {
        toTop.style.display = "flex";
        pushClass(toTop, active);
      } else {
        deleteClass(toTop, active);
      }
    });
  }
}

function lazyLoadMedia(elements = []) {
  elements.forEach(element => {
    let media_items = elems(element);
    if(media_items) {
      Array.from(media_items).forEach(function(item) {
        item.loading = "lazy";
      });
    }
  })
}

function loadActions() {
  updateDate();
  customizeSidebar();
  markExternalLinks();
  copyHeadingLink();
  makeTablesResponsive();
  backToTop();
  lazyLoadMedia(['iframe', 'img']);
}

// Main initialization function
export function initializeApp() {
  initializeNavigation();
  initializeCodeBlocks();
  loadActions();
  
  // Feature heading with delay
  setTimeout(() => {
    featureHeading();
  }, 50);
}

// Auto-initialize when DOM is loaded
if (document.readyState === 'loading') {
  document.addEventListener('DOMContentLoaded', initializeApp);
} else {
  initializeApp();
}
