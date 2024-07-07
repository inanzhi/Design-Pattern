package mediator_pattern

import "fmt"

// 拉货的火车  火车里面要装通信接口装置
type GoodsTrain struct {
	Mediator Mediator
}

func (g *GoodsTrain) RequestArrival() {
	//请求通信接口，问自己是否可以停靠，具体指挥者告诉你能不能停靠
	if g.Mediator.canLand(g) {
		fmt.Println("GoodsTrain: Landing")
	} else {
		fmt.Println("GoodsTrain: Waiting")
	}
}

func (g *GoodsTrain) Departure() {
	//车子离开后，要使用车子的通信讯号发消息给调度员 调度员通知空闲了
	g.Mediator.notifyFree()
	fmt.Println("GoodsTrain: Leaving")
}

func (g *GoodsTrain) StartArrival() {
	fmt.Println("GoodsTrain: Arrival Permitted,start landing")
}
