package parser

import (
	"github.com/skanehira/go-json/ast"
)

type Parser struct {
	values   []ast.Value
	input    []byte
	curPos   int
	peekPos  int
	curToken byte
}

func NewParser(input []byte) *Parser {
	return &Parser{
		input: input,
	}
}

func (p *Parser) Init() *Parser {
	p.Read()
	return p
}

func (p *Parser) Read() {
	if p.peekPos >= len(p.input) {
		p.curToken = 0
		return
	}

	// skip space and tab
	for token := p.input[p.peekPos]; token != 0 &&
		(token == 9 || token == 32); {

		p.curPos = p.peekPos
		p.peekPos++
		token = p.input[p.peekPos]
	}

	p.curToken = p.input[p.peekPos]
	p.curPos = p.peekPos
	p.peekPos++
}

func (p *Parser) CurToken() byte {
	return p.curToken
}

func (p *Parser) CurTokenIs(b byte) bool {
	return p.CurToken() == b
}

func (p *Parser) PeekToken() byte {
	if p.peekPos >= len(p.input) {
		return 0
	}
	return p.input[p.peekPos]
}

func (p *Parser) PeekTokenIs(b byte) bool {
	return p.PeekToken() == b
}

func (p *Parser) ParseString() (ast.Value, error) {
	v := ast.String{
		ValueType: ast.StringType,
	}

	if !p.CurTokenIs('"') {
		return v, ErrInvalidNextToken
	}

	p.Read()

	pos := p.curPos

	for !p.CurTokenIs('"') {
		if p.CurTokenIs(0) {
			return v, ErrInvalidNextToken
		}
		p.Read()
	}

	v.Value = string(p.input[pos:p.curPos])

	return v, nil
}
