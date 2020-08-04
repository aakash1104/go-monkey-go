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
 * @func NextToken: Returns the next token in the input 
 * @return: token.Token struct
 * @receiver lexer
 */
func (l *Lexer) NextToken() token.Token {
    var tok token.Token

    l.SkipWhitespace()

    switch l.ch {
    case '=':
        if (l.PeekChar() == '=') {
            ch := l.ch
            l.ReadChar()
            tok = token.Token{Type: token.EQUAL, Literal: string(ch) + string(l.ch)}
        } else {
            tok = NewToken(token.ASSIGN, l.ch)
        }
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
    case '<':
        if (l.PeekChar() == '=') {
            ch := l.ch
            l.ReadChar()
            tok = token.Token{Type: token.LTE, Literal: string(ch) + string(l.ch)}
        } else {
            tok = NewToken(token.LT, l.ch)
        }
    case '>':
        if (l.PeekChar() == '=') {
            ch := l.ch
            l.ReadChar()
            tok = token.Token{Type: token.GTE, Literal: string(ch) + string(l.ch)}
        } else {
            tok = NewToken(token.GT, l.ch)
        }
    case '!':
        if (l.PeekChar() == '=') {
            ch := l.ch
            l.ReadChar()
            tok = token.Token{Type: token.NEQ, Literal: string(ch) + string(l.ch)}
        } else {
            tok = NewToken(token.NOT, l.ch)
        }
    case 0:
        tok.Literal = ""
        tok.Type = token.EOF
    default:
        if (IsLetter(l.ch)) {
            // If letter, then read till we encounter a non-letter character
            tok.Literal = l.ReadIdentifier()
            tok.Type = token.LookUpIdentifer(tok.Literal)
            return tok
        } else if (IsDigit(l.ch)) {
            // Encountered a digit so read till we encounter a non-digit
            tok.Type = token.INT
            tok.Literal = l.ReadNumber()
            return tok
        } else {
            // Can't recognize character
            tok = NewToken(token.ILLEGAL, l.ch)
        }
    }
    // Read next character
    l.ReadChar()
    return tok
}

/**
 * @func ReadIdentifier: Scan input till a non-letter character is found
 * @receiver: Lexer
 * @return string corresponding to the identifier
 **/

func (l *Lexer) ReadIdentifier() string {
    orig_pos := l.pos
    for (IsLetter(l.ch)) {
        l.ReadChar()
    }
    return l.input[orig_pos:l.pos]
}

func IsLetter(ch byte) bool {
    return ('a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_')
}

func IsDigit (ch byte) bool {
    return '0' <= ch && ch <= '9'
}

/**
 * @func SkipWhitespace: Eat up whitespace.
 **/
func (l *Lexer) SkipWhitespace() {
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
        l.ReadChar()
    }
}

/**
 * @func NewToken: Function to return a new token struct with type @token_type
 * and literal @ch
 *
 **/
func NewToken(token_type token.TokenType, ch byte) token.Token {
    return token.Token{Type: token_type, Literal: string(ch)}
}

/**
 * @func: ReadNumber - read number till we encounter a non digit
 * @receiver: l *Lexer
 * @return string that contains the int
 **/
func (l *Lexer) ReadNumber() string {
    orig_pos := l.pos
    for IsDigit(l.ch) {
        l.ReadChar()
    }
    return l.input[orig_pos:l.pos]
}

func (l *Lexer) PeekChar() byte {
    if (l.read_pos < len(l.input)) {
        return l.input[l.read_pos]
    } else {
        return 0
    }
}
