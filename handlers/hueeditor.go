package handlers

import (
	"fmt"
	"image"
	"os"
	"path/filepath"
	"strconv"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/imgio"
)

// ProcessHueEdit adjusts the hue of an image and saves the edited image.
func ProcessHueEdit(path string, hue string, result string) error {
	// Extract filename from the path
	filename := filepath.Base(path)

	// Open the original image
	img, err := imgio.Open("./data/images/inputs/" + filename)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Adjust the hue of the image
	var edited *image.RGBA
	if i, err := strconv.Atoi(hue); err == nil {
		edited = adjust.Hue(img, i)
	}

	// Remove the previously edited image, if any
	err = os.Remove("./data/images/outputs/" + filename)
	if err != nil {
		fmt.Println(err)
	}

	// Save the edited image
	if err := imgio.Save("./data/images/outputs/"+result, edited, imgio.JPEGEncoder(90)); err != nil {
		fmt.Println(err)
		return err
	}

	// Remove the original image
	err = os.Remove("./data/images/inputs/" + filename)

	// Return any error that occurred during the process
	return err
}
