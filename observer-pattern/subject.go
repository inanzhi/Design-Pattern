package observer_pattern

import "fmt"

// 观察者模式是一种行为型设计模式。这种模式允许一个实例（可以称为目标对象）发布各种事件（event）给其他实例（观察者）。这些观察者会对目标对象进行订阅，这样每当目标对象发生变化时，观察者就会收到事件（event）通知。
//
// 来看个例子：在电商网站经常会有各种商品脱销，假如有顾客已经在关注这些商品，这些商品的脱销就会对他们产生不好的体验。如果顾客还想买这些商品，那么通常有如下解决方案：
//
// 顾客以一定的频率检查这些商品是否在售
// 电商平台将所有的上架的商品信息定期推送给用户
// 顾客只订阅他所关注的特定商品的信息，当这些商品再次上架时他们会收到通知；多个顾客可以订阅同一个商品的信息
// 选项3是最为可行的一种方案。这也正是观察者模式所能做到的事情。观察者模式的核心组件为：
//
// Subject ： 目标对象，是有变化发生时就会发布相关事件的实例
// Observer ： 观察者，订阅目标对象的信息，会收到一些特定事件的通知

// ISubject 主题接口
type ISubject interface {
	RegisterObserver(observer IObserver)
	DeregisterObserver(observer IObserver)
	notifyAllObserverUpdateSubject()
}

// Subject 主题
type Subject struct {
	//调用关注该主题的观察者的内部更新自己的主题对象   每个观察者都有一个主题对象

	observerList []IObserver
	name         string
	//消息
	content string
	//有新消息
	newMessage bool
}

// RegisterObserver 观察者注册主题
func (s *Subject) RegisterObserver(observer IObserver) {
	s.observerList = append(s.observerList, observer)
}

func (s *Subject) DeregisterObserver(o IObserver) {
	s.observerList = removeFromSlice(s.observerList, o)
}

// notifyAllObserverUpdateSubject 主题通知每个内部的观察者，让它们自己进行更新
func (s *Subject) notifyAllObserverUpdateSubject() {
	for _, observer := range s.observerList {
		observer.UpdateSubject(s.content)
	}
}

func NewSubject(name string) *Subject {
	return &Subject{
		name: name,
	}
}

func removeFromSlice(observerList []IObserver, observerToRemove IObserver) []IObserver {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.GetObserverName() == observer.GetObserverName() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

// SetNewMessage 主题设置状态 ，通知每一个观察者，观察者然后进行更新
func (s *Subject) SetNewMessage(newMessage bool, content string) {
	fmt.Printf("Subject %s  now  have  newMessage\n", s.name)
	s.newMessage = newMessage
	s.content = content
	s.notifyAllObserverUpdateSubject()
}
