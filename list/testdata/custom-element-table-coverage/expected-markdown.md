# <root>

# module src/base-elem.js

# `<base-elem>`

## Attributes

Name             | Summary            
---------------- | -------------------
`inherited-attr` | Inherited attribute

## Slots

Name             | Summary       
---------------- | --------------
`inherited-slot` | Inherited slot

## Events

Name              | Summary        
----------------- | ---------------
`inherited-event` | Inherited event

## CSS Properties

Name                | Syntax    | Summary        
------------------- | --------- | ---------------
`--inherited-color` | `<color>` | Inherited color

## CSS Parts

Name             | Summary       
---------------- | --------------
`inherited-part` | Inherited part


# module src/test-elem.js

# `<test-elem>`

## Attributes

Name             | DOM Property | Reflects | Summary             | Default | Type     | Inherited From
---------------- | ------------ | -------- | ------------------- | ------- | -------- | --------------
`inherited-attr` |              |          | Inherited attribute |         |          | BaseElem      
`foo`            | `foo`        | ✅        | Foo summary         | `bar`   | `string` |               
`bar`            | `bar`        |          |                     |         |          |               

## Slots

Name             | Summary           | Inherited From
---------------- | ----------------- | --------------
`inherited-slot` | Inherited slot    | BaseElem      
`<default>`      | Default slot      |               
`icon`           | Slot for the icon |               

## Events

Name              | Type                       | Summary         | Inherited From
----------------- | -------------------------- | --------------- | --------------
`inherited-event` |                            | Inherited event | BaseElem      
`submitted`       | `CustomEvent<'a'|'b'|'c'>` | Fires on submit |               
`reset`           |                            |                 |               

## Fields

Name  | Summary          
----- | -----------------
`foo` | Foo field summary
`bar` |                  

## Methods

Name           | Return Type            | Privacy | Static  | Summary     
-------------- | ---------------------- | ------- | ------- | ------------
`doThing`      | `void | Promise<void>` | public  | `false` | Does a thing
`privateStuff` | `string`               | private | `true`  |             

## CSS Properties

Name                | Syntax    | Default | Summary         | Inherited From
------------------- | --------- | ------- | --------------- | --------------
`--inherited-color` | `<color>` |         | Inherited color | BaseElem      
`--main-color`      | `<color>` | `#fff`  | Main color      |               
`--unused`          |           |         |                 |               

## CSS Parts

Name             | Summary        | Inherited From
---------------- | -------------- | --------------
`inherited-part` | Inherited part | BaseElem      
`label`          | The label part |               

## CSS States

Name         | Summary     
------------ | ------------
`--active`   | Active state
`--inactive` |             

