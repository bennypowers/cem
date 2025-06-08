package generate

import (
	"errors"
	"fmt"

	M "bennypowers.dev/cem/manifest"
	ts "github.com/tree-sitter/go-tree-sitter"
)

func generateClassDeclaration(
	captures CaptureMap,
	root *ts.Node,
	code []byte,
) (err error, declaration *M.ClassDeclaration) {
	declaration = &M.ClassDeclaration{ Kind: "class" }

	className, ok := captures["class.name"]
	if (!ok || len(className) <= 0) {
		str := string(code)
		fmt.Println(str)
		return errors.Join(err, &NoCaptureError{ "class.name", "classDeclaration" }), nil
	}

	declaration.ClassLike.Name = className[0].Text

	error, members := getClassMembersFromClassDeclarationNode(code, root)
	if error != nil {
		return errors.Join(err, error), nil
	}

	for _, method := range members {
		declaration.Members = append(declaration.Members, method)
	}

	jsdoc, ok := captures["jsdoc"]
	if (ok && len(jsdoc) > 0) {
		error, info := NewClassInfo(jsdoc[0].Text)
		if error != nil {
			return errors.Join(err, error), nil
		} else {
			info.MergeToClassDeclaration(declaration)
		}
	}

	return nil, declaration
}
