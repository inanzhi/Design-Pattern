package strategy_pattern

import (
	"fmt"
	"testing"
)

// 策略接口 策略执行者 策略设置到执行者中 策略执行
// 武器接口 机器人  机器人装备武器 机器人攻击
// 定义一组算法，将每个算法都封装起来，并且使他们之间可以互换

// IStrategy 不同的情况使用不同的策略
type IStrategy interface {
	do(int, int) int
}

// Operator 具体策略的执行者 里面包含策略方法
type Operator struct {
	strategy IStrategy
}

// 操作者设置策略
func (operator *Operator) setStrategy(strategy IStrategy) {
	operator.strategy = strategy
}

// 操作者执行方法
func (operator *Operator) calculate(a, b int) int {
	return operator.strategy.do(a, b)
}

// 策略A本体
type add struct{}

// 策略A具体的策略
func (*add) do(a, b int) int {
	return a + b
}

// 策略B本体
type reduce struct{}

// 策略B的实现
func (*reduce) do(a, b int) int {
	return a - b
}

func TestStrategy(t *testing.T) {
	operator := &Operator{}
	operator.setStrategy(&add{})
	result := operator.calculate(1, 2)
	fmt.Println(result)

	operator.setStrategy(&reduce{})
	result = operator.calculate(5, 3)
	fmt.Println(result)

}
