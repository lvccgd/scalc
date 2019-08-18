package parser

import (
	"regexp/syntax"
)

// Lexer represent current lexical analysis state
type Lexer struct {
	expr    string
	pos     int  // position at current symbol
	readPos int  // position after current symbol
	ch      rune // current symbol
}

// NewLexer returns new lexer with task to read
func NewLexer(str string) *Lexer {
	lex := &Lexer{expr: str}
	lex.readChar()
	return lex
}

func (lex *Lexer) readChar() {
	if lex.readPos >= len(lex.expr) {
		lex.ch = 0
	} else {
		lex.ch = rune(lex.expr[lex.readPos])
	}
	lex.pos = lex.readPos
	lex.readPos++
}

// RecogniseToken sequential read tokens
func (lex *Lexer) RecogniseToken() Token {
	var tok Token
	lex.ignoreWhitespaces()

	switch lex.ch {
	case '[':
		tok = newToken(LPARENSTR, lex.ch)
	case ']':
		tok = newToken(RPARENSTR, lex.ch)
	case 0:
		tok.Type = EOF
		tok.Literal = ""
	default:
		if isValidSymbol(lex.ch) {
			word := lex.readWord()
			if isOperator(word) {
				tok.Type = TokenType(word)
				tok.Literal = word
			} else {
				tok.Type = FILENAME
				tok.Literal = word
			}
			return tok
		}
		tok = newToken(ILLEGAL, lex.ch)
	}

	lex.readChar()
	return tok
}

func isOperator(word string) bool {
	return word == SUM || word == DIF || word == INT
}

func isValidSymbol(ch rune) bool {
	return syntax.IsWordChar(rune(ch)) || ch == '.' || ch == '-'
}

func newToken(tokenType TokenType, ch rune) Token {
	return Token{tokenType, string(ch)}
}

func (lex *Lexer) readWord() string {
	pos := lex.pos
	for isValidSymbol(lex.ch) {
		lex.readChar()
	}
	return lex.expr[pos:lex.pos]
}

func (lex *Lexer) ignoreWhitespaces() {
	for lex.ch == ' ' || lex.ch == '\t' || lex.ch == '\n' || lex.ch == '\r' {
		lex.readChar()
	}
}
