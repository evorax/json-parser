package json

// lexer
type Token struct {
	Type    Type
	Literal string
}

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// parser
type Parser struct {
	lexer     *Lexer
	curToken  Token
	peekToken Token
}
