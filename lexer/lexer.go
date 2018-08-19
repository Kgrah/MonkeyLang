package lexer

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
	return l
}
