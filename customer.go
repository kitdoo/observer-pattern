package main

import "fmt"

type customer struct {
	email string
}

func (c *customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s about availability %s\n", c.email, itemName)
}

func (c *customer) getEmail() string {
	return c.email
}
