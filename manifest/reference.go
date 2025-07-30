package manifest

// Reference to an export of a module.
type Reference struct {
	Name    string `json:"name"`
	Package string `json:"package,omitempty"`
	Module  string `json:"module,omitempty"`
}

func NewReference(name string, pkg string, module string) *Reference {
	return &Reference{
		Name:    name,
		Package: pkg,
		Module:  normalizePath(module),
	}
}

// Clone creates a deep copy of the Reference.
func (r Reference) Clone() Reference {
	return Reference{
		Name:    r.Name,
		Package: r.Package,
		Module:  r.Module,
	}
}
