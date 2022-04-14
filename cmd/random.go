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
	"io/fs"
	"math/rand"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var fileName string
var fileList = []string{}
var recursion bool

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Picks a random image in a directory to use as the wallpaper",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {

		// use either provided directory or configured directory
		if len(args) < 1 {
			e := filepath.WalkDir(DIR, walk)
			check(e)
		} else {
			e := filepath.WalkDir(args[0], walk)
			check(e)
		}

		// walk through the given directory recursivly to select a wallpaper

		selection := rand.Intn(len(fileList))

		for !validateImg(fileList[selection]) {
			selection = rand.Intn(len(fileList))
		}

		setWallpaper(fileList[selection])

	},
}

func walk(s string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if !d.IsDir() {
		fileList = append(fileList, s)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(randomCmd)
	rand.Seed(time.Now().UnixNano())
	// TODO: implement
	randomCmd.Flags().BoolVarP(&recursion, "recursive", "r", false, "Search every subdirectory for images.")
}
