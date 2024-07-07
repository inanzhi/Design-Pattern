package visitor_pattern

// Square 正方形
type Square struct {
	side int
}

func NewSquare(side int) *Square {
	return &Square{side}
}

func (s *Square) Accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *Square) GetType() string {
	return "Square"
}
