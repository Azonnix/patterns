package main

type car struct {
	name      string
	observers []iObserver
}

func (c *car) register(newObserver iObserver) {
	c.observers = append(c.observers, newObserver)
}

func (c *car) allNotify() {
	for _, o := range c.observers {
		o.update(c.name)
	}
}
