package command_pattern

//具体命令 包含执行者  执行方式->请求执行者执行

// 命令中要包含接受者
// ConcreteCommand1 实现 Command 接口并封装接收者的具体操作
type ConcreteCommand1 struct {
	Receiver *Receiver
}

func (c *ConcreteCommand1) Execute() {
	c.Receiver.Action1()
}

// ConcreteCommand2 实现 Command 接口并封装接收者的具体操作
type ConcreteCommand2 struct {
	Receiver *Receiver
}

func (c *ConcreteCommand2) Execute() {
	c.Receiver.Action2()
}
