package main

import "fmt"

type Key struct {
	ID   int
	Pass string
}

type Value struct {
	Name *string
	Age  *int
	Sex  *string
}

var complexMap map[Key]Value = map[Key]Value{
	{ID: 1, Pass: "pass1"}: {Name: Ptr("John, Doe"), Age: Ptr(30)},
	{ID: 2, Pass: "pass2"}: {Name: Ptr("Jane, Doe"), Age: Ptr(28)},
}

func main() {
	v, _ := complexMap[Key{ID: 1, Pass: "pass1"}]
	v.Sex = Ptr("m")

	// Sexフィールドが更新される
	fmt.Println(v)
	// Sexフィールドは更新されていない
	fmt.Println(complexMap[Key{ID: 1, Pass: "pass1"}])
}

func Ptr[T any](e T) *T {
	return &e
}
