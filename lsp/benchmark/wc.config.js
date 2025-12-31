export default {
  // include: ['**/*.html', '**/*.js', '**/*.md'],
  // exclude: ['**/*.json'],
  // tagFormatter: (tag) => `${tag}-formatted`,
  // diagnosticSeverity: {
  //   invalidBoolean: 'warning',
  //   invalidNumber: 'error',
  //   invalidAttributeValue: 'error',
  //   deprecatedAttribute: 'warning',
  //   deprecatedElement: 'warning'
  // }
  // manifestSrc: "./_custom-elements.json",
  // manifestSrc:
  //   "https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.20.1/cdn/custom-elements.json",

  libraries: {
    "@shoelace-style": {
      // manifestSrc: 'https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.20.1/cdn/custom-elements.json'
      // manifestSrc: "./_custom-elements.json",
    },
    "@awesome.me/webawesome": {
      /**
       * Fetch manifest from a URL
       * This isn't needed if you have the NPM package installed
       */
      manifestSrc:
        "https://cdn.jsdelivr.net/npm/@awesome.me/webawesome@3.0.0-beta.4/dist/custom-elements.json",

      /**
       * Replace `wa-` prefix with `awesome-` for all Web Awesome components
       * Language server options will now work for `awesome-button` instead of `wa-button`
       */
      tagFormatter: (tagName) => tagName.replace("wa-", "awesome-"),

      diagnosticSeverity: {
        /** Duplicate attributes will show as a warning for Web Awesome components */
        duplicateAttribute: "warning",
      },
    },
  },
};
