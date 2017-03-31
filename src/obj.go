package main

import "fmt"

type Integer int

func (a *Integer) Add(b Integer) {
	fmt.Println("a = ", *a)
	*a += b
}

func main() {
	var a Integer = 1
	a.Add(2)
	fmt.Println("a = ", a)

}
