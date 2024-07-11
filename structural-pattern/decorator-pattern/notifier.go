package decorator_pattern

import "fmt"

// Notifier 接口，定义发送通知的方法
type Notifier interface {
	Send(message string)
}

// EmailNotifier 邮件通知装饰器
type EmailNotifier struct {
	NotifierDecorator
}

func NewEmailNotifier(notifier Notifier) *EmailNotifier {
	return &EmailNotifier{NotifierDecorator{notifier}}
}

func (e *EmailNotifier) Send(message string) {
	e.notifier.Send(message)
	fmt.Println("Sending email notification:", message)
}

// SMSNotifier 短信通知装饰器
type SMSNotifier struct {
	NotifierDecorator
}

func NewSMSNotifier(notifier Notifier) *SMSNotifier {
	return &SMSNotifier{NotifierDecorator{notifier}}
}

func (s *SMSNotifier) Send(message string) {
	s.notifier.Send(message)
	fmt.Println("Sending SMS notification:", message)
}
