(
  (comment)? @method.jsdoc
  (method_definition
    (accessibility_modifier)? @method.privacy
    name: (property_identifier) @method.name
    parameters: (formal_parameters (_
                                     pattern: [
                                               (identifier) @param.name
                                               (rest_pattern (identifier) @param.name)
                                               ]
                                     type: (type_annotation (_) @param.type))+ ) @params
    return_type: (type_annotation (_) @method.returns)?
    (#not-any-of? @method.name
      "constructor"
      "render"
      "connectedCallback"
      "disconnectedCallback"
      "attributeChangedCallback"
      "firstUpdated"
      "update"
      "updated"
      "willUpdated"
      "getUpdateComplete"))
)
