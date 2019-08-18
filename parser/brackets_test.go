package parser

import (
	"testing"
)

func TestBrackets(t *testing.T) {
	var useCases = []struct {
		expr   string
		expect bool
	}{
		{
			expr:   "[]",
			expect: true,
		},
		{
			expr:   "][",
			expect: false,
		},
		{
			expr:   "]]",
			expect: false,
		},
		{
			expr:   "[[",
			expect: false,
		},
		{
			expr:   "[[]]",
			expect: true,
		},
		{
			expr:   "[ SUM [ DIF a.txt b.txt c.txt ] [ INT b.txt c.txt ] ]",
			expect: true,
		},
	}

	for _, useCase := range useCases {
		result, err := CheckBrackets(useCase.expr)
		if err != nil {
			t.Error(err)
		}
		if result != useCase.expect {
			t.Errorf("Expression: %q Excpect: %v", useCase.expr, useCase.expect)
		}
	}
}
