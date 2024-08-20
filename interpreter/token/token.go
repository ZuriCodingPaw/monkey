package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals

	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 12468532

	// operators

	ASSIGN = "ASSIGN"
	PLUS   = "+"
	MINUS  = "-"
	BANG   = "!"
	STAR   = "*"
	SLASH  = "/"
	LT     = "<"
	GT     = ">"

	EQ     = "=="
	NOT_EQ = "!="
	// delimiters

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords

	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdent(ident string) TokenType {
	if keyword, ok := keywords[ident]; ok {
		return keyword
	}
	return IDENT
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
