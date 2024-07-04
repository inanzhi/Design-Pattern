package template_pattern

import (
	"fmt"
	"testing"
)

//
//// 做菜方法集合 做菜流程
//// Cooker 包换所有步骤的方法
//type Cooker interface {
//	fire()
//	cooke()
//	outFire()
//}
//
//// 模版方法 也就是算法骨架 将过程串联
//func doCook(cook Cooker) {
//	cook.fire()
//	cook.cooke()
//	cook.outFire()
//}
//
//// 厨师  包含算法的骨架方法
//type Chef struct{}
//
//func (Chef) fire() {
//	fmt.Println("Fire")
//}
//
//// 做菜，中间过程，交给子类实现
//func (Chef) cooke() {
//}
//
//func (Chef) outFire() {
//	fmt.Println("OutFire")
//}
//
//// 不同的具体实现，实心接口方法
//type dishesYouBaoDaXia struct {
//	Chef
//}
//
//func (d dishesYouBaoDaXia) cook() {
//	fmt.Println("做油爆大虾")
//}
//
//// 不同的具体实现，实心接口方法
//type dishesLongXuMian struct {
//	Chef
//}
//
//func (d dishesLongXuMian) cook() {
//	fmt.Println("做龙须面")
//}
//
//func TestTemplate(t *testing.T) {
//	//做油爆大虾
//	dishesYouBaoDaXia := &dishesYouBaoDaXia{}
//	doCook(dishesYouBaoDaXia)
//
//	fmt.Println("开始做另一道菜...")
//
//	//做龙须面
//	dishesLongXuMian := &dishesLongXuMian{}
//	doCook(dishesLongXuMian)
//}

// Document 是一个接口，定义了读取和写入操作
type Document interface {
	Read() string
	Write(content string)
}

type DocumentTemplate struct {
	Document
}

// Process 是模版方法，定义了算法的骨架
func (dt *DocumentTemplate) Process(content string) {
	data := dt.Read()
	fmt.Println("Reading data:", data)
	dt.Write(content)
	fmt.Println("Writing data:", content)
}

type TextDocument struct {
	content string
}

// Read 实现了 Document 接口的 Read 方法
func (td *TextDocument) Read() string {
	return td.content
}

// Write 实现了 Document 接口的 Write 方法
func (td *TextDocument) Write(content string) {
	td.content = content
}

func main() {
	textDoc := &TextDocument{content: "Initial content"}
	template := &DocumentTemplate{Document: textDoc}

	template.Process("New content")
}

// 做菜方法集合 做菜流程
// Cooker 包换所有步骤的方法
type Cooker interface {
	fire()
	cooke()
	outFire()
}

// 厨师  包含算法的骨架方法
type Chef struct {
	Cooker
}

// 模版方法 也就是算法骨架 将过程串联
func (chef Chef) doCook() {
	chef.fire()
	chef.cooke()
	chef.outFire()
}

func (Chef) fire() {
	fmt.Println("Fire")
}

// 做菜，中间过程，交给子类实现
func (Chef) cooke() {
}

func (Chef) outFire() {
	fmt.Println("OutFire")
}

// 不同的具体实现，实心接口方法
type dishesYouBaoDaXiaChef struct {
	Chef
}

func (chef dishesYouBaoDaXiaChef) cook() {
	fmt.Println("做油爆大虾")
}

// 不同的具体实现，实心接口方法
type dishesLongXuMianChef struct {
	Chef
}

func (chef dishesLongXuMianChef) cook() {
	fmt.Println("做龙须面")
}

func TestTemplate(t *testing.T) {
	//做油爆大虾

	dishesYouBaoDaXiaChef := &dishesYouBaoDaXiaChef{}
	chef := &Chef{dishesYouBaoDaXiaChef}
	chef.doCook()

	fmt.Println("开始做另一道菜...")

	//做龙须面
	dishesLongXuMian := &dishesLongXuMianChef{}
	dishesLongXuMian.doCook()
}
