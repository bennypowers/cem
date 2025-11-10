// Debug view transitions
if (window.navigation) {
  console.log('[view-transition] Navigation API supported');

  navigation.addEventListener('navigate', (event) => {
    console.log('[view-transition] Navigate event:', {
      destination: event.destination.url,
      canIntercept: event.canIntercept,
      navigationType: event.navigationType
    });
  });

  navigation.addEventListener('navigatesuccess', (event) => {
    console.log('[view-transition] Navigate success');
  });

  navigation.addEventListener('navigateerror', (event) => {
    console.log('[view-transition] Navigate error:', event.error);
  });
} else {
  console.log('[view-transition] Navigation API NOT supported');
}

// Check if document has view transition support
if (document.startViewTransition) {
  console.log('[view-transition] document.startViewTransition() supported (same-doc only)');
} else {
  console.log('[view-transition] document.startViewTransition() NOT supported');
}

// Monitor page transitions
const observer = new PerformanceObserver((list) => {
  for (const entry of list.getEntries()) {
    if (entry.entryType === 'navigation') {
      console.log('[view-transition] Navigation timing:', entry);
    }
  }
});
observer.observe({ entryTypes: ['navigation'] });

// Log on page load
console.log('[view-transition] Page loaded, checking computed styles on html...');
const htmlStyle = getComputedStyle(document.documentElement);
console.log('[view-transition] html background:', htmlStyle.background);

// Try to detect if view transition is happening via CSS
const testDiv = document.createElement('div');
testDiv.style.viewTransitionName = 'test-name';
document.body.appendChild(testDiv);
const testStyle = getComputedStyle(testDiv);
console.log('[view-transition] Test element view-transition-name:', testStyle.viewTransitionName);
testDiv.remove();
