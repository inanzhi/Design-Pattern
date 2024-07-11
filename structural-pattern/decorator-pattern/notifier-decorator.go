package decorator_pattern

//在 Go 语言中，是否提供装饰器基类取决于具体的需求和场景。如果装饰器较少且简单，可以直接实现每个装饰器而不提供基类。
//如果装饰器较多且共享通用逻辑，提供一个基类可以减少代码重复，
//提高代码的可维护性和可扩展性。最终的选择应基于项目的复杂性和开发团队的偏好。

// NotifierDecorator 装饰器基类，实现 Notifier 接口
type NotifierDecorator struct {
	notifier Notifier
}

func (d *NotifierDecorator) Send(message string) {
	d.notifier.Send(message)
}
