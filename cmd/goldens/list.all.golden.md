# <root>

# module src/test-elem.js

# `<test-elem>`

## Attributes

Name  | DOM Property | Reflects | Summary     | Default
----- | ------------ | -------- | ----------- | -------
`foo` | `foo`        | ✅       | Foo summary | `bar`  
`bar` | `bar`        |          |             |        

## Slots

Name        | Summary          
----------- | -----------------
`<default>` | Default slot     
`icon`      | Slot for the icon

## Events

Name        | Type                       | Summary        
----------- | -------------------------- | ---------------
`submitted` | `CustomEvent<'a'|'b'|'c'>` | Fires on submit
`reset`     | `Event`                    |                

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

Name           | Syntax    | Default | Summary   
-------------- | --------- | ------- | ----------
`--main-color` | `<color>` | `#fff`  | Main color
`--unused`     |           |         |           

## CSS Parts

Name    | Summary       
------- | --------------
`label` | The label part

## CSS States

Name         | Summary     
------------ | ------------
`--active`   | Active state
`--inactive` |             

