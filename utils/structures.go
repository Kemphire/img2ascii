package utils

import (
	"image"
	"io"
)

type ImgMatrix struct {
	width, height int
	Pixels        [][][3]uint32
}

type ImageRepr struct {
	imgReader io.Reader
	imgStruct image.Image
}

type Pixels [][][3]uint32

type BrightnessCalcAlgo func([3]uint32) int
