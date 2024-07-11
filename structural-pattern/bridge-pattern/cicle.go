package bridge_pattern
//
//// Circle 是 Shape 的具体实现
//type Circle struct {
//	*AbstractShape
//}
//
//func NewCircle(color Color) *Circle {
//	return &Circle{AbstractShape: NewAbstractShape(color)}
//}
//
//func (c *Circle) Draw() string {
//	return "Circle painted in " + c.color.Paint()
//}



// Circle 是 Shape 的具体实现
type Circle struct {
	color Color
}

func NewCircle(color Color) *Circle {
	return &Circle{color}
}

func (c *Circle) Draw() string {
	return "Circle painted in " + c.color.Paint()
}