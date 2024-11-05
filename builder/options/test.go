package people

import "fmt"

type Person struct {
	age    int
	name   string
	salary float64
}

type PersonOptionFunc func(*Person)

func WithName(name string) PersonOptionFunc {
	return func(p *Person) {
		p.name = name
	}
}
func WithAge(age int) PersonOptionFunc {
	return func(p *Person) {
		p.age = age
	}
}
func NewPerson1(opts ...PersonOptionFunc) *Person {
	p := &Person{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

// usage:
func test1() {
	p := NewPerson1(WithName("Anton"), WithAge(25))
	fmt.Println(p) // {name: "Anton", age: 25}
}

// The downside here is that the amount of WithXXX functions are not obvious,
// and the consumers of the package would either have to search them in the
// package documentation or depend on their IDE autocomplete to show them the
// possible options. In my opinion, it does not give you any benefits over the
// Options struct, but bring the downside that the options are not obvious.

// Middle ground

// A middle ground between constructor with required positional arguments and
// optional arguments would be to have the required fields for a given struct as
// positional arguments so that the consumer MUST pass them, and have everything
// else as optional parameters, which may be skipped.

type PersonOptions struct {
	Age    int
	Name   string
	Salary float64
}

func NewPerson2(name string, options *PersonOptions) *Person {
	p := &Person{name: name}
	if options == nil {
		return p
	}
	if options.Age != 0 /* OR options.Age != nil */ {
		p.age = options.Age /* OR p.age = *options.Age */
	}
	if options.Salary != 0 /* OR options.Salary != nil */ {
		p.salary = options.Salary /* OR p.salary = *options.Salary */
	}
	return p
}

// usage:

func test2() {
	p := NewPerson2("Anton", &PersonOptions{Age: 25})
	fmt.Println(p) // {name: "Anton", age: 25}
}

// Summary

// There are 2 build-in ways to initialize a Go struct. Both are quite limited
// and more often than not they are not enough. That is why people came up with
// more sophisticated solutions (based on the built-in ones).

// The Solution: Options Pattern The Options pattern can be used to create
// objects with many optional parameters. In this pattern, we define a struct
// with optional parameters and provide methods to set those parameters. This
// pattern can be more concise than the Builder pattern and can be easier to use
// for objects with fewer parameters. Example In Golang, the Options pattern can
// be implemented using functional options. Functional options are functions
// that take a struct as an argument and return a modified version of that
// struct. Here's an example of how the Options pattern can be used to create a
// pizza object:
type Pizza struct {
	dough    string
	sauce    string
	cheese   string
	toppings []string
}
type PizzaOptions struct {
	Dough    string
	Sauce    string
	Cheese   string
	Toppings []string
}
type PizzaOption func(*PizzaOptions)

func WithDough(dough string) PizzaOption {
	return func(po *PizzaOptions) {
		po.Dough = dough
	}
}
func WithSauce(sauce string) PizzaOption {
	return func(po *PizzaOptions) {
		po.Sauce = sauce
	}
}
func WithCheese(cheese string) PizzaOption {
	return func(po *PizzaOptions) {
		po.Cheese = cheese
	}
}
func WithToppings(toppings []string) PizzaOption {
	return func(po *PizzaOptions) {
		po.Toppings = toppings
	}
}
func NewPizza(options ...PizzaOption) *Pizza {
	opts := &PizzaOptions{}
	for _, option := range options {
		option(opts)
	}
	pizza := &Pizza{
		dough:    opts.Dough,
		sauce:    opts.Sauce,
		cheese:   opts.Cheese,
		toppings: opts.Toppings,
	}
	return pizza
}

// In this example, we define the Pizza struct and the PizzaOptions struct,
// which is a struct with optional parameters. We then define functions to set
// each option, such as WithDough, WithSauce, and WithToppings. These functions
// return a PizzaOption that sets the corresponding field on the PizzaOptions
// struct. Finally, we define a NewPizza function that takes any number of
// PizzaOptions and constructs a Pizza object.
func test3() {
	pizza := NewPizza(
		WithDough("Regular"),
		WithSauce("Tomato"),
		WithCheese("Mozzarella"),
		WithToppings([]string{"Pepperoni", "Olives", "Mushrooms"}),
	)
	println(pizza.dough)
	println(pizza.sauce)
	println(pizza.cheese)
	println(pizza.toppings)
}

// The Options pattern can be a good alternative to the Builder pattern for
// creating objects with many optional parameters, especially if the object has
// fewer parameters. However, it can become unwieldy for objects with many
// parameters, as the number of functions needed to set all the options can
// become large. Usage in the Golang stdlib The Options pattern is used in the
// Golang standard library for creating objects such as the http.Request object,
// which has many optional parameters. The http.NewRequest function takes a
// method, URL, and optional headers and a body, among other parameters, and
// returns a new http.Request object. The headers and body are optional
// parameters that can be set using functional options.
