package null_object_pattern

import "fmt"

//注意

//在Go语言中，接口（interface）的设计目的是为了定义行为契约，而不是为了构造具体类型的实例。
//接口中不能包含构造方法或者说初始化逻辑，因为接口本身并不包含任何状态，也不直接实例化任何对象。
//接口仅仅定义了一组方法签名，任何实现了这些方法签名的类型都被认为实现了这个接口。
//
//Go语言没有传统意义上的构造函数，但在实践中，可以通过定义函数来初始化结构体并返回一个实例，这通常被称为工厂函数或构造器函数。
//这些函数可以位于任何地方，包括实现了接口的类型中，但它们不属于接口定义的一部分。

// 大学
type College interface {
	AddDepartment(department department)
	FindDepartment(departmentName string) department
}

type college struct {
	collegeName string
	departments []department
}

func NewCollege(name string) College {
	fmt.Printf("%s大学创建成功！\n", name)
	return college{collegeName: name}
}

func (c college) AddDepartment(department department) {
	c.departments = append(c.departments, department)
}

func (c college) FindDepartment(departmentName string) department {
	for _, departments := range c.departments {
		if departmentName == departments.GetName() {
			return departments
		}
	}
	return &NullDepartment{}
}
