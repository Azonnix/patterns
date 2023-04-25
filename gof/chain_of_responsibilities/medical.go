package main

import "fmt"

type medical struct {
	next iDepartment
}

func (m *medical) setNext(next iDepartment) {
	m.next = next
}

func (m *medical) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("Medicine already done!")
		return
	}

	p.medicineDone = true
	fmt.Println("Medicine done!")
}
