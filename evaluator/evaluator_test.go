package evaluator

import (
	"testing"

	"github.com/aitumik/interpreter/lexer"
	"github.com/aitumik/interpreter/object"
	"github.com/aitumik/interpreter/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	prog := p.ParseProgram()

	return Eval(prog)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	panic("implement me")
}
