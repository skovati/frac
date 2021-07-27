package main

import (
    "fmt"
    "image"
    "image/color"
    "image/png"
    "os"
)

func main() {
    c := Config{4000, 3000, -2.0, 0.6, -1.1, 1.1}
    image := image.NewRGBA(image.Rect(0, 0, c.width, c.height))
    black := color.RGBA{0, 0, 0, 255}
    white := color.RGBA{255, 255, 255, 255}
    color := white
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
                color = black
            } else {
                color = white
            }
            image.Set(x, y, color)
        }
    }
    f, err := os.Create("/home/skovati/brot.png")
    if err != nil {
        panic(err)
    }
    defer f.Close()
    png.Encode(f, image)
}
