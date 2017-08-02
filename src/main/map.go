package main

import "fmt"

func main() {
	m1 := map[int]string{1: "a", 2: "b", 3: "c"}
	m2 := map[string]int{}
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m2)
}
