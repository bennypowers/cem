package types

// Registry interface for accessing custom elements manifest data
// This allows tools to access registry without circular dependency
type Registry interface {
	ElementInfo(tagName string) (ElementInfo, error)
	AllElements() map[string]ElementInfo
	LoadManifests() error
}

// ElementInfo represents element information as needed by tools
type ElementInfo interface {
	TagName() string
	Name() string
	Description() string
	Module() string
	Package() string
	Attributes() []Attribute
	Slots() []Slot
	Events() []Event
	CssProperties() []CssProperty
	CssParts() []CssPart
	CssStates() []CssState
	Guidelines() []string
	Examples() []Example
	ItemsByKind(kind string) []Item
}

// Item interfaces for different element features
type Item interface {
	Name() string
	Description() string
	Guidelines() []string
	Examples() []string
	Kind() string
}

type Attribute interface {
	Item
	Type() string
	Default() string
	Required() bool
	Values() []string
}

type Slot interface {
	Item
}

type Event interface {
	Item
	Type() string
}

type CssProperty interface {
	Item
	Syntax() string
	Inherits() bool
	Initial() string
}

type CssPart interface {
	Item
}

type CssState interface {
	Item
}

type Example interface {
	Title() string
	Description() string
	Code() string
	Language() string
}

// ExampleInfo represents usage examples
type ExampleInfo struct {
	TitleValue       string `json:"title"`
	DescriptionValue string `json:"description,omitempty"`
	CodeValue        string `json:"code"`
	LanguageValue    string `json:"language,omitempty"`
}

func (e ExampleInfo) Title() string       { return e.TitleValue }
func (e ExampleInfo) Description() string { return e.DescriptionValue }
func (e ExampleInfo) Code() string        { return e.CodeValue }
func (e ExampleInfo) Language() string    { return e.LanguageValue }
