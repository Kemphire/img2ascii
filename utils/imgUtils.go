package utils

import (
	"fmt"
	"image"
	"io"
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

func ConstructImg(reader io.Reader) (*ImageRepr, error) {
	i := ImageRepr{}
	i.imgReader = reader
	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Println("Error happened in Constructing Img")
		return nil, err
	}
	i.imgStruct = img
	return &i, nil
}
