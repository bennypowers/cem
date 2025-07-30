package manifest

// SourceReference is a reference to the source of a declaration or member.
type SourceReference struct {
	Href string `json:"href"`
}

// Clone creates a deep copy of the SourceReference.
func (s SourceReference) Clone() SourceReference {
	return SourceReference{
		Href: s.Href,
	}
}
