/*
Copyright © 2026 Benny Powers

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package generate

import (
	"fmt"
	"strings"
	"testing"

	M "bennypowers.dev/cem/manifest"
)

// Inline assertions: dropOverloadImplementationSignatures and
// dropOverloadImplementationMethods are Tier-1 pure functions, which
// docs/TESTING.md requires be covered by exhaustive table-driven in-package
// tests rather than fixture/golden files (same rationale as
// TestParameterFromNode). End-to-end overload dropping is covered by the
// overload golden fixtures.

func TestDropOverloadImplementationSignatures(t *testing.T) {
	// fnDecl builds a minimal function declaration; the filter reads only
	// the name and start byte.
	fnDecl := func(name string, startByte uint) *M.FunctionDeclaration {
		return &M.FunctionDeclaration{
			FunctionLike:   M.FunctionLike{StartByte: startByte},
			FullyQualified: M.FullyQualified{Name: name},
		}
	}
	// varDecl builds a minimal non-function declaration, proving that
	// interleaved declarations are untouched.
	varDecl := func(name string, startByte uint) *M.VariableDeclaration {
		return &M.VariableDeclaration{
			PropertyLike: M.PropertyLike{
				FullyQualified: M.FullyQualified{Name: name},
				StartByte:      startByte,
			},
		}
	}

	// valid overload set: two body-less signatures, an unrelated variable
	// interleaved, then the body-carrying implementation.
	greetOverloadA := fnDecl("greet", 0)
	greetOverloadB := fnDecl("greet", 100)
	countVar := varDecl("count", 150)
	greetImpl := fnDecl("greet", 200)

	// single function (group size 1 fast path).
	solo := fnDecl("solo", 0)

	// lone body-less declaration (a stray overload signature with no
	// implementation).
	loneSignature := fnDecl("declared", 0)

	// same-name group where no member carries a body.
	pairA := fnDecl("pair", 0)
	pairB := fnDecl("pair", 100)

	// body-carrying declaration not last: the last member carries a body,
	// but so does a preceding member, so the group is not overload-shaped.
	seqA := fnDecl("seq", 0)
	seqB := fnDecl("seq", 100)
	seqImpl := fnDecl("seq", 200)

	// implementation captured before its signatures: the exported and local
	// query loops append in separate passes, so slice order is not source
	// order across them.
	crossImpl := fnDecl("cross", 0)
	crossSigA := fnDecl("cross", 100)
	crossSigB := fnDecl("cross", 200)

	// implementation JSDoc merges onto the first kept signature.
	docsSig := fnDecl("docs", 100)
	docsImpl := fnDecl("docs", 200)
	docsImpl.Description = "impl description"
	docsImpl.Summary = "impl summary"

	tests := []struct {
		name          string
		declarations  []M.Declaration
		hasBody       map[*M.FunctionDeclaration]bool
		want          []M.Declaration
		wantSameSlice bool
		wantDesc      string
		wantSummary   string
	}{
		{
			name:         "valid overload set drops only the implementation signature",
			declarations: []M.Declaration{greetOverloadA, greetOverloadB, countVar, greetImpl},
			hasBody:      map[*M.FunctionDeclaration]bool{greetImpl: true},
			want:         []M.Declaration{greetOverloadA, greetOverloadB, countVar},
		},
		{
			name:          "single function kept unchanged",
			declarations:  []M.Declaration{solo},
			hasBody:       map[*M.FunctionDeclaration]bool{solo: true},
			want:          []M.Declaration{solo},
			wantSameSlice: true,
		},
		{
			name:          "lone body-less declaration kept",
			declarations:  []M.Declaration{loneSignature},
			hasBody:       map[*M.FunctionDeclaration]bool{},
			want:          []M.Declaration{loneSignature},
			wantSameSlice: true,
		},
		{
			name:          "all-body-less same-name group keeps everything",
			declarations:  []M.Declaration{pairA, pairB},
			hasBody:       map[*M.FunctionDeclaration]bool{pairA: false, pairB: false},
			want:          []M.Declaration{pairA, pairB},
			wantSameSlice: true,
		},
		{
			name:          "body-carrying declaration not last keeps everything",
			declarations:  []M.Declaration{seqA, seqB, seqImpl},
			hasBody:       map[*M.FunctionDeclaration]bool{seqA: true, seqImpl: true},
			want:          []M.Declaration{seqA, seqB, seqImpl},
			wantSameSlice: true,
		},
		{
			name:         "implementation before signatures still drops only the implementation",
			declarations: []M.Declaration{crossImpl, crossSigA, crossSigB},
			hasBody:      map[*M.FunctionDeclaration]bool{crossImpl: true},
			want:         []M.Declaration{crossSigA, crossSigB},
		},
		{
			name:         "implementation JSDoc merges onto the first kept signature",
			declarations: []M.Declaration{docsSig, docsImpl},
			hasBody:      map[*M.FunctionDeclaration]bool{docsImpl: true},
			want:         []M.Declaration{docsSig},
			wantDesc:     "impl description",
			wantSummary:  "impl summary",
		},
		{
			name:         "empty declarations returns empty",
			declarations: nil,
			hasBody:      map[*M.FunctionDeclaration]bool{},
			want:         nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dropOverloadImplementationSignatures(tt.declarations, tt.hasBody)
			if len(got) != len(tt.want) {
				t.Fatalf("dropOverloadImplementationSignatures() kept %d declarations (%s), want %d (%s)",
					len(got), declNames(got), len(tt.want), declNames(tt.want))
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("kept declaration %d = %q, want %q", i, got[i].Name(), tt.want[i].Name())
				}
			}
			if tt.wantDesc != "" || tt.wantSummary != "" {
				first := got[0].(*M.FunctionDeclaration)
				if first.Description != tt.wantDesc {
					t.Errorf("merged description = %q, want %q", first.Description, tt.wantDesc)
				}
				if first.Summary != tt.wantSummary {
					t.Errorf("merged summary = %q, want %q", first.Summary, tt.wantSummary)
				}
			}
			if tt.wantSameSlice {
				if &got[0] != &tt.declarations[0] {
					t.Errorf("dropOverloadImplementationSignatures() returned a new slice, want the input slice unchanged (no-drop fast path)")
				}
			}
		})
	}
}

func TestDropOverloadImplementationMethods(t *testing.T) {
	// method builds a minimal class method; the filter reads only the
	// name, static flag, and start byte.
	method := func(name string, static bool, startByte uint) *M.ClassMethod {
		return &M.ClassMethod{
			FunctionLike:   M.FunctionLike{StartByte: startByte},
			FullyQualified: M.FullyQualified{Name: name},
			Kind:           "method",
			Static:         static,
		}
	}
	// field builds a minimal non-method member, proving that interleaved
	// fields are untouched even when they share a method's name.
	field := func(name string, startByte uint) *M.CustomElementField {
		return &M.CustomElementField{
			ClassField: M.ClassField{
				Kind: "field",
				PropertyLike: M.PropertyLike{
					FullyQualified: M.FullyQualified{Name: name},
					StartByte:      startByte,
				},
			},
		}
	}

	// valid overload set: two body-less signatures, an unrelated field
	// interleaved, then the body-carrying implementation.
	fmtSigA := method("format", false, 0)
	fmtSigB := method("format", false, 100)
	countField := field("count", 150)
	fmtImpl := method("format", false, 200)

	// static overload set alongside a same-name instance method: static and
	// instance methods are distinct members, so the instance method must not
	// make the static group lose its overload shape.
	instFoo := method("foo", false, 0)
	staticFooSig := method("foo", true, 100)
	staticFooImpl := method("foo", true, 200)

	// same-name group where no member carries a body.
	pairA := method("pair", false, 0)
	pairB := method("pair", false, 100)

	// two members carry bodies, so the group is not overload-shaped.
	seqA := method("seq", false, 0)
	seqB := method("seq", false, 100)
	seqImpl := method("seq", false, 200)

	// implementation captured before its signatures: query match order is
	// not source order, so the body-carrying member is identified
	// position-independently.
	crossImpl := method("cross", false, 0)
	crossSigA := method("cross", false, 100)
	crossSigB := method("cross", false, 200)

	// implementation JSDoc merges onto the first kept signature.
	docsSig := method("docs", false, 100)
	docsImpl := method("docs", false, 200)
	docsImpl.Description = "impl description"
	docsImpl.Summary = "impl summary"
	docsImpl.Deprecated = M.DeprecatedReason("use docs2")

	// lone body-less declaration (a stray overload signature with no
	// implementation).
	loneSignature := method("declared", false, 0)

	// a field sharing a method's name never groups with the methods.
	shadowField := field("shadowed", 300)
	shadowSig := method("shadowed", false, 100)
	shadowImpl := method("shadowed", false, 200)

	tests := []struct {
		name           string
		members        []M.ClassMember
		hasBody        map[*M.ClassMethod]bool
		want           []M.ClassMember
		wantSameSlice  bool
		wantDesc       string
		wantSummary    string
		wantDeprecated M.Deprecated
	}{
		{
			name:    "valid overload set drops only the implementation signature",
			members: []M.ClassMember{fmtSigA, fmtSigB, countField, fmtImpl},
			hasBody: map[*M.ClassMethod]bool{fmtImpl: true},
			want:    []M.ClassMember{fmtSigA, fmtSigB, countField},
		},
		{
			name:    "static overload drops only the static implementation",
			members: []M.ClassMember{instFoo, staticFooSig, staticFooImpl},
			hasBody: map[*M.ClassMethod]bool{instFoo: true, staticFooImpl: true},
			want:    []M.ClassMember{instFoo, staticFooSig},
		},
		{
			name:          "single method kept unchanged",
			members:       []M.ClassMember{instFoo},
			hasBody:       map[*M.ClassMethod]bool{instFoo: true},
			want:          []M.ClassMember{instFoo},
			wantSameSlice: true,
		},
		{
			name:          "lone body-less method kept",
			members:       []M.ClassMember{loneSignature},
			hasBody:       map[*M.ClassMethod]bool{},
			want:          []M.ClassMember{loneSignature},
			wantSameSlice: true,
		},
		{
			name:          "all-body-less same-name group keeps everything",
			members:       []M.ClassMember{pairA, pairB},
			hasBody:       map[*M.ClassMethod]bool{pairA: false, pairB: false},
			want:          []M.ClassMember{pairA, pairB},
			wantSameSlice: true,
		},
		{
			name:          "two body-carrying same-name methods keep everything",
			members:       []M.ClassMember{seqA, seqB, seqImpl},
			hasBody:       map[*M.ClassMethod]bool{seqA: true, seqImpl: true},
			want:          []M.ClassMember{seqA, seqB, seqImpl},
			wantSameSlice: true,
		},
		{
			name:    "implementation before signatures still drops only the implementation",
			members: []M.ClassMember{crossImpl, crossSigA, crossSigB},
			hasBody: map[*M.ClassMethod]bool{crossImpl: true},
			want:    []M.ClassMember{crossSigA, crossSigB},
		},
		{
			name:           "implementation JSDoc merges onto the first kept signature",
			members:        []M.ClassMember{docsSig, docsImpl},
			hasBody:        map[*M.ClassMethod]bool{docsImpl: true},
			want:           []M.ClassMember{docsSig},
			wantDesc:       "impl description",
			wantSummary:    "impl summary",
			wantDeprecated: M.DeprecatedReason("use docs2"),
		},
		{
			name:    "field sharing a method name never groups with the methods",
			members: []M.ClassMember{shadowField, shadowSig, shadowImpl},
			hasBody: map[*M.ClassMethod]bool{shadowImpl: true},
			want:    []M.ClassMember{shadowField, shadowSig},
		},
		{
			name:    "empty members returns empty",
			members: nil,
			hasBody: map[*M.ClassMethod]bool{},
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dropOverloadImplementationMethods(tt.members, tt.hasBody)
			if len(got) != len(tt.want) {
				t.Fatalf("dropOverloadImplementationMethods() kept %d members (%s), want %d (%s)",
					len(got), memberNames(got), len(tt.want), memberNames(tt.want))
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("kept member %d = %q, want %q", i, memberName(got[i]), memberName(tt.want[i]))
				}
			}
			if tt.wantDesc != "" || tt.wantSummary != "" || tt.wantDeprecated != nil {
				first := got[0].(*M.ClassMethod)
				if first.Description != tt.wantDesc {
					t.Errorf("merged description = %q, want %q", first.Description, tt.wantDesc)
				}
				if first.Summary != tt.wantSummary {
					t.Errorf("merged summary = %q, want %q", first.Summary, tt.wantSummary)
				}
				if first.Deprecated != tt.wantDeprecated {
					t.Errorf("merged deprecated = %v, want %v", first.Deprecated, tt.wantDeprecated)
				}
			}
			if tt.wantSameSlice {
				if &got[0] != &tt.members[0] {
					t.Errorf("dropOverloadImplementationMethods() returned a new slice, want the input slice unchanged (no-drop fast path)")
				}
			}
		})
	}
}

// declNames renders a declaration slice as a comma-separated name list for
// failure messages.
func declNames(decls []M.Declaration) string {
	if len(decls) == 0 {
		return "<empty>"
	}
	names := make([]string, len(decls))
	for i, d := range decls {
		names[i] = d.Name()
	}
	return strings.Join(names, ", ")
}

// memberNames renders a class member slice as a comma-separated name list
// for failure messages.
func memberNames(members []M.ClassMember) string {
	if len(members) == 0 {
		return "<empty>"
	}
	names := make([]string, len(members))
	for i, m := range members {
		names[i] = memberName(m)
	}
	return strings.Join(names, ", ")
}

// memberName renders one class member's name for failure messages; the
// ClassMember interface has no name accessor, so the concrete types are
// read directly.
func memberName(m M.ClassMember) string {
	switch v := m.(type) {
	case *M.ClassMethod:
		return v.Name
	case *M.CustomElementField:
		return v.Name
	default:
		return fmt.Sprintf("%T", m)
	}
}
