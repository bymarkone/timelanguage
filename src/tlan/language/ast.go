package language

type Node interface {
	TokenLiteral() string
}

type Item struct {
	Token       Token
	Type        *ItemType
	Name        *Name
	Description *Description
	Category    *Category
	Children    []*Item
	Marked      bool
	Annotations []Annotation
}

func (s *Item) TokenLiteral() string { return s.Token.Literal }

type ItemType struct {
	Token Token
	Value string
}

func (s *ItemType) TokenLiteral() string { return s.Token.Literal }

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
	Annotations []Annotation
}

func (i *Category) TokenLiteral() string { return i.Token.Literal }

type Annotation interface {
	Type() string
}

const (
	UNARY  = "UNARY"
	BINARY = "BINARY"
)

type UnaryAnnotation struct {
	Token Token
	Name  *Name
}

func (i *UnaryAnnotation) TokenLiteral() string { return i.Token.Literal }
func (i *UnaryAnnotation) Type() string         { return UNARY }

type BinaryAnnotation struct {
	Token    Token
	Left     *Name
	Right    *Name
	Operator *Operator
}

func (i *BinaryAnnotation) TokenLiteral() string { return i.Token.Literal }
func (i *BinaryAnnotation) Type() string         { return BINARY }

type Operator struct {
	Token Token
}

func (i *Operator) TokenLiteral() string { return i.Token.Literal }
