(class_declaration
  (comment) @jsdoc
  body: (class_body
          (method_definition
            name: (property_identifier) @method.name))) (#not-eq? @method.name "constructor")
