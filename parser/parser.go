package parser

import (
	"log"
	"strconv"

	"github.com/skanehira/go-json/ast"
	"github.com/skanehira/go-json/lexer"
	"github.com/skanehira/go-json/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) Parse() ast.Value {
	var v ast.Value
	switch p.curToken.Type {
	case token.LBRACE: // object
		v = p.parseObject()
	case token.LBRACKET: // array
		v = p.parseArray()
	case token.LETTER: // string
		v = p.parseString()
	case token.INT:
		v = p.parseInt()
	case token.TRUE, token.FALSE:
		v = p.parseBool()
	case token.NULL:
		v = ast.Null{}
	default:
		log.Panicf("wrong current token: %s", p.curToken.Type)
	}

	p.nextToken()
	return v
}

func (p *Parser) parseString() ast.String {
	return ast.String{
		Value: p.curToken.Literal,
	}
}

func (p *Parser) parseBool() ast.Bool {
	b, err := strconv.ParseBool(p.curToken.Literal)
	if err != nil {
		log.Panicf("wront bool literal: %s", p.curToken.Literal)
	}
	return ast.Bool{
		Value: b,
	}
}

func (p *Parser) parseInt() ast.Int {
	i, err := strconv.Atoi(p.curToken.Literal)
	if err != nil {
		log.Panicf("wrong int literal: %s", p.curToken.Literal)
	}
	return ast.Int{
		Value: int64(i),
	}
}

func (p *Parser) parseArray() ast.Array {
	array := ast.Array{
		Values: []ast.Value{},
	}

	for p.curToken.Type != token.RBRACKET {
		if p.peekTokenIs(token.RBRACKET) {
			break
		}
		p.nextToken()
		array.Values = append(array.Values, p.Parse())
	}

	return array
}

func (p *Parser) parseObject() ast.Object {
	obj := ast.Object{
		Elements: []ast.Element{},
	}

	for p.curToken.Type != token.RBRACE {
		if p.peekTokenIs(token.RBRACE) {
			break
		}
		if !p.peekTokenIs(token.LETTER) {
			log.Panicf("peek token wrong: %s", p.peekToken.Type)
		}

		p.nextToken()

		key := p.curToken.Literal

		if !p.peekTokenIs(token.COLON) {
			log.Panicf("peek token wrong: %s", p.peekToken.Type)
		}

		p.nextToken()
		p.nextToken()

		el := ast.Element{
			Key:   key,
			Value: p.Parse(),
		}

		obj.Elements = append(obj.Elements, el)
	}

	return obj
}

type JSON struct {
	Value ast.Value
}

func (j *JSON) String() string {
	return j.Value.String()
}
