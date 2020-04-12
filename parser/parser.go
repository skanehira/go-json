package parser

import (
	"strconv"
	"strings"

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

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (p *Parser) Parse() ast.Value {
	p.readChar()
	var value ast.Value

	for !p.curCharIs(0) {
		if p.curCharIs('"') {
			value = p.ParseString()
			break
		} else if p.curCharIs('{') {
			value = p.ParseObject()
			break
		} else if isDigit(p.ch) {
			value = p.ParseInt()
			break
		}
		p.readChar()
	}

	return value
}

func (p *Parser) ParseObject() ast.Object {
	o := ast.Object{
		Type: token.Object,
	}

	elements := []ast.Element{}
	for !p.curCharIs(0) && !p.curCharIs(byte('}')) {
		p.readChar()

		name := strings.Trim(p.ParseString().String(), "\"")
		e := ast.Element{Name: name, Value: p.Parse()}
		elements = append(elements, e)
	}

	o.Elements = elements

	return o
}

func (p *Parser) ParseInt() ast.Int {
	i := ast.Int{
		Type: token.Int,
	}

	pos := p.position

	for !p.curCharIs(0) && isDigit(p.ch) {
		p.readChar()
	}

	v, _ := strconv.Atoi(string(p.input[pos:p.position]))
	i.Value = v

	return i
}

func (p *Parser) ParseString() ast.String {
	str := ast.String{
		Type: token.String,
	}

	p.readChar()
	pos := p.position

	for !p.curCharIs(0) && !p.curCharIs(byte('"')) {
		p.readChar()
	}

	str.Value = string(p.input[pos:p.position])

	p.readChar()
	return str
}
