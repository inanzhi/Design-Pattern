package bridge_pattern

// Shape 是形状的抽象类
type Shape struct {
	color Color
}

// NewShape 创建一个有颜色的形状
func NewShape(color Color) *Shape {
	return &Shape{color: color}
}

func (s *Shape) SetColor(color Color) {
	s.color = color
}

// Circle 是 Shape 的具体实现
type Circle struct {
	*Shape
}

func NewCircle(color Color) *Circle {
	return &Circle{Shape: NewShape(color)}
}

func (c *Circle) Draw() string {
	return "Circle painted in " + c.color.Paint()
}

// Square 是 Shape 的具体实现
type Square struct {
	*Shape
}

func NewSquare(color Color) *Square {
	return &Square{Shape: NewShape(color)}
}

func (s *Square) Draw() string {
	return "Square painted in " + s.color.Paint()
}
