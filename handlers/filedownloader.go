package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

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
