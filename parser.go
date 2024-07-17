package json

import (
	"reflect"
	"strconv"
	"strings"
)

func NewParser(l *Lexer) *Parser {
	p := &Parser{lexer: l}
	p.next()
	p.next()
	return p
}

func (p *Parser) next() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) parseValue(v reflect.Value) {
	switch p.curToken.Type {
	case LEFT_BRACE:
		p.parseObject(v)
	case LEFT_BRACKET:
		p.parseArray(v)
	case STRING:
		if v.IsValid() && v.CanSet() {
			v.SetString(p.curToken.Literal)
		}
		p.next()
	case NUMBER:
		num, _ := strconv.ParseFloat(p.curToken.Literal, 64)
		if v.IsValid() && v.CanSet() {
			if v.Kind() == reflect.Float64 {
				v.SetFloat(num)
			} else if v.Kind() == reflect.Int {
				v.SetInt(int64(num))
			}
		}
		p.next()
	case TRUE:
		if v.IsValid() && v.CanSet() {
			v.SetBool(true)
		}
		p.next()
	case FALSE:
		if v.IsValid() && v.CanSet() {
			v.SetBool(false)
		}
		p.next()
	case NULL:
		p.next()
	}
}

func (p *Parser) parseObject(v reflect.Value) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	p.next()

	for p.curToken.Type != RIGHT_BRACE {
		if p.curToken.Type == STRING {
			key := p.curToken.Literal
			p.next()
			if p.curToken.Type == COLON {
				p.next()
				field := find(v, key)
				if field.IsValid() {
					p.parseValue(field)
				} else {
					p.parseValue(reflect.ValueOf(nil))
				}
			}
		}
		if p.curToken.Type == COMMA {
			p.next()
		}
	}

	p.next()
}

func (p *Parser) parseArray(v reflect.Value) {
	p.next()

	for p.curToken.Type != RIGHT_BRACKET {
		elem := reflect.New(v.Type().Elem()).Elem()
		p.parseValue(elem)
		v.Set(reflect.Append(v, elem))
		if p.curToken.Type == COMMA {
			p.next()
		}
	}

	p.next()
}

func find(v reflect.Value, key string) reflect.Value {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag.Get("json")
		tagParts := strings.Split(tag, ",")
		if tagParts[0] == key {
			return v.Field(i)
		}
	}
	return reflect.Value{}
}

func ParseJSON(input string, output any) {
	lexer := NewLexer(input)
	parser := NewParser(lexer)
	parser.parseValue(reflect.ValueOf(output))
}
