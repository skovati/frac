package main

import (
	"math/cmplx"
)

var two complex128 = 2 + 0i

func iterate(c complex128) int {
	z := complex(0, 0)
	for i := 0; i < maxIter; i++ {
		if cmplx.Abs(z) > 2 {
			return i
		}
		z = cmplx.Pow(z, two) + c
	}
	return maxIter
}
