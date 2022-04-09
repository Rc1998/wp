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
	"os"

	"github.com/spf13/cobra"
)

var directory string
var wallPath string

// setenvCmd represents the setenv command
var setenvCmd = &cobra.Command{
	Use:   "setenv",
	Short: "Set env values for settings",
	Long: `wp setenv - Can be set via shell or using this command
Sets the individual env values for each configurable setting.
	wp setenv -d [directory]		set location for random wallpaper selection
	wp setend -w [wallpaper path]  set location to save wallpaper`,

	Run: func(cmd *cobra.Command, args []string) {
		if directory != "" {
			os.Setenv("WP_DIR", directory)
		}

		if wallPath != "" {
			os.Setenv("WP_WALL_PATH", wallPath)
		}
	},
}

func init() {
	rootCmd.AddCommand(setenvCmd)

	setenvCmd.Flags().StringVarP(&directory, "directory", "d", "", "Dir to use for random wallpaper selection")
	setenvCmd.Flags().StringVarP(&wallPath, "wallpath", "w", "", "Place to save wallpaper file")
}
