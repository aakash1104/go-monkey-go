package token

type TokenType string

type Token struct {
    Type TokenType
    Literal string
}

const (
    // Parser specific
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"

    // Keywords
    FUNCTION =  "FUNCTION"
    LET = "LET"

    // Operators
    PLUS = "+"
    MINUS = "-"
    MULTIPLY = "*"
    DIVIDE = "/"
    MODULO = "%"
    ASSIGN = "="

    // Delimiters
    COMMA = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"

    LBRACE = "{"
    RBRACE = "}"

    LSQUARE = "["
    RSQUARE = "]"

    // Identifiers and literal values
    IDENT = "IDENT"
    INT = "INT"
    STRING = "STRING"
)

var keywords = map[string]TokenType {
    "let": LET,
    "fn": FUNCTION,
}

func LookUpIdentifer(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    // Not found in keywords, so must be an identifier
    return IDENT
}
