package parser

import (
	"testing"
)

func TestString(t *testing.T) {
	expr := &Expression{
		Token:    Token{Type: LPARENSTR, Literal: "["},
		Operator: "SUM",
		Sets: []IExpression{
			Expression{
				Token:    Token{Type: LPARENSTR, Literal: "["},
				Operator: "DIF",
				Sets: []IExpression{
					Set{
						Token: Token{Type: FILENAME, Literal: "a.txt"},
					},
					Set{
						Token: Token{Type: FILENAME, Literal: "b.txt"},
					},
					Set{
						Token: Token{Type: FILENAME, Literal: "c.txt"},
					},
				},
			},
			Expression{
				Token:    Token{Type: LPARENSTR, Literal: "["},
				Operator: "INT",
				Sets: []IExpression{
					Set{
						Token: Token{Type: FILENAME, Literal: "b.txt"},
					},
					Set{
						Token: Token{Type: FILENAME, Literal: "c.txt"},
					},
				},
			},
		},
	}

	if expr.String() != "[ SUM [ DIF a.txt b.txt c.txt ] [ INT b.txt c.txt ] ]" {
		t.Errorf("Incorrect expression representation. Result: %q", expr.String())
	}
}
