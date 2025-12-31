export default {
  manifestSrc: "./custom-elements.json",
  
  diagnosticSeverity: {
    invalidBoolean: 'error',
    invalidNumber: 'error',
    invalidAttributeValue: 'error',
    deprecatedAttribute: 'warning',
    deprecatedElement: 'warning',
    duplicateAttribute: 'error'
  }
};