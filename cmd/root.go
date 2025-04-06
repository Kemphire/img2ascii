/*
Copyright © 2025 Kartikey Shahi

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/kemphire/terminal_art/utils"
	"github.com/spf13/cobra"
)

var filePath string
var inverse bool
var color bool
var imgDimension []int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "terminal_art",
	Short: "A Image to ascii converter",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if imgDimension != nil && len(imgDimension) != 2 {
			fmt.Println(len(imgDimension))
			return fmt.Errorf("only two vaues are allowed in dimmension/-d flag")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		reader, err := utils.FetchFileUrl(filePath)

		if err != nil {
			fmt.Println(err)
			return
		}
		defer reader.Close()

		img, err := utils.ConstructImg(reader)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(inverse)
		fmt.Println(color)
		fmt.Println(len(imgDimension))

		if inverse && color {
			fmt.Fprintf(os.Stderr, "Colored inverse images are not supported")
			os.Exit(2)
		}

		if inverse {
			img.PrintImageIverted(utils.BrightnessLuminosity)
		} else if color {
			if len(imgDimension) == 2 {
				img.PrintImageColored(utils.BrightnessLuminosity, imgDimension...)
			} else {
				img.PrintImageColored(utils.BrightnessLuminosity)
			}
		} else {
			img.PrintImg(utils.BrightnessLuminosity)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.terminal_art.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVar(&filePath, "file", "", "Image file path or url")
	rootCmd.Flags().BoolVarP(&inverse, "inverse", "i", false, "Toggle brighness inversion")
	rootCmd.Flags().BoolVarP(&color, "color", "c", false, "Toggle color for ascii image")
	rootCmd.Flags().IntSliceVarP(&imgDimension, "dimmension", "d", nil, "Dimmension of the image, in the format widht,height")
	rootCmd.MarkFlagRequired("file")
}
