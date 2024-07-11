package decorator_pattern

import "fmt"

// ConcreteNotifier 具体的通知者，实现 Notifier 接口
type ConcreteNotifier struct{}

func (n *ConcreteNotifier) Send(message string) {
	fmt.Println("Sending notification:", message)
}
