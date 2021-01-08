package parser

type Node interface {
	TokenLiteral() string
}

type Item struct {
	Token       Token
	Name        *Name
	Description *Description
	Category    *Category
	Children    []*Item
	Marked      bool
}

func (s *Item) TokenLiteral() string { return s.Token.Literal }

type Name struct {
	Token Token
	Value string
}

func (s *Name) TokenLiteral() string { return s.Token.Literal }

type Description struct {
	Token Token
	Value string
}

func (s *Description) TokenLiteral() string { return s.Token.Literal }

type Category struct {
	Token Token
	Value string
}

func (i *Category) TokenLiteral() string { return i.Token.Literal }
