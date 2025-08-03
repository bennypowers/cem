import { translations } from './config.js';

// Navigation toggle
function initializeNavigation() {
  const navToggle = document.getElementById('nav-toggle');
  navToggle?.addEventListener('click', function() {
    this.parentElement?.querySelector('nav')?.classList.toggle('show');
  });
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

// Tables made responsive with CSS: table { overflow-x: auto; }

function backToTop(){
  const toTop = document.querySelector("#toTop");
  if (toTop) {
    window.addEventListener("scroll", () => {
      toTop.classList.toggle('active', window.scrollY >= 200);
    }, { passive: true });
  }
}

// Lazy loading handled by native HTML loading="lazy" attribute

function loadActions() {
  copyHeadingLink();
  backToTop();
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
