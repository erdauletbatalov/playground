package main

import "fmt"

type Item struct {
	subscriberList []Subscriber
	name           string
	inStock        bool
}

func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}
func (i *Item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}
func (i *Item) subscribe(o Subscriber) {
	i.subscriberList = append(i.subscriberList, o)
}

func (i *Item) unsubscribe(o Subscriber) {
	i.subscriberList = removeFromSlice(i.subscriberList, o)
}

func (i *Item) notifyAll() {
	for _, subscriber := range i.subscriberList {
		subscriber.update(i.name)
	}
}

func removeFromSlice(subList []Subscriber, subToRemove Subscriber) []Subscriber {
	subListLen := len(subList)
	for i, sub := range subList {
		if subToRemove.getID() == sub.getID() {
			subList[subListLen-1], subList[i] = subList[i], subList[subListLen-1]
			return subList[:subListLen-1]
		}
	}
	return subList
}
