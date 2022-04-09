package ast

import (
	"bytes"
	"monkeylang/src/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	StatementNode()
}

type Expression interface {
	Node
	ExpressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type LetStatement struct {
	Token token.Token
	Name *Identifier
	Value Expression
}

func (l *LetStatement) StatementNode() {}
func (l *LetStatement) TokenLiteral() string { return l.Token.Literal }
func (l *LetStatement) String() string {
	var out bytes.Buffer
	
	out.WriteString(l.TokenLiteral() + " ")
	out.WriteString(l.Name.String())
	out.WriteString(" = ")

	if l.Value != nil {
		out.WriteString(l.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

// func (ls *LetStatement) String() string {
// 	return "let"
// }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) ExpressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string {
	return i.Value
}

type ReturnStatement struct {
	Token token.Token
	ReturnValue Expression
}

func (l *ReturnStatement) StatementNode() {}
func (l *ReturnStatement) TokenLiteral() string { return l.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
  
	out.WriteString(rs.TokenLiteral() + " ")
	
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
  
	out.WriteString(";")
	return out.String()
}

type ExpressionStatement struct {
	Token token.Token
	Expression Expression
}

func (es *ExpressionStatement) StatementNode() {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}