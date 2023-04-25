package main

type iterator struct {
	users []*user
	index int
}

func (i *iterator) hasNext() bool {
	if i.index < len(i.users) {
		return true
	}
	return false
}

func (i *iterator) getNext() *user {
	if i.hasNext() {
		user := i.users[i.index]
		i.index++
		return user
	}
	return nil
}
