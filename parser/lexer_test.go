package parser

import (
	"testing"
)

func TestRecogniseToken(t *testing.T) {
	useCases := []Token{
		{
			LPARENSTR, "[",
		},
		{
			SUM, "SUM",
		},
		{
			LPARENSTR, "[",
		},
		{
			INT, "INT",
		},
		{
			FILENAME, "a.txt",
		},
		{
			FILENAME, "b.txt",
		},
		{
			RPARENSTR, "]",
		},
		{
			FILENAME, "c.txt",
		},
		{
			RPARENSTR, "]",
		},
	}

	lex := NewLexer("[SUM [INT a.txt b.txt] c.txt]")

	for _, useCase := range useCases {
		tok := lex.RecogniseToken()
		if tok.Type != useCase.Type {
			t.Errorf("Incorrect token type. Expected: %q Actual: %q",
				useCase.Type, tok.Type)
		}

		if tok.Literal != useCase.Literal {
			t.Errorf("Incorrect literal. Expected: %q Actual: %q",
				useCase.Literal, tok.Literal)
		}
	}
}
