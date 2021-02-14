package language

import (
	"fmt"
)

type Parser struct {
	lexer           *Lexer
	curToken        Token
	peekToken       Token
	Errors          []ParseError
	currentCategory *Category
	currentItem     *Item
	currentLevel    int
	firstLevelItems []*Item
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
		if !p.parseCategory() {
			for _, err := range p.Errors {
				fmt.Printf("Expected %s, got %s\n", err.expected, err.got)
			}
			break
		}
	}
	return p.firstLevelItems
}

func (p *Parser) parseCategory() bool {
	if !p.expectPeek(IDENT) {
		return false
	}

	p.currentItem = nil
	p.currentCategory = &Category{Token: p.curToken, Value: p.curToken.Literal}
	p.parseAnnotations()
	p.parseItem()
	return true
}

func (p *Parser) parseItem() {
	var level = p.findLevel()
	var item = &Item{Category: p.currentCategory}
	if p.expectPeek(DASH) {
		item.Type = Project
	} else if p.expectPeek(STAR) {
		item.Type = Task
	} else if p.expectPeek(PLUS) {
		item.Type = Pointer
	} else {
		return
	}

	if level == 0 {
		p.firstLevelItems = append(p.firstLevelItems, item)
	}
	if level > p.currentLevel {
		p.currentItem.Children = append(p.currentItem.Children, item)
	}
	p.currentItem = item
	p.currentLevel = level

	p.parseMarker()
	p.parseName()
	p.parseDescription()
	p.parseAnnotations()
	p.parseTarget()

	p.expectSkip(RP)
	p.expectSkip(SEMICOLON)

	if p.peekTokenIs(LEVEL) || p.peekTokenIs(DASH) || p.peekTokenIs(STAR) {
		p.parseItem()
	}
}

func (p *Parser) parseMarker() {
	if p.peekTokenIs(LP) {
		p.currentItem.Marked = true
		p.nextToken()
	}
}

func (p *Parser) parseName() {
	if !p.expectPeek(IDENT) {
		return
	}
	p.currentItem.Name = &Name{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseDescription() {
	if !p.expectPeek(STRING) {
		return
	}
	p.currentItem.Description = &Description{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseTarget() {
	if !p.expectPeek(DUALARROW) {
		return
	}
	if !p.expectPeek(IDENT) {
		return
	}
	p.currentItem.Target = p.curToken.Literal
}

func (p *Parser) parseAnnotations() {
	if !p.expectPeek(LSB) {
		return
	}
	for !p.peekTokenIs(RSB) {
		p.parseAnnotation()
	}
	p.expectSkip(RSB)
}

func (p *Parser) parseAnnotation() {
	if !p.expectPeek(IDENT) {
		return
	}
	if p.peekTokenIs(RSB) || p.peekTokenIs(COMMA) {
		unary := &UnaryAnnotation{Token: p.curToken, Name: Name{Token: p.curToken, Value: p.curToken.Literal}}
		p.addAnnotationTo(unary)
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
		p.addAnnotationTo(binary)
	}
	p.expectSkip(COMMA)
}

func (p *Parser) addAnnotationTo(ann Annotation) {
	if p.currentItem != nil {
		p.currentItem.Annotations = append(p.currentItem.Annotations, ann)
	} else if p.currentCategory != nil {
		p.currentCategory.Annotations = append(p.currentCategory.Annotations, ann)
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
