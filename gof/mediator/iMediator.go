package main

type iMediator interface {
	canLand(train iTrain) bool
	notifyFree()
}
