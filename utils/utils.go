package utils

import (
	"fmt"
	"image"
	"io"
	"os"

	"github.com/nfnt/resize"
	"golang.org/x/term"
)

func panic(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func ImgDimmension(file io.Reader) (int, int) {
	imgConfig, _, err := image.DecodeConfig(file)
	panic(err)
	return imgConfig.Width, imgConfig.Height
}

func ImgTo2dMatrix(img image.Image) Pixels {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	var matrix Pixels = make([][][3]uint32, height)

	for i := range matrix {
		matrix[i] = make([][3]uint32, width)
	}

	for i := range height {
		for j := range width {
			// here our coordingates are inverted ig,thats why i'm considering x=j and y=i
			r, g, b, _ := img.At(j, i).RGBA()
			matrix[i][j] = [3]uint32{r / 257, g / 257, b / 257}
		}
	}

	return matrix
}

func BrightnessLuminosity(rgb [3]uint32) int {
	return int(0.21*float64(rgb[0]) + 0.72*float64(rgb[1]) + 0.07*float64(rgb[2]))
}

func BrightnessAverageValue(rgb [3]uint32) int {
	var sum uint32 = 0
	for _, i := range rgb {
		sum += i
	}
	return int(sum / 3)
}

func BrightnessMinMax(rgb [3]uint32) int {
	return int((max(rgb[0], rgb[1], rgb[2]) + min(rgb[0], rgb[1], rgb[2])) / 2)
}

func brightnessMatrix(img *ImgMatrix, brightnessCalc BrightnessCalcAlgo) [][]int {
	if brightnessCalc == nil {
		brightnessCalc = BrightnessLuminosity
	}
	matrix := make([][]int, img.height)

	for i := range img.height {
		matrix[i] = make([]int, img.width)
	}
	for i := range img.height {
		for j := range img.width {
			matrix[i][j] = brightnessCalc(img.Pixels[i][j])
		}
	}
	return matrix
}

func getNormalized(value, minOut, maxOut, minIn, maxIn int) int {
	return ((maxOut - minOut) * (value - minIn) / (maxIn - minIn)) + minOut
}

func MapBrightnessToChar(b int) rune {
	asciiString := "`^\",:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"

	asciiRunes := []rune(asciiString)

	normalizedIdx := getNormalized(b, 0, len(asciiString)-1, 0, 255)
	if normalizedIdx == len(asciiString)-1 {
		return '\x00'
	}
	return asciiRunes[normalizedIdx]
}

func getTerminalSize(fd int) (int, int) {
	width, height, err := term.GetSize(fd)
	panic(err)
	return width, height
}

func resizeImage(img image.Image, newWidth, newHeight int) image.Image {
	return resize.Resize(uint(newWidth), uint(newHeight), img, resize.NearestNeighbor)
}
