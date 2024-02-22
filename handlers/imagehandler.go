package handlers

import (
	"boilerplate/database"
	"boilerplate/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func HueEdit(c *fiber.Ctx) error {

	image := &models.Image{
		// Note: when writing to external database,
		// we can simply use - Name: c.FormValue("image")
		Name: utils.CopyString(c.FormValue("image")),
		Url:  utils.CopyString(c.FormValue("image")),
		Hue:  utils.CopyString(c.FormValue("hue")),
	}
	database.Insert(image)

	filetype := FileDownloader(image.Url)
	switch filetype {
	case "image/jpeg", "image/jpg":
		return c.JSON(fiber.Map{
			"success": true,
			"image":   image,
		})
	default:
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"success":         false,
			"responseMessage": "image file not supported",
		})
	}
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

func FileDownloader(url string) string {
	path := url
	fname := filepath.Base(path)

	f, err := os.Create("./data/images/inputs/" + fname)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	_, err = f.ReadFrom(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("image downloaded")

	file, err := os.Open("./data/images/inputs/" + fname)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	buff := make([]byte, 512) // why 512 bytes ? see http://golang.org/pkg/net/http/#DetectContentType
	_, err = file.Read(buff)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	filetype := http.DetectContentType(buff)

	fmt.Println(filetype)

	return filetype
}
