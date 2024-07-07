package visitor_pattern

import "fmt"

type AreaCalculator struct {
	name string
	area float64
}

func NewAreaCalculator(name string) Visitor {

	return AreaCalculator{name, 0}
}

func (a AreaCalculator) visitForSquare(s *Square) {
	//Calculate area for square. After calculating the area assign in to the area instance variable
	fmt.Printf("Calculating %s area success!\n", s.GetType())
}

func (a AreaCalculator) visitForCircle(c *Circle) {
	//Calculate are for circle. After calculating the area assign in to the area instance variable
	fmt.Printf("Calculating %s area success!\n", c.GetType())
}
