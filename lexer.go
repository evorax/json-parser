package json

import (
	"unicode"
)

type Type string

const (
	LEFT_BRACE    Type = "{"
	RIGHT_BRACE   Type = "}"
	LEFT_BRACKET  Type = "["
	RIGHT_BRACKET Type = "]"
	COLON         Type = ":"
	COMMA         Type = ","
	STRING        Type = "STRING"
	NUMBER        Type = "NUMBER"
	TRUE          Type = "true"
	FALSE         Type = "false"
	NULL          Type = "null"
	EOF           Type = "EOF"
)

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

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skip()

	switch l.ch {
	case '{':
		tok = set(LEFT_BRACE, l.ch)
	case '}':
		tok = set(RIGHT_BRACE, l.ch)
	case '[':
		tok = set(LEFT_BRACKET, l.ch)
	case ']':
		tok = set(RIGHT_BRACKET, l.ch)
	case ':':
		tok = set(COLON, l.ch)
	case ',':
		tok = set(COMMA, l.ch)
	case '"':
		tok.Type = STRING
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if isNum(l.ch) {
			tok.Type = NUMBER
			tok.Literal = l.readNumber()
			return tok
		} else if isLetter(l.ch) {
			literal := l.readIdentifier()
			tok.Type = lookupIdent(literal)
			tok.Literal = literal
			return tok
		} else {
			tok = set(EOF, l.ch)
		}
	}

	l.readChar()
	return tok
}

func set(Type Type, ch byte) Token {
	return Token{Type: Type, Literal: string(ch)}
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isNum(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skip() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isNum(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch))
}

func lookupIdent(ident string) Type {
	switch ident {
	case "true":
		return TRUE
	case "false":
		return FALSE
	case "null":
		return NULL
	default:
		return STRING
	}
}
