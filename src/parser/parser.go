package parser

import (
  "fmt"
)

var tree = make(map[int]*Valuable)
var all = make(map[string]*Valuable)

type Parser struct {
  lexer       *Lexer
  curToken    Token
  peekToken   Token
  Errors      []ParseError
}

type ParseError struct {
  expected    TokenType
  got         TokenType
  line        int
  column      int
}
func (pe *ParseError) message() string {
  return fmt.Sprintf("Parse error, expected '%s', got '%s'", pe.expected, pe.got)
}

func NewParser(lexer *Lexer) *Parser {
  parser := &Parser{
    lexer: lexer,
  }

  parser.nextToken()

  return parser
}

func (p *Parser) Parse() *Valuable {
  p.parseValuable()
  p.parseValuable()
  p.parseValuable()
  return tree[0]
}

func (p *Parser) parseValuable() {
  level := p.findLevel()

  if !p.expectPeek(ITEM) {
    return
  }

  itemToken := p.curToken

  if !p.expectPeek(IDENT) {
    return
  }

  nameToken := p.curToken
  name := p.curToken.Literal

  var valuable *Valuable
  if all[name] != nil {
    valuable = all[name]
  } else {
    valuable = &Valuable {Token: itemToken}
    valuable.Name = &Name{Token: nameToken, Value: name}
    valuable.Level = &Level{Value: level}
    valuable.Children = []*Valuable{}
    p.expectSkip(SEMICOLON)
    if p.expectPeek(STRING) {
      valuable.Description = &Description{Token: p.curToken, Value: p.curToken.Literal}
    }
    all[name] = valuable
    if level >= 1 {
      tree[level - 1].Children = append(tree[level - 1].Children, valuable)
    }
    tree[level] = valuable
  }
}

func (p *Parser) findLevel() int {
  level := 0
  for p.expectPeek(LEVEL) {
    level += 1
  }
  return level
}

func (p *Parser) nextToken() {
  p.curToken = p.peekToken
  p.peekToken = p.lexer.NextToken()
}

func (p *Parser) expectSkip(t TokenType) {
  if p.peekTokenIs(t) {
    p.nextToken()
  }
}


func (p *Parser) expectPeek(t TokenType) bool {
  if p.peekTokenIs(t) {
    p.nextToken()
    return true
  } else {
    p.peekError(t)
    return false
  }
}

func (p *Parser) peekTokenIs(t TokenType) bool {
  return p.peekToken.Type == t
}

func (p *Parser) curTokenIs(t TokenType) bool {
  return p.curToken.Type == t
}

func (p *Parser) peekError(t TokenType) {
  parseError := ParseError{expected: t, got: p.peekToken.Type}
  p.Errors = append(p.Errors, parseError)
}
