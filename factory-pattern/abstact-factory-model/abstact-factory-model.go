package abstact_factory_model

import "fmt"

type Person interface {
	Greet()
}

type person struct {
	name string
	age  int
}

// 通过该方法 结构体实现了接口 所以可以相互赋值
func (p person) Greet() {
	fmt.Printf("Hi! My name is %s!\n", p.name)
}

// 可以使用不同的工厂函数来返回不同的接口实现
// NewPerson 返回一个接口，而不是一个结构体
func NewPerson(name string, age int) Person {
	return person{
		name: name,
		age:  age,
	}
}
