package visitor_pattern

type Shape interface {
	GetType() string

	//接受一个访客，开了一个后门，允许后门通过这个方法间接执行后门结构体提供的方法
	Accept(Visitor)
}
