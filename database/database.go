package database

import (
	"boilerplate/models"
	"fmt"
	"sync"
)

var (
	db []*models.Image
	mu sync.Mutex
)

// Connect with database
func Connect() {
	db = make([]*models.Image, 0)
	fmt.Println("Connected with Database")
}

func Insert(image *models.Image) {
	mu.Lock()
	db = append(db, image)
	mu.Unlock()
}

func Get() []*models.Image {
	return db
}
