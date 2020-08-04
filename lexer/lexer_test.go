package lexer

import (
    "testing"
    "go-monkey-go/token"
)

func TestSimpleToken(t *testing.T) {
    input := `=`
    l := New(input)
    tok := l.NextToken()

     if (tok.Type != token.ASSIGN) {
        t.Fatalf("SimpleTokenTest failed - Expected %s and got %s",
                    token.ASSIGN, tok.Type)
    }
}

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

func TestTokenAdvancedSourceCode(t *testing.T) {
    input := `let f = 5;
let t = 10;
let add = fn(x, y) {
    x + y;
};

let result = add(f, t);

10 < 5 > 10!;

if (5 < 10) {
    return true;
} 
else
    return false;`

    tests := []struct {
        ExpectedType token.TokenType
        ExpectedLiteral string
    } {
        {token.LET, "let"},
        {token.IDENT, "f"},
        {token.ASSIGN, "="},
        {token.INT, "5"},
        {token.SEMICOLON, ";"},
        {token.LET, "let"},
        {token.IDENT, "t"},
        {token.ASSIGN, "="},
        {token.INT, "10"},
        {token.SEMICOLON, ";"},
        {token.LET, "let"},
        {token.IDENT, "add"},
        {token.ASSIGN, "="},
        {token.FUNCTION, "fn"},
        {token.LPAREN, "("},
        {token.IDENT, "x"},
        {token.COMMA, ","},
        {token.IDENT, "y"},
        {token.RPAREN, ")"},
        {token.LBRACE, "{"},
        {token.IDENT, "x"},
        {token.PLUS, "+"},
        {token.IDENT, "y"},
        {token.SEMICOLON, ";"},
        {token.RBRACE, "}"},
        {token.SEMICOLON, ";"},
        {token.LET, "let"},
        {token.IDENT, "result"},
        {token.ASSIGN, "="},
        {token.IDENT, "add"},
        {token.LPAREN, "("},
        {token.IDENT, "f"},
        {token.COMMA, ","},
        {token.IDENT, "t"},
        {token.RPAREN, ")"},
        {token.SEMICOLON, ";"},
        {token.INT, "10"},
        {token.LT, "<"},
        {token.INT, "5"},
        {token.GT, ">"},
        {token.INT, "10"},
        {token.NOT, "!"},
        {token.SEMICOLON, ";"},
        {token.IF, "if"},
        {token.LPAREN, "("},
        {token.INT, "5"},
        {token.LT, "<"},
        {token.INT, "10"},
        {token.RPAREN, ")"},
        {token.LBRACE, "{"},
        {token.RETURN, "return"},
        {token.TRUE, "true"},
        {token.SEMICOLON, ";"},
        {token.RBRACE, "}"},
        {token.ELSE, "else"},
        {token.RETURN, "return"},
        {token.FALSE, "false"},
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
