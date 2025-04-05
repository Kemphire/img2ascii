package utils

import (
	"fmt"
	"os"
)

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

	// fmt.Println(brightnessMatrix)

	for i := range height {
		for j := range width {
			fmt.Printf("%c", MapBrightnessToChar(brightnessMatrix[i][j]))
		}
		fmt.Println()
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
		fmt.Println()
	}
}

func (i ImageRepr) PrintImageColored(brightnessCalc BrightnessCalcAlgo) {
	width := i.width()
	height := i.height()

	if tWidth, tHeight := getTerminalSize(int(os.Stdin.Fd())); width > tWidth || height > tHeight {
		if width > tWidth {
			width = tWidth - (tWidth / 10)
		}
		if height > tHeight {
			height = tHeight - (tHeight / 10)
		}
		i.imgStruct = resizeImage(i.imgStruct, width, height)
	}

	coloredStringMatrix := make([][]string, height)

	for i := range height {
		coloredStringMatrix[i] = make([]string, width)
	}

	brightnessMatrix := i.getInverseBrighnessMatrix(brightnessCalc)
	colorMatrix := i.get2dmatrix()

	for i := range height {
		for j := range width {
			red := colorMatrix[i][j][0]
			green := colorMatrix[i][j][1]
			blue := colorMatrix[i][j][2]
			coloredStringMatrix[i][j] = fmt.Sprintf("\033[38;2;%d;%d;%dm%c%s", red, green, blue, MapBrightnessToChar(brightnessMatrix[i][j]), RESET)
		}
	}

	for i := range height {
		for j := range width {
			fmt.Printf("%s", coloredStringMatrix[i][j])
		}
		fmt.Println()
	}
}
