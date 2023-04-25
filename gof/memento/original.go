package main

type original struct {
	state string
}

func (o *original) setState(newState string) {
	o.state = newState
}

func (o *original) getState() string {
	return o.state
}

func (o *original) saveState() *memento {
	return NewMemento(o.state)
}

func (o *original) loadState(m *memento) {
	o.state = m.getState()
}
