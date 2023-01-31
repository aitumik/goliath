package token

const (
	ILLEGAL = "ILLEGAL"

	// operators
	PLUS   = "PLUS"
	MINUS  = "MINUS"
	ASSIGN = "="

	IDENT = "IDENT"
	INT   = "INT"

	// keywords
	LET = "LET"
	FUN = "FUN"

	LPAREN = "("
	RPAREN = "("

	LBRACE = "{"
	RBRACE = "}"

	// delimeters
	SEMICOLON = ";"
	COMMA     = ","

	// !-/*5;

	BANG     = "!"
	SLASH    = "/"
	ASTERISK = "*"

	LT = "<"
	GT = ">"

	IF   = "IF"
	ELSE = "ELSE"

	TRUE  = "TRUE"
	FALSE = "FALSE"

	RETURN = "RETURN"

	EQ  = "=="
	NEQ = "!="

	EOF = "EOF"
)

var keywords = map[string]TokenType{
	"fn":     FUN,
	"let":    LET,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdent(ident string) TokenType {
	if value, ok := keywords[ident]; ok {
		return value
	}

	return IDENT
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
