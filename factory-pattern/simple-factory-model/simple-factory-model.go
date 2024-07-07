package simple_factory_model

import "fmt"

// Person 简单工厂模式
// 相比于 p:=&Person{}可以确保实例具有需要的参数
// 但缺点就是要构建不同参数类型和数量的产品 就需要修改用来创建的函数，耦合性高，于是我们使用工厂方法模式
type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("My name is %s", p.Name)
}
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}
