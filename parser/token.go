package parser

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	LPAREN              = '['
	RPAREN              = ']'
	LPARENSTR TokenType = "["
	RPARENSTR           = "]"
	SUM                 = "SUM"
	DIF                 = "DIF"
	INT                 = "INT"
	FILENAME            = "FILENAME"
	ILLEGAL             = "ILLEGAL"
	EOF                 = "EOF"
)
