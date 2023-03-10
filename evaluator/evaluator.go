package evaluator

import (
	"github.com/aitumik/interpreter/ast"
	"github.com/aitumik/interpreter/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	}

	return nil
}
