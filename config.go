package main

import (
    "errors"
)

type Config struct {
    width       int
    height      int
    minReal     float64
    maxReal     float64
    minImag     float64
    maxImag     float64
}

func toReal(c Config, x int) (float64, error) {
    if x >= c.width || x < 0 {
        return 0, errors.New("X OOB")
    }
    size := ((c.maxReal - c.minReal) / float64(c.width-1))
    return c.minReal + float64(x)*size, nil
}

func toImag(c Config, y int) (float64, error) {
  if y >= c.height || y < 0 {
    return 0, errors.New("Y OOB")
  }
  size := ((c.maxImag - c.minImag) / float64(c.height-1))

  return c.maxImag - float64(y)*size, nil
}
