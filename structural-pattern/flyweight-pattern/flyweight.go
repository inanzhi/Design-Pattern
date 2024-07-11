package flyweight_pattern

import "fmt"

// Flyweight interface
type Flyweight interface {
	Operation(extrinsicState string)
}

// ConcreteFlyweight 实现 Flyweight 接口
type ConcreteFlyweight struct {
	intrinsicState string
}

func (c *ConcreteFlyweight) Operation(extrinsicState string) {
	fmt.Printf("Intrinsic State = %s, Extrinsic State = %s\n", c.intrinsicState, extrinsicState)
}
