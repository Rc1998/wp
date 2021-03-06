/*
Copyright © 2022 Ryan Campbell <ryan@rcampbell.xyz>

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
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/h2non/filetype"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a given image as a wallpaper",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		// ensure file is an image
		if !validateImg(args[0]) {
			fmt.Printf("Error: %s is not a valid image\n", args[0])
			os.Exit(1)
		}

		setWallpaper(args[0])

	},
}

func setWallpaper(imgName string) {
	// copy image to $XDG_CONFIG_HOME/wallpaper
	srcImg, e := os.Open(imgName) // new image to use
	check(e)
	wallFile, e := os.Create(WALL_PATH) // destination
	check(e)
	defer srcImg.Close()
	defer wallFile.Close()

	_, e = io.Copy(wallFile, srcImg)
	check(e)

	e = wallFile.Sync()
	check(e)

	// set as wallpaper
	err := exec.Command("feh", "--no-fehbg", "--bg-fill", WALL_PATH).Run()
	if err != nil {
		log.Fatal(err)
	}
	message := "Image " + imgName + " set as wallpaper."
	err = exec.Command("notify-send", message).Run()
}

func validateImg(fileName string) bool {

	file, e := os.Open(fileName)
	check(e)
	defer file.Close()

	// check for valid image
	head := make([]byte, 261)
	file.Read(head)

	return filetype.IsImage(head)

}

func init() {
	rootCmd.AddCommand(setCmd)
}
