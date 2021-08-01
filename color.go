package main

import (
	"image/color"
)

func colorize(pal []color.RGBA, iter int) color.RGBA {
	return pal[iter%(len(pal)-1)]
}

var greyscale []color.RGBA = []color.RGBA{
	{0x00, 0x00, 0x00, 0xFF},
	{0x10, 0x10, 0x10, 0xFF},
	{0x20, 0x20, 0x20, 0xFF},
	{0x30, 0x30, 0x30, 0xFF},
	{0x40, 0x40, 0x40, 0xFF},
	{0x50, 0x50, 0x50, 0xFF},
	{0x60, 0x60, 0x60, 0xFF},
	{0x70, 0x70, 0x70, 0xFF},
	{0x80, 0x80, 0x80, 0xFF},
	{0x90, 0x90, 0x90, 0xFF},
	{0xA0, 0xA0, 0xA0, 0xFF},
	{0xB0, 0xB0, 0xB0, 0xFF},
	{0xC0, 0xC0, 0xC0, 0xFF},
	{0xD0, 0xD0, 0xD0, 0xFF},
	{0xE0, 0xE0, 0xE0, 0xFF},
	{0xF0, 0xF0, 0xF0, 0xFF},
}

var bobRoss []color.RGBA = []color.RGBA{
	{0x00, 0x00, 0x00, 0xFF}, // Midnight black
	{0x02, 0x1E, 0x44, 0xFF}, // Prussian blue
	{0x0A, 0x34, 0x10, 0xFF}, // Sap green
	{0x0C, 0x00, 0x40, 0xFF}, // Phthalo blue
	{0x10, 0x2E, 0x3C, 0xFF}, // Phthalo green
	{0x22, 0x1B, 0x15, 0xFF}, // Van Dyke brown
	{0x4E, 0x15, 0x00, 0xFF}, // Alizarin crimson
	{0x5F, 0x2E, 0x1F, 0xFF}, // Dark Sienna
	{0xC7, 0x9B, 0x00, 0xFF}, // Yellow ochre
	{0xDB, 0x00, 0x00, 0xFF}, // Bright red
	{0xFF, 0x3C, 0x00, 0xFF}, // Cadmium yellow
	{0xFF, 0xB8, 0x00, 0xFF}, // Indian yellow
	{0xFF, 0xFF, 0xFF, 0xFF}, // Titanium white
}

var gruvbox []color.RGBA = []color.RGBA{
	{0xa8, 0x99, 0x84, 0xFF}, // grey
	{0xcc, 0x24, 0x1d, 0xFF}, // red
	{0x98, 0x97, 0x1a, 0xFF}, // green
	{0xd7, 0x99, 0x21, 0xFF}, // yellow
	{0x45, 0x85, 0x88, 0xFF}, // blue
	{0x68, 0x9D, 0x6A, 0xFF}, // cyan
	{0xB1, 0x62, 0x86, 0xFF}, // magenta
	{0x28, 0x28, 0x28, 0xFF}, // black
}

var base20 []color.RGBA = []color.RGBA{
	{0xE6, 0x19, 0x4B, 0xFF},
	{0x3C, 0xB4, 0x4B, 0xFF},
	{0xFF, 0xE1, 0x19, 0xFF},
	{0x43, 0x63, 0xD8, 0xFF},
	{0xF5, 0x82, 0x31, 0xFF},
	{0x91, 0x1E, 0xB4, 0xFF},
	{0x46, 0xF0, 0xF0, 0xFF},
	{0xF0, 0x32, 0xE6, 0xFF},
	{0xBC, 0xF6, 0x0C, 0xFF},
	{0xFA, 0xBE, 0xBE, 0xFF},
	{0x00, 0x80, 0x80, 0xFF},
	{0xE6, 0xBE, 0xFF, 0xFF},
	{0x9A, 0x63, 0x24, 0xFF},
	{0xFF, 0xFA, 0xC8, 0xFF},
	{0x80, 0x00, 0x00, 0xFF},
	{0xAA, 0xFF, 0xC3, 0xFF},
	{0x80, 0x80, 0x00, 0xFF},
	{0xFF, 0xD8, 0xB1, 0xFF},
	{0x00, 0x00, 0x75, 0xFF},
	{0x80, 0x80, 0x80, 0xFF},
	{0xFF, 0xFF, 0xFF, 0xFF},
	{0x00, 0x00, 0x00, 0xFF},
}

var plan9 []color.RGBA = []color.RGBA{
	{0x00, 0x00, 0x00, 0xFF},
	{0x00, 0x00, 0x44, 0xFF},
	{0x00, 0x00, 0x88, 0xFF},
	{0x00, 0x00, 0xCC, 0xFF},
	{0x00, 0x44, 0x00, 0xFF},
	{0x00, 0x44, 0x44, 0xFF},
	{0x00, 0x44, 0x88, 0xFF},
	{0x00, 0x44, 0xCC, 0xFF},
	{0x00, 0x88, 0x00, 0xFF},
	{0x00, 0x88, 0x44, 0xFF},
	{0x00, 0x88, 0x88, 0xFF},
	{0x00, 0x88, 0xCC, 0xFF},
	{0x00, 0xCC, 0x00, 0xFF},
	{0x00, 0xCC, 0x44, 0xFF},
	{0x00, 0xCC, 0x88, 0xFF},
	{0x00, 0xCC, 0xCC, 0xFF},
	{0x00, 0xDD, 0xDD, 0xFF},
	{0x11, 0x11, 0x11, 0xFF},
	{0x00, 0x00, 0x55, 0xFF},
	{0x00, 0x00, 0x99, 0xFF},
	{0x00, 0x00, 0xDD, 0xFF},
	{0x00, 0x55, 0x00, 0xFF},
	{0x00, 0x55, 0x55, 0xFF},
	{0x00, 0x4C, 0x99, 0xFF},
	{0x00, 0x49, 0xDD, 0xFF},
	{0x00, 0x99, 0x00, 0xFF},
	{0x00, 0x99, 0x4C, 0xFF},
	{0x00, 0x99, 0x99, 0xFF},
	{0x00, 0x93, 0xDD, 0xFF},
	{0x00, 0xDD, 0x00, 0xFF},
	{0x00, 0xDD, 0x49, 0xFF},
	{0x00, 0xDD, 0x93, 0xFF},
	{0x00, 0xEE, 0x9E, 0xFF},
	{0x00, 0xEE, 0xEE, 0xFF},
	{0x22, 0x22, 0x22, 0xFF},
	{0x00, 0x00, 0x66, 0xFF},
	{0x00, 0x00, 0xAA, 0xFF},
	{0x00, 0x00, 0xEE, 0xFF},
	{0x00, 0x66, 0x00, 0xFF},
	{0x00, 0x66, 0x66, 0xFF},
	{0x00, 0x55, 0xAA, 0xFF},
	{0x00, 0x4F, 0xEE, 0xFF},
	{0x00, 0xAA, 0x00, 0xFF},
	{0x00, 0xAA, 0x55, 0xFF},
	{0x00, 0xAA, 0xAA, 0xFF},
	{0x00, 0x9E, 0xEE, 0xFF},
	{0x00, 0xEE, 0x00, 0xFF},
	{0x00, 0xEE, 0x4F, 0xFF},
	{0x00, 0xFF, 0x55, 0xFF},
	{0x00, 0xFF, 0xAA, 0xFF},
	{0x00, 0xFF, 0xFF, 0xFF},
	{0x33, 0x33, 0x33, 0xFF},
	{0x00, 0x00, 0x77, 0xFF},
	{0x00, 0x00, 0xBB, 0xFF},
	{0x00, 0x00, 0xFF, 0xFF},
	{0x00, 0x77, 0x00, 0xFF},
	{0x00, 0x77, 0x77, 0xFF},
	{0x00, 0x5D, 0xBB, 0xFF},
	{0x00, 0x55, 0xFF, 0xFF},
	{0x00, 0xBB, 0x00, 0xFF},
	{0x00, 0xBB, 0x5D, 0xFF},
	{0x00, 0xBB, 0xBB, 0xFF},
	{0x00, 0xAA, 0xFF, 0xFF},
	{0x00, 0xFF, 0x00, 0xFF},
	{0x44, 0x00, 0x44, 0xFF},
	{0x44, 0x00, 0x88, 0xFF},
	{0x44, 0x00, 0xCC, 0xFF},
	{0x44, 0x44, 0x00, 0xFF},
	{0x44, 0x44, 0x44, 0xFF},
	{0x44, 0x44, 0x88, 0xFF},
	{0x44, 0x44, 0xCC, 0xFF},
	{0x44, 0x88, 0x00, 0xFF},
	{0x44, 0x88, 0x44, 0xFF},
	{0x44, 0x88, 0x88, 0xFF},
	{0x44, 0x88, 0xCC, 0xFF},
	{0x44, 0xCC, 0x00, 0xFF},
	{0x44, 0xCC, 0x44, 0xFF},
	{0x44, 0xCC, 0x88, 0xFF},
	{0x44, 0xCC, 0xCC, 0xFF},
	{0x44, 0x00, 0x00, 0xFF},
	{0x55, 0x00, 0x00, 0xFF},
	{0x55, 0x00, 0x55, 0xFF},
	{0x4C, 0x00, 0x99, 0xFF},
	{0x49, 0x00, 0xDD, 0xFF},
	{0x55, 0x55, 0x00, 0xFF},
	{0x55, 0x55, 0x55, 0xFF},
	{0x4C, 0x4C, 0x99, 0xFF},
	{0x49, 0x49, 0xDD, 0xFF},
	{0x4C, 0x99, 0x00, 0xFF},
	{0x4C, 0x99, 0x4C, 0xFF},
	{0x4C, 0x99, 0x99, 0xFF},
	{0x49, 0x93, 0xDD, 0xFF},
	{0x49, 0xDD, 0x00, 0xFF},
	{0x49, 0xDD, 0x49, 0xFF},
	{0x49, 0xDD, 0x93, 0xFF},
	{0x49, 0xDD, 0xDD, 0xFF},
	{0x4F, 0xEE, 0xEE, 0xFF},
	{0x66, 0x00, 0x00, 0xFF},
	{0x66, 0x00, 0x66, 0xFF},
	{0x55, 0x00, 0xAA, 0xFF},
	{0x4F, 0x00, 0xEE, 0xFF},
	{0x66, 0x66, 0x00, 0xFF},
	{0x66, 0x66, 0x66, 0xFF},
	{0x55, 0x55, 0xAA, 0xFF},
	{0x4F, 0x4F, 0xEE, 0xFF},
	{0x55, 0xAA, 0x00, 0xFF},
	{0x55, 0xAA, 0x55, 0xFF},
	{0x55, 0xAA, 0xAA, 0xFF},
	{0x4F, 0x9E, 0xEE, 0xFF},
	{0x4F, 0xEE, 0x00, 0xFF},
	{0x4F, 0xEE, 0x4F, 0xFF},
	{0x4F, 0xEE, 0x9E, 0xFF},
	{0x55, 0xFF, 0xAA, 0xFF},
	{0x55, 0xFF, 0xFF, 0xFF},
	{0x77, 0x00, 0x00, 0xFF},
	{0x77, 0x00, 0x77, 0xFF},
	{0x5D, 0x00, 0xBB, 0xFF},
	{0x55, 0x00, 0xFF, 0xFF},
	{0x77, 0x77, 0x00, 0xFF},
	{0x77, 0x77, 0x77, 0xFF},
	{0x5D, 0x5D, 0xBB, 0xFF},
	{0x55, 0x55, 0xFF, 0xFF},
	{0x5D, 0xBB, 0x00, 0xFF},
	{0x5D, 0xBB, 0x5D, 0xFF},
	{0x5D, 0xBB, 0xBB, 0xFF},
	{0x55, 0xAA, 0xFF, 0xFF},
	{0x55, 0xFF, 0x00, 0xFF},
	{0x55, 0xFF, 0x55, 0xFF},
	{0x88, 0x00, 0x88, 0xFF},
	{0x88, 0x00, 0xCC, 0xFF},
	{0x88, 0x44, 0x00, 0xFF},
	{0x88, 0x44, 0x44, 0xFF},
	{0x88, 0x44, 0x88, 0xFF},
	{0x88, 0x44, 0xCC, 0xFF},
	{0x88, 0x88, 0x00, 0xFF},
	{0x88, 0x88, 0x44, 0xFF},
	{0x88, 0x88, 0x88, 0xFF},
	{0x88, 0x88, 0xCC, 0xFF},
	{0x88, 0xCC, 0x00, 0xFF},
	{0x88, 0xCC, 0x44, 0xFF},
	{0x88, 0xCC, 0x88, 0xFF},
	{0x88, 0xCC, 0xCC, 0xFF},
	{0x88, 0x00, 0x00, 0xFF},
	{0x88, 0x00, 0x44, 0xFF},
	{0x99, 0x00, 0x4C, 0xFF},
	{0x99, 0x00, 0x99, 0xFF},
	{0x93, 0x00, 0xDD, 0xFF},
	{0x99, 0x4C, 0x00, 0xFF},
	{0x99, 0x4C, 0x4C, 0xFF},
	{0x99, 0x4C, 0x99, 0xFF},
	{0x93, 0x49, 0xDD, 0xFF},
	{0x99, 0x99, 0x00, 0xFF},
	{0x99, 0x99, 0x4C, 0xFF},
	{0x99, 0x99, 0x99, 0xFF},
	{0x93, 0x93, 0xDD, 0xFF},
	{0x93, 0xDD, 0x00, 0xFF},
	{0x93, 0xDD, 0x49, 0xFF},
	{0x93, 0xDD, 0x93, 0xFF},
	{0x93, 0xDD, 0xDD, 0xFF},
	{0x99, 0x00, 0x00, 0xFF},
	{0xAA, 0x00, 0x00, 0xFF},
	{0xAA, 0x00, 0x55, 0xFF},
	{0xAA, 0x00, 0xAA, 0xFF},
	{0x9E, 0x00, 0xEE, 0xFF},
	{0xAA, 0x55, 0x00, 0xFF},
	{0xAA, 0x55, 0x55, 0xFF},
	{0xAA, 0x55, 0xAA, 0xFF},
	{0x9E, 0x4F, 0xEE, 0xFF},
	{0xAA, 0xAA, 0x00, 0xFF},
	{0xAA, 0xAA, 0x55, 0xFF},
	{0xAA, 0xAA, 0xAA, 0xFF},
	{0x9E, 0x9E, 0xEE, 0xFF},
	{0x9E, 0xEE, 0x00, 0xFF},
	{0x9E, 0xEE, 0x4F, 0xFF},
	{0x9E, 0xEE, 0x9E, 0xFF},
	{0x9E, 0xEE, 0xEE, 0xFF},
	{0xAA, 0xFF, 0xFF, 0xFF},
	{0xBB, 0x00, 0x00, 0xFF},
	{0xBB, 0x00, 0x5D, 0xFF},
	{0xBB, 0x00, 0xBB, 0xFF},
	{0xAA, 0x00, 0xFF, 0xFF},
	{0xBB, 0x5D, 0x00, 0xFF},
	{0xBB, 0x5D, 0x5D, 0xFF},
	{0xBB, 0x5D, 0xBB, 0xFF},
	{0xAA, 0x55, 0xFF, 0xFF},
	{0xBB, 0xBB, 0x00, 0xFF},
	{0xBB, 0xBB, 0x5D, 0xFF},
	{0xBB, 0xBB, 0xBB, 0xFF},
	{0xAA, 0xAA, 0xFF, 0xFF},
	{0xAA, 0xFF, 0x00, 0xFF},
	{0xAA, 0xFF, 0x55, 0xFF},
	{0xAA, 0xFF, 0xAA, 0xFF},
	{0xCC, 0x00, 0xCC, 0xFF},
	{0xCC, 0x44, 0x00, 0xFF},
	{0xCC, 0x44, 0x44, 0xFF},
	{0xCC, 0x44, 0x88, 0xFF},
	{0xCC, 0x44, 0xCC, 0xFF},
	{0xCC, 0x88, 0x00, 0xFF},
	{0xCC, 0x88, 0x44, 0xFF},
	{0xCC, 0x88, 0x88, 0xFF},
	{0xCC, 0x88, 0xCC, 0xFF},
	{0xCC, 0xCC, 0x00, 0xFF},
	{0xCC, 0xCC, 0x44, 0xFF},
	{0xCC, 0xCC, 0x88, 0xFF},
	{0xCC, 0xCC, 0xCC, 0xFF},
	{0xCC, 0x00, 0x00, 0xFF},
	{0xCC, 0x00, 0x44, 0xFF},
	{0xCC, 0x00, 0x88, 0xFF},
	{0xDD, 0x00, 0x93, 0xFF},
	{0xDD, 0x00, 0xDD, 0xFF},
	{0xDD, 0x49, 0x00, 0xFF},
	{0xDD, 0x49, 0x49, 0xFF},
	{0xDD, 0x49, 0x93, 0xFF},
	{0xDD, 0x49, 0xDD, 0xFF},
	{0xDD, 0x93, 0x00, 0xFF},
	{0xDD, 0x93, 0x49, 0xFF},
	{0xDD, 0x93, 0x93, 0xFF},
	{0xDD, 0x93, 0xDD, 0xFF},
	{0xDD, 0xDD, 0x00, 0xFF},
	{0xDD, 0xDD, 0x49, 0xFF},
	{0xDD, 0xDD, 0x93, 0xFF},
	{0xDD, 0xDD, 0xDD, 0xFF},
	{0xDD, 0x00, 0x00, 0xFF},
	{0xDD, 0x00, 0x49, 0xFF},
	{0xEE, 0x00, 0x4F, 0xFF},
	{0xEE, 0x00, 0x9E, 0xFF},
	{0xEE, 0x00, 0xEE, 0xFF},
	{0xEE, 0x4F, 0x00, 0xFF},
	{0xEE, 0x4F, 0x4F, 0xFF},
	{0xEE, 0x4F, 0x9E, 0xFF},
	{0xEE, 0x4F, 0xEE, 0xFF},
	{0xEE, 0x9E, 0x00, 0xFF},
	{0xEE, 0x9E, 0x4F, 0xFF},
	{0xEE, 0x9E, 0x9E, 0xFF},
	{0xEE, 0x9E, 0xEE, 0xFF},
	{0xEE, 0xEE, 0x00, 0xFF},
	{0xEE, 0xEE, 0x4F, 0xFF},
	{0xEE, 0xEE, 0x9E, 0xFF},
	{0xEE, 0xEE, 0xEE, 0xFF},
	{0xEE, 0x00, 0x00, 0xFF},
	{0xFF, 0x00, 0x00, 0xFF},
	{0xFF, 0x00, 0x55, 0xFF},
	{0xFF, 0x00, 0xAA, 0xFF},
	{0xFF, 0x00, 0xFF, 0xFF},
	{0xFF, 0x55, 0x00, 0xFF},
	{0xFF, 0x55, 0x55, 0xFF},
	{0xFF, 0x55, 0xAA, 0xFF},
	{0xFF, 0x55, 0xFF, 0xFF},
	{0xFF, 0xAA, 0x00, 0xFF},
	{0xFF, 0xAA, 0x55, 0xFF},
	{0xFF, 0xAA, 0xAA, 0xFF},
	{0xFF, 0xAA, 0xFF, 0xFF},
	{0xFF, 0xFF, 0x00, 0xFF},
	{0xFF, 0xFF, 0x55, 0xFF},
	{0xFF, 0xFF, 0xAA, 0xFF},
	{0xFF, 0xFF, 0xFF, 0xFF}}
