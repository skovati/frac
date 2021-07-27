package main

import (
    "fmt"
)

func main() {
    c := Config{151, 51, -2.0, 0.5, -1.0, 1.0}
    for y := 0; y < c.height; y++ {
        cImag, err := toImag(c, y)
        if err != nil {
            fmt.Errorf("%v", err)
        }
        for x := 0; x < c.width; x++ {
            cReal, err := toReal(c, x)
            if err != nil {
                fmt.Errorf("%v", err)
            }
            if isElement(complex(cReal, cImag)) {
                fmt.Print("*")
            } else {
                fmt.Print(".")
            }
        }
        fmt.Println()
    }
}
