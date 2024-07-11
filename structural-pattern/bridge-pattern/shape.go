package bridge_pattern

type Shape interface {
	Draw()
}


//多套一层抽象类也可以  不影响  结构体A套结构体B  则结构体A可以直接调用结构体B中的方法
//// AbstractShape 是形状的抽象类
//type AbstractShape struct {
//	color Color
//}
//
//func (as *AbstractShape) Draw() {
//	as.color.Paint()
//}
//
//// NewShape 创建一个有颜色的形状
//func NewAbstractShape(color Color) *AbstractShape {
//	return &AbstractShape{color: color}
//}
//
//func (as *AbstractShape) SetColor(color Color) {
//	as.color = color
//}


