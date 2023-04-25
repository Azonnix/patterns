package main

import "sync"

type stationManager struct {
	trainQueue     []iTrain
	lock           *sync.Mutex
	isPlatformFree bool
}

func newStationManager() *stationManager {
	return &stationManager{
		lock:           &sync.Mutex{},
		isPlatformFree: true,
	}
}

func (m *stationManager) canLand(train iTrain) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.isPlatformFree {
		m.isPlatformFree = false
		return true
	}
	m.trainQueue = append(m.trainQueue, train)
	return false
}

func (m *stationManager) notifyFree() {
	m.lock.Lock()
	defer m.lock.Unlock()
	if !m.isPlatformFree {
		m.isPlatformFree = true
	}
	if len(m.trainQueue) > 0 {
		train := m.trainQueue[0]
		m.trainQueue = m.trainQueue[1:]
		train.permitArrival()
	}
}
