package main

import "fmt"

type doctor struct {
	next iDepartment
}

func (d *doctor) setNext(next iDepartment) {
	d.next = next
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor already check up!")
		d.next.execute(p)
		return
	}

	p.doctorCheckUpDone = true
	fmt.Println("Doctor check up!")
	d.next.execute(p)
}
