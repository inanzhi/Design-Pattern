package observer_pattern

import "fmt"

// IObserver 定义了观察者接口
type IObserver interface {
	UpdateSubject(string)
	GetObserverName() string
}

// Observer 观察者
type Observer struct {
	//观察者关注的主题
	subject *Subject
	//观察者的名称
	name string
}

func (o *Observer) UpdateSubject(content string) {
	fmt.Printf("Observer %s received state change notification: %s\n", o.name, content)
}
func (o *Observer) GetObserverName() string {
	return o.name
}

func (o *Observer) GetContent() string {
	return o.subject.content
}

func NewObserver(subject *Subject, name string) *Observer {
	return &Observer{
		subject: subject,
		name:    name,
	}
}
