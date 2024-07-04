package chain_of_responsibility_pattern

import "fmt"

// 收银处
type Cashier struct {
	next Department
}

func (c *Cashier) Execute(p *Patient) {
	if p.PaymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from Patient Patient")
	c.next.Execute(p)
}

func (c *Cashier) SetNext(next Department) {
	c.next = next
}
