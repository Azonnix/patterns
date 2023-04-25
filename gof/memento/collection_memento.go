package main

type collectionMemento struct {
	queueMemento []*memento
}

func (m *collectionMemento) addMemento(newMemento *memento) {
	m.queueMemento = append(m.queueMemento, newMemento)
}

func (m *collectionMemento) getMemento() *memento {
	if len(m.queueMemento) > 0 {
		currentMemento := m.queueMemento[len(m.queueMemento)-1]
		m.queueMemento = m.queueMemento[:len(m.queueMemento)-1]
		return currentMemento
	}
	return nil
}
