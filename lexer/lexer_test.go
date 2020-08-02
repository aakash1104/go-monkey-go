package lexer

import (
    "testing"
    "go-monkey-go/token"
)

func TestNextToken(t *testing.T) {
    input := `=+-*/%(){}[],;`

    // Array of structs each has a test case
    tests := []struct {
        ExpectedType token.TokenType
        ExpectedLiteral string
    } {
        {token.ASSIGN, "="},
        {token.PLUS, "+"},
        {token.MINUS, "-"},
        {token.MULTIPLY, "*"},
        {token.DIVIDE, "/"},
        {token.MODULO, "%"},
        {token.LPAREN, "("},
        {token.RPAREN, ")"},
        {token.LBRACE, "{"},
        {token.RBRACE, "}"},
        {token.LSQUARE, "["},
        {token.RSQUARE, "]"},
        {token.COMMA, ","},
        {token.SEMICOLON, ";"},
        {token.EOF, ""},
    }

    l := New(input)
    for i, tt := range tests {
        tok := l.NextToken()

        if (tok.Type != tt.ExpectedType) {
            t.Fatalf("test[%d] - incorrect token type. Expected %s and got %s", 
                        i, tt.ExpectedType, tok.Type)
        }
        
        if (tok.Literal != tt.ExpectedLiteral) {
            t.Fatalf("test[%d] - incorrect literal. Expected %s, got %s", i, 
                        tt.ExpectedLiteral, tok.Literal)
        }
    }
}
