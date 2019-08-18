package parser

import (
	"bytes"
	"strings"
)

type INode interface {
	String() string
}

type IExpression interface {
	INode
	expressionNode()
}

type Set struct {
	Token Token
}

func (set Set) expressionNode() {}

func (set Set) String() string {
	return set.Token.Literal
}

type Expression struct {
	Token    Token
	Operator string
	Sets     []IExpression
}

func (expr Expression) expressionNode() {}

func (expr Expression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, arg := range expr.Sets {
		args = append(args, arg.String())
	}
	out.WriteString("[ ")
	out.WriteString(expr.Operator + " ")
	out.WriteString(strings.Join(args, " "))
	out.WriteString(" ]")

	return out.String()
}
