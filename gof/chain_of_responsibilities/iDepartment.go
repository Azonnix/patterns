package main

type iDepartment interface {
	setNext(iDepartment)
	execute(*patient)
}
