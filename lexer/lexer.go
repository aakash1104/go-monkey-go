package lexer

import (
    "go-monkey-go/token"
)

type Lexer struct {
    input string
    pos int // current position scanning
    read_pos int // current reading position (after current char)
    ch byte // current character we're reading
}

func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.ReadChar()
    return l
}

/**
 * @func ReadChar: Helper function to read character at current position
 */
func (l *Lexer) ReadChar() {
    if (l.read_pos >= len(l.input)) {
        l.ch = 0;
    } else {
        l.ch = l.input[l.read_pos]
    }
    l.pos = l.read_pos
    l.read_pos += 1
}

/**
 * @func NextToken: Returns a token.Token struct
 * @receiver lexer
 */
func (l *Lexer) NextToken() token.Token {
    var tok token.Token
    
    switch l.ch {
    case '=':
        tok = NewToken(token.ASSIGN, l.ch)
    case ';':
        tok = NewToken(token.SEMICOLON, l.ch)
    case '(':
        tok = NewToken(token.LPAREN, l.ch)
    case ')':
        tok = NewToken(token.RPAREN, l.ch)
    case '{':
        tok = NewToken(token.LBRACE, l.ch)
    case '}':
        tok = NewToken(token.RBRACE, l.ch)
    case '[':
        tok = NewToken(token.LSQUARE, l.ch)
    case ']':
        tok = NewToken(token.RSQUARE, l.ch)
    case ',':
        tok = NewToken(token.COMMA, l.ch)
    case '+':
        tok = NewToken(token.PLUS, l.ch)
    case '-':
        tok = NewToken(token.MINUS, l.ch)
    case '*':
        tok = NewToken(token.MULTIPLY, l.ch)
    case '/':
        tok = NewToken(token.DIVIDE, l.ch)
    case '%':
        tok = NewToken(token.MODULO, l.ch)
    case 0:
        tok.Literal = ""
        tok.Type = token.EOF
    }
    // Read next character
    l.ReadChar()
    return tok
}

func NewToken(token_type token.TokenType, ch byte) token.Token {
    return token.Token{Type: token_type, Literal: string(ch)}
}

