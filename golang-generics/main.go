package main

type Person[T any] struct {
	Name       string
	Age        int
	serializeT T
}

func (p *Person[T]) Bark() *T {}

func main() {
	person := Person{
		Name: "John",
		Age:  25,
	}
}
