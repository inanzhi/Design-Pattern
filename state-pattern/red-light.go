package state_pattern

import "fmt"

type RedLight struct{}

func (r *RedLight) ChangeToRedLight() {
	fmt.Println("已经是红灯")
}

func (r *RedLight) ChangeToYellowLight() {
	fmt.Println("红灯变为黄灯")
}

func (r *RedLight) ChangeToGreenLight() {
	fmt.Println("红灯变为绿灯")
}

// YellowLight 和 GreenLight 类似地实现...
