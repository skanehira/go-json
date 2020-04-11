package parser

import (
	"strconv"

	"github.com/skanehira/go-json/ast"
	"github.com/skanehira/go-json/token"
)

type Parser struct {
	input        []byte
	position     int
	readPosition int
	ch           byte
}

func (p *Parser) skipWhitespace() {
	for p.ch == ' ' || p.ch == '\t' || p.ch == '\n' || p.ch == '\r' {
		p.readChar()
	}
}

func (p *Parser) peekChar() byte {
	if p.readPosition >= len(p.input) {
		return 0
	}
	return p.input[p.readPosition]
}

func (p *Parser) readChar() {
	if p.readPosition >= len(p.input) {
		p.ch = 0
	} else {
		p.skipWhitespace()
		p.ch = p.input[p.readPosition]
	}

	p.position = p.readPosition
	p.readPosition += 1
}

func (p *Parser) curCharIs(ch byte) bool {
	return p.ch == ch
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (p *Parser) Parse() ast.Value {
	p.readChar()
	var value ast.Value

	for !p.curCharIs(0) {
		switch p.ch {
		case '"':
			v := p.ParseString()
			if p.peekChar() == 0 {
				value = v
				break
			}
		case '{':
		case '[':

		default:
			if isDigit(p.ch) {
				v := p.ParseInt()
				if p.peekChar() == 0 {
					value = v
					break
				}
			} else {

			}
		}
		p.readChar()
	}

	return value
}

func (p *Parser) ParseInt() ast.Int {
	i := ast.Int{
		Type: token.Int,
	}

	b := []byte{}

	println(string(p.ch))

	for !p.curCharIs(0) && isDigit(p.ch) {
		b = append(b, p.ch)
		p.readChar()
	}

	v, _ := strconv.Atoi(string(b))
	i.Value = v

	return i
}

func (p *Parser) ParseString() ast.String {
	str := ast.String{
		Type: token.String,
	}

	value := []byte{}

	p.readChar()

	for !p.curCharIs(0) && !p.curCharIs(byte('"')) {
		value = append(value, p.ch)
		p.readChar()
	}

	str.Value = string(value)
	return str
}
