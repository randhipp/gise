package handlers

import (
	"boilerplate/database"
	"boilerplate/models"
	"path/filepath"

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
