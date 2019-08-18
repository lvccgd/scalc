package parser

type parseFn func() IExpression

// Parser store structured expresseion
type Parser struct {
	lex       *Lexer
	curToken  Token
	nextToken Token
	parseFns  map[TokenType]parseFn
}

// NewParser returns correct parser
func NewParser(lex *Lexer) *Parser {
	parser := &Parser{
		lex:      lex,
		parseFns: make(map[TokenType]parseFn),
	}

	parser.addParseFn(LPARENSTR, parser.parseExpression)
	parser.addParseFn(FILENAME, parser.parseFilename)

	parser.getToken()

	return parser
}

// Parse returns the all expression structure
func (parser *Parser) Parse() IExpression {
	if parser.nextTokenIs(EOF) {
		return Expression{}
	}

	return parser.parseExpression()
}

// getToken iterates the available tokens with lexer
func (parser *Parser) getToken() {
	parser.curToken = parser.nextToken
	parser.nextToken = parser.lex.RecogniseToken()
}

// parseExpression returns expression structure
func (parser *Parser) parseExpression() IExpression {
	parser.getToken()
	expr := Expression{Token: parser.curToken, Operator: parser.nextToken.Literal}
	expr.Sets = parser.parseExpressions()
	return expr
}

// parseFilename returns set ()
func (parser *Parser) parseFilename() IExpression {
	set := Set{}
	set.Token = parser.nextToken
	return set
}

func (parser *Parser) parseExpressions() []IExpression {
	sets := []IExpression{}
	if parser.nextTokenIs(RPARENSTR) {
		parser.getToken()
		return sets
	}

	for {
		parser.getToken()
		if parser.nextTokenIs(RPARENSTR) || parser.nextTokenIs(EOF) {
			return sets
		}
		fn := parser.parseFns[parser.nextToken.Type]
		sets = append(sets, fn())
	}

	return sets
}

func (parser *Parser) nextTokenIs(tokenType TokenType) bool {
	return parser.nextToken.Type == tokenType
}

func (parser *Parser) addParseFn(tokenType TokenType, fn parseFn) {
	parser.parseFns[tokenType] = fn
}
