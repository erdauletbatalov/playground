package main

func main() {

	shirtItem := newItem("Nike Shirt")

	subscriber1 := &User{id: "abc@gmail.com"}
	subscriber2 := &User{id: "xyz@gmail.com"}

	shirtItem.subscribe(subscriber1)
	shirtItem.subscribe(subscriber2)

	shirtItem.updateAvailability()
}
