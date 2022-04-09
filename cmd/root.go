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
	"log"
	"os"

	"github.com/spf13/cobra"
)

var WALL_PATH = getWallPath("WP_WALL_PATH")
var DIR = getDirPath("WP_DIR")

// TODO: add nitrogen as alternative
var ENGINE = "feh"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wp",
	Short: "wallpaper - Select, at will or random, an image to use as a wallpaper",
	Long: `wallpaper - Given an image or directory, wp will set that image or a
random image from the given directory and set is as the background. 
	
	wp
	wp random /some/directory/of/images
	wp set /some/image/file.png`,

	//Run: func(cmd *cobra.Command, args []string) { },

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var dir string

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.wp.yaml)")
	// rootCmd.PersistentFlags().StringVarP(&dir, "dir", "d", "~/Images", "Directory of images")
}

// TOTALLY FLAWLESS ERROR HANDLING
func check(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(2)
	}
}

func getWallPath(key string) string {
	val, e := os.LookupEnv(key)
	if !e {
		return os.ExpandEnv("$HOME/.config/wallpaper")
	} else {
		return val
	}
}

func getDirPath(key string) string {
	val, e := os.LookupEnv(key)
	if !e {
		return os.ExpandEnv("$HOME/Images")
	} else {
		return val
	}
}
