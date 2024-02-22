package handlers

import (
	"boilerplate/database"
	"boilerplate/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func HueEdit(c *fiber.Ctx) error {

	image := &models.Image{
		// Note: when writing to external database,
		// we can simply use - Name: c.FormValue("image")
		Name: utils.CopyString(c.FormValue("image")),
	}
	database.Insert(image)

	return c.JSON(fiber.Map{
		"success": true,
		"image":   image,
	})
}

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}

func ImageList(c *fiber.Ctx) error {
	images := database.Get()

	return c.JSON(fiber.Map{
		"success": true,
		"images":  images,
	})
}
