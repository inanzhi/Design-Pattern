package mediator_pattern

// Mediator 接口，定义了组件之间通信的方法
type Mediator interface {

	//某辆车是否可以停靠
	canLand(train) bool
	// NotifyFree 通知站台空闲了
	notifyFree()
}
