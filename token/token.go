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
)

const (
    // Keywords
    FUNCTION =  "FUNCTION"
    LET = "LET"
)

const (
    // Operators
    PLUS = "+"
    MINUS = "-"
    MULTIPLY = "*"
    DIVIDE = "/"
    MODULO = "%"
    ASSIGN = "="
)

const (
    // Delimiters
    COMMA = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"

    LBRACE = "{"
    RBRACE = "}"

    LSQUARE = "["
    RSQUARE = "]"
)

const (
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
