package language

import (
	"fmt"
)

var currentCategory *Category = nil
var currentItem *Item = nil
var currentLevel = 0
var firstLevelItems []*Item

type Parser struct {
	lexer     *Lexer
	curToken  Token
	peekToken Token
	Errors    []ParseError
}

type ParseError struct {
	expected TokenType
	got      TokenType
	line     int
	column   int
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

func (p *Parser) Parse() []*Item {
	for !p.peekTokenIs(EOF) {
		p.parseCategory()
	}
	return firstLevelItems
}

func (p *Parser) parseCategory() {
	if !p.expectPeek(IDENT) {
		return
	}

	currentCategory = &Category{Token: p.curToken, Value: p.curToken.Literal}
	p.parseItem()
}

func (p *Parser) parseItem() {
	var level = p.findLevel()
	if !p.expectPeek(DASH) {
		return
	}

	var item = &Item{Token: p.curToken, Category: currentCategory, Children: []*Item{}}
	if level == 0 {
		firstLevelItems = append(firstLevelItems, item)
	}
	if level > currentLevel {
		currentItem.Children = append(currentItem.Children, item)
	}
	currentItem = item
	currentLevel = level

	p.parseMarker()
	p.parseName()
	p.parseDescription()
	p.parseAnnotation()

	p.expectSkip(RSB)
	p.expectSkip(RP)
	p.expectSkip(SEMICOLON)

	if p.peekTokenIs(LEVEL) || p.peekTokenIs(DASH) {
		p.parseItem()
	}
}

func (p *Parser) parseMarker() {
	if p.peekTokenIs(LP) {
		currentItem.Marked = true
		p.nextToken()
	}
}

func (p *Parser) parseName() {
	if !p.expectPeek(IDENT) {
		return
	}
	currentItem.Name = &Name{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseDescription() {
	if !p.expectPeek(STRING) {
		return
	}
	currentItem.Description = &Description{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseAnnotation() {
	if !p.expectPeek(LSB) {
		return
	}
	if !p.expectPeek(IDENT) {
		return
	}
	if p.peekTokenIs(RSB) {
		unary := &UnaryAnnotation{Token: p.curToken, Name: &Name{Token: p.curToken, Value: p.curToken.Literal}}
		currentItem.Annotations = append(currentItem.Annotations, unary)
	} else if p.peekTokenIs(DASH) {
		left := &Name{Token: p.curToken, Value: p.curToken.Literal}
		p.expectPeek(DASH)
		operator := Operator{Token: p.curToken}
		if !p.expectPeek(IDENT) {
			return
		}
		right := &Name{Token: p.curToken, Value: p.curToken.Literal}

		binary := &BinaryAnnotation{
			Token:    p.curToken,
			Left:     left,
			Operator: &operator,
			Right:    right}
		currentItem.Annotations = append(currentItem.Annotations, binary)
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
