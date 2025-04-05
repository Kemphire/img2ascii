package utils

import (
	"image"
	"io"
)

const (
	RED   string = "\033[31m"
	GREEN string = "\033[32m"
	BLUE  string = "\033[34m"
	RESET string = "\033[0m"
)

type ImgMatrix struct {
	width, height int
	Pixels        [][][3]uint32
}

type ImageRepr struct {
	imgReader io.Reader
	imgStruct image.Image
}

// type ColorMatrix struct {
// 	width, height int
// 	matrix        [][]string
// }

type Pixels [][][3]uint32

type BrightnessCalcAlgo func([3]uint32) int
