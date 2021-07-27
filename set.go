package main

import (
	"math/cmplx"
)

var two complex128 = 2 + 0i
var maxIter = 80

func isElement(c complex128) bool {
    z := complex(0, 0) 
    for i := 0; i < maxIter; i++ {
        if cmplx.Abs(z) > 2 {
            return false
        }
        z = iter(z, c)
    }
    return true
}

func iter(z, c complex128) complex128 {
    return cmplx.Pow(z, two) + c
}
