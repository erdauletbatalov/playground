package main

import "fmt"

type User struct {
	id string
}

func (c *User) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *User) getID() string {
	return c.id
}
