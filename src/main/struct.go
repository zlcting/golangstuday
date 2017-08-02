package main

import "fmt"

type human struct {
	sex int
}

type per struct {
	human
	age int
	sex int
}

func main() {
	a := per{}
	a.age = 15
	a.sex = 1
	a.human.sex = 2
	fmt.Println(a)
}
