package language

import (
	"fmt"
)

type Parser struct {
	filename        string
	lexer           *Lexer
	curToken        Token
	peekToken       Token
	Errors          []ParseError
	currentCategory *Category
	currentItem     *Item
	currentParent   *Item
	currentLevel    int
	firstLevelItems []*Item
	categories      []*Category
}

type ParseError struct {
	origin   string
	expected []TokenType
	got      TokenType
	literal  string
	line     int
	column   int
}

func (pe *ParseError) message() string {
	return fmt.Sprintf("Parse error, expected '%s', got '%s'", pe.expected, pe.got)
}

func NewParser(filename string, lexer *Lexer) *Parser {
	parser := &Parser{
		lexer:    lexer,
		filename: filename,
	}

	parser.nextToken()
	return parser
}

func (p *Parser) Parse() ([]*Category, []*Item) {
	for !p.peekTokenIs(EOF) {
		if !p.parseCategory() {
			for _, err := range p.Errors {
				fmt.Printf("Error parsing file %s while doing '%s' (line %d,column %d). Expected %s, got %s (%s) \n",
					p.filename, err.origin, err.line, err.column, err.expected, err.got, err.literal)
			}
			break
		}
	}
	return p.categories, p.firstLevelItems
}

func (p *Parser) parseCategory() bool {
	if !p.expectPeek("Category", IDENT) {
		return false
	}

	p.currentItem = nil
	p.currentCategory = &Category{Token: p.curToken, Value: p.curToken.Literal}
	p.categories = append(p.categories, p.currentCategory)
	p.parseAnnotations()
	p.parseItem()
	return true
}

func (p *Parser) parseItem() {
	var level = p.findLevel()
	var item = &Item{Category: p.currentCategory}
	if p.couldPeek(DASH) {
		item.Type = Normal
	} else if p.couldPeek(STAR) {
		item.Type = Star
	} else if p.couldPeek(PLUS) {
		item.Type = Plus
	} else {
		p.peekErrors(DASH, STAR, PLUS)
		return
	}

	if level == 0 {
		p.firstLevelItems = append(p.firstLevelItems, item)
	} else {
		currentParent := p.firstLevelItems[len(p.firstLevelItems)-1]
		currentParent.Children = append(currentParent.Children, item)
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

	if p.peekTokenIs(LEVEL) || p.peekTokenIs(DASH) || p.peekTokenIs(STAR) || p.peekTokenIs(PLUS) {
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
	if !p.expectPeek("Name", IDENT) {
		return
	}
	if p.peekTokenIs(DUALCOLON) {
		p.currentItem.PreName = &Name{Token: p.curToken, Value: p.curToken.Literal}
		p.expectSkip(DUALCOLON)
		if !p.expectPeek("Name", IDENT) {
			return
		}
		p.currentItem.Name = &Name{Token: p.curToken, Value: p.curToken.Literal}
	} else {
		p.currentItem.Name = &Name{Token: p.curToken, Value: p.curToken.Literal}
	}
}

func (p *Parser) parseDescription() {
	if !p.couldPeek(STRING) {
		return
	}
	p.currentItem.Description = &Description{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseTarget() {
	if !p.couldPeek(DUALARROW) {
		return
	}
	if !p.expectPeek("Target", IDENT) {
		return
	}
	p.currentItem.Target = p.curToken.Literal
}

func (p *Parser) parseAnnotations() {
	if !p.couldPeek(LSB) {
		return
	}
	for !p.peekTokenIs(RSB) {
		p.parseAnnotation()
	}
	p.expectSkip(RSB)
}

func (p *Parser) parseAnnotation() {
	if !p.expectPeek("Annotations", IDENT) {
		return
	}
	if p.peekTokenIs(RSB) || p.peekTokenIs(DOT) {
		unary := &UnaryAnnotation{Token: p.curToken, Name: Name{Token: p.curToken, Value: p.curToken.Literal}}
		p.addAnnotationTo(unary)
	} else if p.peekTokenIs(DASH) {
		left := &Name{Token: p.curToken, Value: p.curToken.Literal}
		p.expectPeek("Annotations : Binary", DASH)
		operator := Operator{Token: p.curToken}
		if !p.expectPeek("Annotations : Binary : Identity", IDENT) {
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
	p.expectSkip(DOT)
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
	for p.couldPeek(LEVEL) {
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

func (p *Parser) couldPeek(t TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	return false
}

func (p *Parser) expectPeek(origin string, t TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(origin, t)
		return false
	}
}

func (p *Parser) peekTokenIs(t TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) curTokenIs(t TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekError(origin string, t TokenType) {
	parseError := ParseError{expected: []TokenType{t}, got: p.peekToken.Type, column: p.lexer.col - 1, line: p.lexer.line, literal: p.peekToken.Literal, origin: origin}
	p.Errors = append(p.Errors, parseError)
}

func (p *Parser) peekErrors(t ...TokenType) {
	parseError := ParseError{expected: t, got: p.peekToken.Type, column: p.lexer.col - 1, line: p.lexer.line, literal: p.peekToken.Literal}
	p.Errors = append(p.Errors, parseError)
}
