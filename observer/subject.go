package main

type Publisher interface {
	subscribe(observer Subscriber)
	unsubscribe(observer Subscriber)
	notifyAll()
}
