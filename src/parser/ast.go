package parser

type Node interface {
  TokenLiteral() string
}

type Valuable struct {
  Token         Token
  Name          *Name
  Description   *Description
  Level         *Level
  Children      []*Valuable
}
func (s *Valuable) TokenLiteral() string { return s.Token.Literal }

type Name struct {
  Token         Token
  Value         string
}
func (s *Name) TokenLiteral() string { return s.Token.Literal }

type Description struct {
  Token         Token
  Value         string
}
func (s *Description) TokenLiteral() string { return s.Token.Literal }

type Level struct {
  Token         Token
  Value         int
}
func (s *Level) TokenLiteral() string { return s.Token.Literal }

