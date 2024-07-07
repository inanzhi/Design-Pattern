package mediator_pattern

import "fmt"

// 拉客的火车 火车里面要装通信接口装置
type PassengerTrain struct {
	Mediator Mediator
}

func (g *PassengerTrain) RequestArrival() {
	if g.Mediator.canLand(g) {
		fmt.Println("PassengerTrain: Landing")
	} else {
		fmt.Println("PassengerTrain: Waiting")
	}
}

func (g *PassengerTrain) Departure() {
	fmt.Println("PassengerTrain: Leaving")
	g.Mediator.notifyFree()
}

func (g *PassengerTrain) StartArrival() {
	fmt.Println("PassengerTrain: Arrival Permitted,start Landing")
}
