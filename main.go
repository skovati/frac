package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"log"
	"math"
	"os"
	"sync"
	"time"

	"github.com/lucasb-eyer/go-colorful"
)

// WaitGroup for concurrent rendering
var wg sync.WaitGroup
var pal []colorful.Color

// vars to hold command line flag values (or defaults)
var (
	width, height    int
	posReal, posImag float64
	fileName         string
	maxIter          int
	size             float64
)

// initialize command line flags
func init() {
	flag.IntVar(&width, "width", 8000, "Output image width")
	flag.IntVar(&height, "height", 6000, "Output image height")
	flag.IntVar(&maxIter, "iter", 255, "Number of mandelbrot iterations per pixel")
	flag.Float64Var(&posReal, "xpos", -0.6, "Real value of centered point")
	flag.Float64Var(&posImag, "ypos", 0.0, "Imaginary value of centered point")
	flag.Float64Var(&size, "size", 2.3, "Size of area in complex plane to render")
	flag.StringVar(&fileName, "output", "frac.png", "File name to store rendered image in")
	// and parse them
	flag.Parse()
}

func render(done chan struct{}) *image.RGBA {
	ratio := float64(width) / float64(height)
	minReal, maxReal := posReal-size*ratio/2.0, math.Abs(posReal+size*ratio/2.0)
	minImag, maxImag := posImag-size/2.0, math.Abs(posImag+size/2.0)
	fmt.Println("x", minReal, maxReal)
	fmt.Println("y", minImag, maxImag)

	image := image.NewRGBA(image.Rect(0, 0, width, height))

	fmt.Printf("%s", "Rendering fractal")
	for row := 0; row < height; row++ {
		wg.Add(1)
		go func(row int) {
			defer wg.Done()
			cImag := minImag + (maxImag-minImag)*float64(row)/float64(height-1)
			for col := 0; col < width; col++ {
				cReal := minReal + (maxReal-minReal)*float64(col)/float64(width-1)
				iter := iterate(complex(cReal, cImag))
				image.Set(col, row, colorize(greyscale, iter))
			}
		}(row)
	}
	wg.Wait()
	done <- struct{}{}
	return image
}

func must(s string, err error) {
	if err != nil {
		log.Printf("%v\n", err)
	}
}

func main() {
	done := make(chan struct{})
	ticker := time.NewTicker(time.Millisecond * 100)

	// print "." if still rendering, otherwise break
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Print(".")
			case <-done:
				ticker.Stop()
				fmt.Printf("\n\nFractal rendered into \"%s\"\n", fileName)
				return
			}
		}
	}()

	// render fractal
	image := render(done)

	// get file path, and save image
	pwd, err := os.Getwd()
	must("os.Getwd", err)
	f, err := os.Create(pwd + string(os.PathSeparator) + fileName)
	must("os.Create", err)
	defer f.Close()
	png.Encode(f, image)
}
