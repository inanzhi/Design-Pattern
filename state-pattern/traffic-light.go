package state_pattern

import "fmt"

type TrafficLight struct {
	name         string
	currentState LightState
}

func NewTrafficLight(name string, state LightState) *TrafficLight {
	fmt.Printf("%s创建成功!\n", name)
	return &TrafficLight{name, state}
}

func (tl *TrafficLight) SetState(state LightState) {
	tl.currentState = state
}

func (tl *TrafficLight) SwitchToRed() {
	tl.currentState.ChangeToRedLight()
}

func (tl *TrafficLight) SwitchToYellow() {
	tl.currentState.ChangeToYellowLight()
}

func (tl *TrafficLight) SwitchToGreen() {
	tl.currentState.ChangeToGreenLight()
}
