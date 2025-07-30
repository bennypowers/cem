package manifest

// Interface implementation checks
var _ Deprecatable = (*PropertyLike)(nil)

// PropertyLike is the common interface of variables, class fields, and function parameters.
type PropertyLike struct {
	FullyQualified
	StartByte  uint       `json:"-"`
	Type       *Type      `json:"type,omitempty"`
	Default    string     `json:"default,omitempty"`
	Deprecated Deprecated `json:"deprecated,omitempty"`
	Readonly   bool       `json:"readonly,omitempty"`
}

func (x *PropertyLike) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

// Clone creates a deep copy of the PropertyLike structure.
func (p PropertyLike) Clone() PropertyLike {
	cloned := PropertyLike{
		FullyQualified: p.FullyQualified.Clone(),
		StartByte:      p.StartByte,
		Default:        p.Default,
		Readonly:       p.Readonly,
	}

	if p.Deprecated != nil {
		cloned.Deprecated = p.Deprecated.Clone()
	}

	if p.Type != nil {
		cloned.Type = p.Type.Clone()
	}

	return cloned
}
