package visitor

import (
	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/parser"
)

type AstWalker struct {
	Strings *[]string
}

func (w AstWalker) Enter(node ast.Node) ast.Visitor {
	switch node := node.(type) {
	case *ast.StringLiteral:
		*w.Strings = append(*w.Strings, node.Value)
	}
	return w
}
func (w AstWalker) Exit(node ast.Node) {}

func ExtractStrings(source []byte) ([]string, error) {
	prg, err := parser.ParseFile(nil, "", source, 1)
	if err != nil {
		return nil, err
	}

	strings := make([]string, 0)
	walker := AstWalker{Strings: &strings}
	for _, st := range prg.Body {
		ast.Walk(walker, ast.Node(st))
	}

	return strings, nil
}
