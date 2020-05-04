package parser

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/skanehira/go-json/ast"
)

type Parser struct {
	value    ast.Value
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

func (p *Parser) SkipWhiteSpace() {
	for token := p.PeekToken(); token != 0 &&
		(token == 9 || token == 32); {

		p.curPos = p.peekPos
		p.peekPos++
		token = p.input[p.peekPos]
	}
}

func (p *Parser) Read() {
	p.SkipWhiteSpace()

	if p.peekPos >= len(p.input) {
		p.curToken = 0
		return
	}

	p.curToken = p.input[p.peekPos]
	p.curPos = p.peekPos
	p.peekPos++
}

func (p *Parser) CurToken() byte {
	return p.curToken
}

func (p *Parser) CurTokenString() string {
	if p.CurTokenIs(0) {
		return "null(0)"
	}
	return string(p.curToken)
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

func (p *Parser) PeekTokenString() string {
	if p.PeekTokenIs(0) {
		return "null(0)"
	}
	return string(p.PeekToken())
}

func (p *Parser) PeekTokenIs(b byte) bool {
	return p.PeekToken() == b
}

func (p *Parser) ParseString() (ast.Value, error) {
	v := ast.String{
		ValueType: ast.StringType,
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

	p.Read()

	return v, nil
}

func (p *Parser) ParseObject() (ast.Value, error) {
	v := ast.Object{
		ValueType: ast.ObjectType,
	}

	if p.PeekTokenIs('}') {
		return v, nil
	}

	if !p.PeekTokenIs('"') {
		return nil, errors.Wrap(ErrInvalidNextToken, "next token="+p.PeekTokenString())
	}

	// parse key
	p.Read()
	str, err := p.ParseString()
	if err != nil {
		return nil, err
	}

	v.Key = str.(ast.String).Value

	if !p.CurTokenIs(':') {
		return nil, errors.Wrap(ErrInvalidToken, "current token="+p.CurTokenString())
	}

	// parse value
	p.Read()
	value, err := p.Parse()
	if err != nil {
		return nil, err
	}

	v.Value = value

	return v, nil
}

func (p *Parser) ParseInteger() (ast.Value, error) {
	v := ast.Integer{
		ValueType: ast.IntType,
	}

	pos := p.curPos

	for 48 <= p.curToken && p.curToken <= 57 {
		p.Read()
	}

	value := p.input[pos : p.curPos+1]
	i, err := strconv.Atoi(string(value))
	if err != nil {
		return nil, fmt.Errorf("failed to convert \"%s\" to integer", value)
	}

	v.Value = int64(i)

	p.Read()

	return v, nil
}

func (p *Parser) Parse() (ast.Value, error) {
	switch p.CurToken() {
	case '"':
		return p.ParseString()
	case '{':
		return p.ParseObject()
	case '[':
	}

	return nil, errors.Wrap(ErrInvalidToken, "current token="+p.CurTokenString())
}
