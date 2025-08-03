export class DocsSidebar extends HTMLElement {
  static is = 'docs-sidebar';
  static { customElements.define(this.is, this); }

  get linkClass() {
    return this.getAttribute('link-class') || 'section_link';
  }

  get titleClass() {
    return this.getAttribute('title-class') || 'section_title';
  }

  connectedCallback() {
    this.#initializeSidebar();
    this.#setupScrollTracking();
    this.#featureHeading();
  }

  #initializeSidebar() {
    for (const toc of this.querySelectorAll('nav') ?? []) {
      toc.id ??= "";
      toc.classList.add('toc');
      const [fst] = toc.children
      for (const item of fst?.children ?? []) {
        item.classList.add('toc_item');
        item.firstElementChild.classList.add('toc_link');
      }
    }
  }

  #setupScrollTracking() {
    // Find the ToC for the current page by looking for the one that follows the active section_link
    const activePageLink = this.querySelector(`.${this.linkClass}.active`);
    let currentToc = null;

    if (activePageLink) {
      // Look for the next sibling nav element (the ToC for this page)
      let nextElement = activePageLink.nextElementSibling;
      while (nextElement) {
        if (nextElement.tagName === 'NAV' && nextElement.classList.contains('toc')) {
          currentToc = nextElement;
          break;
        }
        nextElement = nextElement.nextElementSibling;
      }
    }

    if (currentToc) {
      const pageInternalLinks = Array.from(currentToc.querySelectorAll('a[href*="#"]'));
      const pageLinks = pageInternalLinks.map(link => {
        return document.getElementById(decodeURIComponent(link.hash.replace('#', '')));
      }).filter(Boolean); // Remove null entries

      if (pageLinks.length > 0) {
        const scrollHandler = () => {
          let position = window.scrollY + window.innerHeight / 2;
          let activeIndex = 0;

          for (const [index, element] of pageLinks.entries()) {
            if (element && element.offsetTop < position && element.offsetTop > pageLinks[activeIndex].offsetTop) {
              activeIndex = index;
            }
          }

          this.#updateActiveHeading(activeIndex, pageInternalLinks);
        };

        window.addEventListener('scroll', scrollHandler);

        // Trigger initial active heading setup
        scrollHandler();
      }
    }
  }

  #updateActiveHeading(index, listLinks) {
    const linksToModify = {
      active: listLinks.filter(link => link.classList.contains('active'))[0],
      new: listLinks[index]
    };

    if (linksToModify.active !== linksToModify.new) {
      if (linksToModify.active) {
        linksToModify.active.classList.remove('active');
      }
      linksToModify.new.classList.add('active');
    }
  }

  #featureHeading() {
    // Show active heading at top
    let activeHeading = this.querySelector(`.${this.linkClass}.active`);
    activeHeading = activeHeading || this.querySelector(`.${this.titleClass}.active`);

    if (activeHeading) {
      this.scroll({
        top: activeHeading.offsetTop,
        left: 0,
      });
    }
  }
}
