package main

type observer interface {
	update(string)
	getEmail() string
}
