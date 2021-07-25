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

func (c Config) toReal(x int) (float64, error) {
    if x >= c.width || x < 0 {
        return 0, errors.New("X OOB")
    }
    size := ((c.maxReal - c.minReal) / float64(c.width-1))
    return c.minReal + float64(x)*size, nil
}

func (c Config) toImag(y int) (float64, error) {
  if y >= c.height || y < 0 {
    return 0, errors.New("Y is out of bounds")
  }
  size := ((c.maxImag - c.minImag) / float64(c.height-1))

  return c.maxImag - float64(y)*size, nil
}
