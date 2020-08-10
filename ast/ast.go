package ast

import (
    "go-monkey-go/token"
)

type Node interface {
    TokenLiteral() string
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

type Identifier struct {
    Token token.Token // token.IDENT
    Value string
}

type LetStatement struct {
    Token token.Token // token.LET
    Name *Identifier
    Value Expression
}
/**
 * Helper function for debugging and testing only 
 **/
func (p *Program) TokenLiteral() string {
    if (len(p.Statements) > 0) {
        return p.Statements[0].TokenLiteral();
    } else {
        return ""
    }
}

func (ls *LetStatement) StatementNode() {}
func (ls *LetStatement) TokenLiteral() string {return ls.Token.Literal}

func (i *Identifier) ExpressionNode() {}
func (i *Identifier) TokenLiteral() string {return i.Token.Literal}
