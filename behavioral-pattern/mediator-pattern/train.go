package mediator_pattern

type train interface {

	//请求、通知停靠
	RequestArrival()

	//离开
	Departure()

	//开始停靠
	StartArrival()
}
