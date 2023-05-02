package main

import "fmt"

type client struct {
	email string
}

func (c *client) update(owner string) {
	fmt.Println(`notify to ` + c.email + `about` + owner)
}

func (c *client) getEmail() string {
	return c.email
}
