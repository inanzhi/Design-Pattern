package factory_method_model

import "fmt"

type Person struct {
	name string
	age  int
}

// NewPersonFactory 第一个参数为统一参数，第二个则可以修改
func NewPersonFactory(age int) func(name string) Person {
	return func(name string) Person {
		return Person{
			name: name,
			age:  age,
		}
	}
}
func use() {
	newBaby := NewPersonFactory(1)
	baby := newBaby("john")
	fmt.Println(baby)

	newTeenager := NewPersonFactory(16)
	teen := newTeenager("Alice")
	fmt.Println(teen)
}
