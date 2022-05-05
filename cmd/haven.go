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
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var imgUrl string

// wallhaven API parameters
var q string             // search term
var categories string    // 100, 110, 111 (general, anime, people)
var purity string        // 100, 110, 111 (sfw/sketchy/nsfw)
var minResolution string // 1920x1080
var apiKey string        // users api key
var havenDir = os.Getenv("WP_HAVEN_DIR")
var CACHE = os.Getenv("XDG_CACHE_HOME")

// havenCmd represents the haven command
var havenCmd = &cobra.Command{
	Use:   "haven",
	Short: "Download a random wallpaper from wallhaven",
	Long: `haven - pull a random wallpaper from wallhaven.cc and set it as your wallpaper
	`,
	Run: func(cmd *cobra.Command, args []string) {

		// set defaults for categories and purity
		if categories == "" {
			categories = "100"
		}

		if purity == "" {
			purity = "100"
		}

		if havenDir == "" {
			havenDir = CACHE + "/wallhaven"
		}

		// construct url using flags/defaults
		url := buildApiCall()

		checkApiKey()

		response, err := http.Get(url)
		check(err)
		defer response.Body.Close()

		// check response
		if response.StatusCode != 200 {
			log.Fatal("wallhaven not responding")
			os.Exit(2)
		}

		// read, hopefully json, data into body
		body, err := io.ReadAll(response.Body)
		if !json.Valid([]byte(body)) {
			log.Fatal("not valid json")
			os.Exit(2)
		}

		// parse json
		var data havenData

		err = json.Unmarshal([]byte(body), &data)
		check(err)
		imgUrl = data.Data[0].Path

		// download image
		response, err = http.Get(imgUrl)
		check(err)
		defer response.Body.Close()

		// create file to store image
		imgName := buildImgName(data)
		saveImg, err := os.Create(havenDir + "/" + imgName)
		check(err)
		defer saveImg.Close()

		// copy image to file
		_, err = io.Copy(saveImg, response.Body)
		check(err)

		// set wallpaper
		setWallpaper(havenDir + "/" + imgName)

	},
}

func init() {
	rootCmd.AddCommand(havenCmd)

	// TODO: implement flags for options

	havenCmd.Flags().StringVarP(&categories, "categories", "c", os.Getenv("WP_HAVEN_CAT"), "Set categories to include (1) or exclude (0) (general, anime, people)")
	havenCmd.Flags().StringVarP(&purity, "purity", "p", os.Getenv("WP_HAVEN_PUR"), "Set purity to include (1) or exclude (0) (sfw, sketchy, nsfw)")
	havenCmd.Flags().StringVarP(&q, "search", "s", "", "Search term to look for images")
	havenCmd.Flags().StringVarP(&minResolution, "resolution", "r", "3840x2160", "Minimum resolution [width]x[height]")
	havenCmd.Flags().StringVarP(&apiKey, "apikey", "a", os.Getenv("WP_HAVEN_API"), "Your wallhaven api key, required for nsfw imgaes")

}

func buildImgName(data havenData) string {
	imgNameStr := strings.Builder{}
	imgNameStr.WriteString("wallhaven-")
	imgNameStr.WriteString(data.Data[0].ID)

	imgType := data.Data[0].FileType

	switch {
	case imgType == "image/jpeg":
		imgNameStr.WriteString(".jpg")
	case imgType == "image/png":
		imgNameStr.WriteString(".png")
	default:
		imgNameStr.WriteString(".jpg")
	}

	return imgNameStr.String()
}

func buildApiCall() string {

	// "https://wallhaven.cc/api/v1/search?sorting=random&purity=110&q=space"
	// get all options
	str := strings.Builder{}
	str.WriteString("https://wallhaven.cc/api/v1/search?sorting=random")
	str.WriteString("&purity=")
	str.WriteString(purity)
	str.WriteString("&categories=")
	str.WriteString(categories)
	str.WriteString("&atleast=")
	str.WriteString(minResolution)
	str.WriteString("&q=")
	str.WriteString(q)

	if len(apiKey) > 1 {
		str.WriteString("&apikey=")
		str.WriteString(apiKey)
	}

	return str.String()
}

func checkApiKey() {
	if apiKey == "" && purity[2] == '1' {
		log.Fatal("API key incorrect or not provided. NSFW images requires an API key")
		os.Exit(2)
	}
}

type havenData struct {
	Data []struct {
		ID         string   `json:"id"`
		URL        string   `json:"url"`
		ShortURL   string   `json:"short_url"`
		Views      int      `json:"views"`
		Favorites  int      `json:"favorites"`
		Source     string   `json:"source"`
		Purity     string   `json:"purity"`
		Category   string   `json:"category"`
		DimensionX int      `json:"dimension_x"`
		DimensionY int      `json:"dimension_y"`
		Resolution string   `json:"resolution"`
		Ratio      string   `json:"ratio"`
		FileSize   int      `json:"file_size"`
		FileType   string   `json:"file_type"`
		CreatedAt  string   `json:"created_at"`
		Colors     []string `json:"colors"`
		Path       string   `json:"path"`
		Thumbs     struct {
			Large    string `json:"large"`
			Original string `json:"original"`
			Small    string `json:"small"`
		} `json:"thumbs"`
	} `json:"data"`
	Meta struct {
		CurrentPage int    `json:"current_page"`
		LastPage    int    `json:"last_page"`
		PerPages    string `json:"per_page"` // wallhaven api returns different values here depending on using apikey or not
		PerPagei    int    `json:"per_page"`
		Total       int    `json:"total"`
		Query       string `json:"query"`
		Seed        string `json:"seed"`
	} `json:"meta"`
}
