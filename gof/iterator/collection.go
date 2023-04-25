package main

type collection struct {
	users []*user
}

func (c *collection) createIterator() *iterator {
	return &iterator{
		users: c.users,
	}
}
