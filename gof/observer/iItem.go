package main

type iItem interface {
	register(iObserver)
	allNotify()
}
