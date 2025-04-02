package utils

import (
	"fmt"
	"image"
	"io"
	"os"
)

func (i ImageRepr) width() int {
	return i.imgStruct.Bounds().Dx()
}

func (i ImageRepr) height() int {
	return i.imgStruct.Bounds().Dy()
}

func (i ImageRepr) get2dmatrix() Pixels {
	return ImgTo2dMatrix(i.imgStruct)
}

func (i ImageRepr) getBrightnessMatrix(brightnessCalc BrightnessCalcAlgo) [][]int {
	if brightnessCalc == nil {
		brightnessCalc = BrightnessLuminosity
	}
	width := i.width()
	height := i.height()

	rbgMatrix := i.get2dmatrix()

	matrix := make([][]int, height)

	for i := range height {
		matrix[i] = make([]int, width)
	}

	for i := range height {
		for j := range width {
			matrix[i][j] = brightnessCalc(rbgMatrix[i][j])
		}
	}
	return matrix
}

func invertBrighness(value int) int {
	return 255 - value
}

func (i ImageRepr) getInverseBrighnessMatrix(brightnessCalc BrightnessCalcAlgo) [][]int {
	if brightnessCalc == nil {
		brightnessCalc = BrightnessLuminosity
	}
	width := i.width()
	height := i.height()

	rbgMatrix := i.get2dmatrix()

	matrix := make([][]int, height)

	for i := range height {
		matrix[i] = make([]int, width)
	}

	for i := range height {
		for j := range width {
			matrix[i][j] = invertBrighness(brightnessCalc(rbgMatrix[i][j]))
		}
	}
	return matrix

}

func (i ImageRepr) GetImage() image.Image {
	return i.imgStruct
}

func (i ImageRepr) PrintImg(brightnessCalc BrightnessCalcAlgo) {
	width := i.width()
	height := i.height()

	if tWidth, tHeight := getTerminalSize(int(os.Stdin.Fd())); width > tWidth || height > tHeight {
		if width > tWidth {
			width = tWidth
		}
		if height > tHeight {
			height = tHeight
		}
		i.imgStruct = resizeImage(i.imgStruct, width, height)
	}

	brightnessMatrix := i.getBrightnessMatrix(brightnessCalc)

	for i := range height {
		for j := range width {
			fmt.Printf("%c", MapBrightnessToChar(brightnessMatrix[i][j]))
		}

	}
}

func (i ImageRepr) PrintImageIverted(brightnessCalc BrightnessCalcAlgo) {
	width := i.width()
	height := i.height()

	if tWidth, tHeight := getTerminalSize(int(os.Stdin.Fd())); width > tWidth || height > tHeight {
		if width > tWidth {
			width = tWidth
		}
		if height > tHeight {
			height = tHeight
		}
		i.imgStruct = resizeImage(i.imgStruct, width, height)
	}

	brightnessMatrix := i.getInverseBrighnessMatrix(brightnessCalc)

	for i := range height {
		for j := range width {
			fmt.Printf("%c", MapBrightnessToChar(brightnessMatrix[i][j]))
		}
	}
}

func ConstructImg(reader io.Reader) (*ImageRepr, error) {
	i := ImageRepr{}
	i.imgReader = reader
	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Println("Error happened in ConstructImg")
		return nil, err
	}
	i.imgStruct = img
	return &i, nil
}
