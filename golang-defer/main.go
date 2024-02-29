package main

import (
	"errors"
	"fmt"
)

func main() {
	// var i int

	// defer func() {
	// 	fmt.Println("defer i is ", i)
	// }()

	// i = rand.Intn(10)
	// n, _ := DoSomething(i)
	// fmt.Println("i is ", n)

	i := 0
	defer fmt.Println(i)
	i++
}

func DoSomething(i int) (int, error) {
	if i%2 == 0 {
		return i, errors.New("i is even")
	} else {
		return i, nil
	}
}
