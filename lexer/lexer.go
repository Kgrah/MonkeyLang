package lexer

import "github.com/kgrah/MonkeyLang/token"

//The Lexer type
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

//NewLexer ...Constructor
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = token.NewToken(token.ASSIGN, l.ch)
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
	case ',':
		tok = NewToken(token.COMMA, l.ch)
	case '+':
		toke = NewToken(token.PLUS, l.ch)
	}
}
