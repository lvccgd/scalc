package parser

import (
	"testing"
)

func TestParser(t *testing.T) {
	useCase := "[ SUM [ DIF a.txt b.txt c.txt ] [ INT b.txt c.txt ] ]"
	lex := NewLexer(useCase)
	parser := NewParser(lex)
	expr := parser.Parse()
	if useCase != expr.String() {
		t.Errorf("Expexted: %q Actual: %q", useCase, expr.String())
	}
}
