package parser

import (
  "fmt"
)

var tree = make(map[int]*Item)
var all = make(map[string]*Item)

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

func (p *Parser) Parse() *Item {
  for !p.peekTokenIs(EOF) {
    p.parseValuable()
  }
  return tree[0]
}

func (p *Parser) parseValuable() {
  level := p.findLevel()

  if !p.expectPeek(IDENT) {
    return
  }

  nameToken := p.curToken
  name := p.curToken.Literal

  var item *Item
  if all[name] != nil {
    item = all[name]
  } else {
    item = &Item {Token: nameToken}
    item.Name = &Name{Token: nameToken, Value: name}
    item.Children = []*Item{}
    p.expectSkip(SEMICOLON)
    if p.expectPeek(STRING) {
      item.Description = &Description{Token: p.curToken, Value: p.curToken.Literal}
    }
    item.Level = level
    all[name] = item
    if level >= 1 {
      tree[level - 1].Children = append(tree[level - 1].Children, item)
    }
    tree[level] = item
  }
}

func (p *Parser) findLevel() int {
  level := 0
  for p.expectPeek(LEVEL) || p.expectPeek(ITEM) {
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
