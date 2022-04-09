# wp - wallpaper

A simple cli tool to set, either randomly or specifically, a wallpaper using feh as a backend.

Built using Golang and [Cobra](https://pkg.go.dev/github.com/spf13/cobra#pkg-overview)

## Usage

To set a wallpaper, make sure feh is installed and run the following
```bash
$ wp set image.jpg
```

When an image is set, that file is copied to $XDG_CONFIG_HOME/wallpaper

To get a random wallpaper set the enviornment variable WP_DIR or pass a directory as an argument
```bash
$ wp random /some/folder/of/images
```

or with WP_DIR set simply

```bash
$ wp random
```

`wp random` searches the provided directory and any subdirectories for images