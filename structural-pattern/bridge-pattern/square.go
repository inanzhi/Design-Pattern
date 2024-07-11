package bridge_pattern


//// Square 是 Shape 的具体实现
//type Square struct {
//	*AbstractShape
//}
//
//func NewSquare(color Color) *Square {
//	return &Square{AbstractShape: NewAbstractShape(color)}
//}
//
//func (s *Square) Draw() string {
//	return "Square painted in " + s.color.Paint()
//}


type Square struct {
	color Color
}

func NewSquare(color Color) *Square {
	return &Square{color}
}

func (s *Square) Draw() string {
	return "Square painted in " + s.color.Paint()
}