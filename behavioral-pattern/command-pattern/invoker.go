package command_pattern

// Invoker 用户    用户要包含命令的抽象功能   用户设置具体命令是什么   用户执行命令
type Invoker struct {
	command ICommand
}

func (i *Invoker) SetCommand(command ICommand) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}
