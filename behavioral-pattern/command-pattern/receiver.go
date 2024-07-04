package command_pattern

import "fmt"

// Receiver 实际执行命令的对象  执行命令
type Receiver struct{}

func (r *Receiver) Action1() {
	fmt.Println("Receiver: Executing Action1")
}

func (r *Receiver) Action2() {
	fmt.Println("Receiver: Executing Action2")
}
