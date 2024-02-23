package handlers

import (
	"boilerplate/database"
	"boilerplate/models"
	"fmt"
	"image"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

// Constants for supported image file types
const (
	ImageJPEG = "image/jpeg"
	ImageJPG  = "image/jpg"
)

// NotFound returns a custom 404 page.
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}

// DownloadImageAndHueEdit downloads an image, adjusts its hue, and returns the edited image.
func DownloadImageAndHueEdit(c *fiber.Ctx) error {
	// Get image URL and hue value from the request
	imageURL := utils.CopyString(c.FormValue("image"))
	hue := utils.CopyString(c.FormValue("hue"))

	// Extract filename from the image URL
	filename := filepath.Base(imageURL)

	// Create a new image object
	image := &models.Image{
		Name:   imageURL, // Name and URL can be the same if writing to an external database
		Url:    imageURL,
		Hue:    hue,
		Result: hue + filename,
	}

	// Insert the image into the database
	database.Insert(image)

	// Download the image
	fileType := FileDownloader(image.Url)

	// Check if the file type is supported
	switch fileType {
	case ImageJPEG, ImageJPG:
		// Process hue edit on the image
		err := ProcessHueEdit(image.Url, image.Hue, image.Result)
		if err != nil {
			// Return an error response if editing the image fails
			return c.Status(fiber.ErrUnprocessableEntity.Code).JSON(fiber.Map{
				"success":         false,
				"responseMessage": "error on editing image",
			})
		}
		// Return a success response with the edited image
		return c.JSON(fiber.Map{
			"success": true,
			"image":   image,
		})
	default:
		// Return an error response if the image file type is not supported
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"success":         false,
			"responseMessage": "image file not supported",
		})
	}
}

// ImageList returns a list of images.
func ImageList(c *fiber.Ctx) error {
	// Get the list of images from the database
	images := database.Get()

	// Return the list of images as JSON response
	return c.JSON(fiber.Map{
		"success": true,
		"images":  images,
	})
}

// FileDownloader downloads a file from the given URL and returns its content type.
func FileDownloader(url string) string {
	// Extract filename from the URL
	filename := filepath.Base(url)

	// Create a new file to save the downloaded image
	f, err := os.Create("./data/images/inputs/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Download the image from the URL
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Write the image content to the file
	_, err = f.ReadFrom(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print a message indicating that the image has been downloaded
	fmt.Println("Image downloaded")

	// Open the downloaded file
	file, err := os.Open("./data/images/inputs/" + filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Read a small portion of the file to determine its content type
	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Detect the content type of the file
	fileType := http.DetectContentType(buff)
	fmt.Println(fileType)

	// Return the content type of the file
	return fileType
}

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
