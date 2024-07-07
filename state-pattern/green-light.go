package state_pattern

import "fmt"

type GreenLight struct{}

func (r *GreenLight) ChangeToGreenLight() {
	fmt.Println("已经是绿灯")
}

func (r *GreenLight) ChangeToRedLight() {
	fmt.Println("绿灯变成红灯")
}

func (r *GreenLight) ChangeToYellowLight() {
	fmt.Println("绿灯变为黄灯")
}
