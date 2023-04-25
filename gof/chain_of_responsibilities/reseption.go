package main

import "fmt"

type reseption struct {
	next iDepartment
}

func (r *reseption) setNext(next iDepartment) {
	r.next = next
}

func (r *reseption) execute(p *patient) {
	if p.registrationDone {
		fmt.Println("Registration already done!")
		r.next.execute(p)
		return
	}

	p.registrationDone = true
	fmt.Println("Registration done!")
	r.next.execute(p)
}
