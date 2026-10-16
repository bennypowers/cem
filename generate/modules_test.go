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
	"strings"
	"testing"

	M "bennypowers.dev/cem/manifest"
)

// Inline assertions: dropOverloadImplementationSignatures is a Tier-1 pure
// function, which docs/TESTING.md requires be covered by exhaustive
// table-driven in-package tests rather than fixture/golden files (same
// rationale as TestParameterFromNode). End-to-end overload dropping is
// covered by the overload golden fixtures.

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
