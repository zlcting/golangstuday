package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	const E = 0.000001
	z := float64(1) //注意float型
	k := float64(0)
	for ; ; z = z - (z*z-x)/(2*z) {
		if z-k <= E && z-k >= -E {
			return z
		}
		k = z
	}
}

func main() {
	fmt.Println(Sqrt(10))
	//fmt.Println(E)

}
