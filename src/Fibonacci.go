package main

import "fmt"

func fibonacci() func() int {
	sum1 := 0
	sum2 := 1

	return func() int {
		out := sum1 + sum2
		sum1 = sum2
		sum2 = out

		return out
	}
}

func main() {
	f := fibonacci()

	fmt.Println(f())
	fmt.Println(f())

}
