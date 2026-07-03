### Breaking Change Report

| Severity | Count |
|----------|------:|
| Breaking | 4 |
| Dangerous | 2 |
| Safe | 6 |

#### Breaking (4)

| Element | Change |
|---------|--------|
| `<my-element>` | event "change" type changed from CustomEvent to Event in <my-element> |
| `<my-element>` | method "doStuff" parameters changed from (name: string) to (name: string, value: number) in <my-element> |
| `<my-element>` | method "doStuff" return type changed from void to Promise<void> in <my-element> |
| `<my-element>` | slot "header" removed from <my-element> |

#### Dangerous (2)

| Element | Change |
|---------|--------|
| `<my-element>` | CSS custom property "--bg-color" default changed from white to black in <my-element> |
| `<my-element>` | field "value" type changed from string to number in <my-element> |

#### Safe (6)

| Element | Change |
|---------|--------|
| `<my-element>` | attribute "color" added to <my-element> |
| `<my-element>` | CSS custom property "--text-color" added to <my-element> |
| `<my-element>` | CSS part "icon" added to <my-element> |
| `<my-element>` | event "input" added to <my-element> |
| `<my-element>` | method "reset" added to <my-element> |
| `<new-element>` | element <new-element> added |
