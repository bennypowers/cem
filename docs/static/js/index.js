import {
  translations
} from './config.js';

// Navigation toggle
function initializeNavigation() {
  const navToggle = document.getElementById('nav-toggle');
  navToggle?.addEventListener('click', function() {
    this.parentElement?.querySelector('nav')?.classList.toggle('show');
  });
}

function updateDate() {
  const date = new Date();
  const year = date.getFullYear().toString();
  const year_el = document.querySelector('.year');
  if (year_el) {
    year_el.textContent = year;
  }
}

function markExternalLinks() {
  let links = document.querySelectorAll('a');
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
        attr1 = link.getAttribute(target);
        attr2 = link.getAttribute(rel);

        attr1 ? false : link.setAttribute(target, blank);
        attr2 ? false : link.setAttribute(rel, noopener);
      }
    });
  }
}

function copyFeedback(parent) {
  const copy_txt = document.createElement('div');
  const yanked = 'link_yanked';
  copy_txt.classList.add(yanked);
  copy_txt.innerText = translations.copied_text;
  if(!parent.querySelector(`.${yanked}`)) {
    const icon = parent.getElementsByTagName('svg')?.[0];
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
  deeplinks = document.querySelectorAll(`.${deeplink}`);
  if(deeplinks) {
    document.addEventListener('click', function(event)
    {
      target = event.target;
      parent = target.parentNode;
      if (target && target.classList.contains(deeplink) || parent.classList.contains(deeplink)) {
        event.preventDefault();
        new_link = target.href != undefined ? target.href : target.parentNode.href;
        navigator.clipboard.writeText(new_link);
        target.href != undefined ?  copyFeedback(target) : copyFeedback(target.parentNode);
      }
    });
  }
}

function makeTablesResponsive() {
  const tables = document.querySelectorAll('table');
  if (tables) {
    tables.forEach(function(table){
      const table_wrapper = document.createElement('div');
      table_wrapper.classList.add('scrollable');
      // Wrap table with table_wrapper
      table.parentNode.insertBefore(table_wrapper, table);
      table_wrapper.appendChild(table);
    });
  }
}

function backToTop(){
  const toTop = document.querySelector("#toTop");
  if (toTop) {
    window.addEventListener("scroll", () => {
      const last_known_scroll_pos = window.scrollY;
      if(last_known_scroll_pos >= 200) {
        toTop.style.display = "flex";
        toTop.classList.add('active');
      } else {
        toTop.classList.remove('active');
      }
    });
  }
}

function lazyLoadMedia(elements = []) {
  elements.forEach(element => {
    let media_items = document.querySelectorAll(element);
    if(media_items) {
      Array.from(media_items).forEach(function(item) {
        item.loading = "lazy";
      });
    }
  })
}

function loadActions() {
  updateDate();
  markExternalLinks();
  copyHeadingLink();
  makeTablesResponsive();
  backToTop();
  lazyLoadMedia(['iframe', 'img']);
}

// Critical path - initialize immediately
initializeNavigation();
// Use requestIdleCallback for non-critical tasks
if ('requestIdleCallback' in window) {
  requestIdleCallback(() => {
    loadActions();
  }, { timeout: 1000 });
} else {
  // Fallback for browsers without requestIdleCallback
  setTimeout(() => {
    loadActions();
  }, 0);
}
