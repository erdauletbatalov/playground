package main

type Subscriber interface {
	update(string)
	getID() string
}
