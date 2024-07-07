package visitor_pattern

type Visitor interface {
	visitForSquare(*Square)

	visitForCircle(*Circle)
}
