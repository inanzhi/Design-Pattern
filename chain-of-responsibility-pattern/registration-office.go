package chain_of_responsibility_pattern

import "fmt"

type RegistrationOffice struct {
	next Department
}

func (r *RegistrationOffice) Execute(p *Patient) {
	if p.RegistrationDone {
		fmt.Println("Patient registration already done")
		r.next.Execute(p)
		return
	}
	fmt.Println("Reception registering Patient")
	p.RegistrationDone = true
	r.next.Execute(p)
}

func (r *RegistrationOffice) SetNext(next Department) {
	r.next = next
}
