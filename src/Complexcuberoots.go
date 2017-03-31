package main

import "fmt"
import "math/cmplx"

func Cbrt(x complex128) complex128 {
	z := complex128(1)
	for {
		if y := z - (cmplx.Pow(z, 3)-x)/(3*z*z); cmplx.Abs(y-z) < 1e-10 {
			return y
		} else {
			return z
		}

	}
}

func main() {
	fmt.Println(Cbrt(2))
}
