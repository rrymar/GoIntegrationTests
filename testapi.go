package main

import (
	"testapi/albums"
)

func main() {
	router := albums.SetupRoutes()
	router.Run("localhost:8080")
}
