package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

//The book has this function in the lexer.go file

//NewToken ...Constructor
func NewToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	//Operators
	ASSIGN = "="
	PLUS   = "+"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
