package chain_of_responsibility_pattern

import "fmt"

type Medical struct {
	next Department
}

func (m *Medical) Execute(p *Patient) {

	if p.MedicineDone {
		fmt.Println("Medicine already given to Patient")
		if m.next != nil {
			m.next.Execute(p)
		}
		return
	}
	fmt.Println("Medical giving medicine to Patient")
	p.MedicineDone = true
	if m.next != nil {
		m.next.Execute(p)
	}
}

func (m *Medical) SetNext(next Department) {
	m.next = next
}
