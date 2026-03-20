package main

import (
	"golang-exercises/internal/integration/jsonplaceholder"
	"golang-exercises/internal/integration/jsonplaceholder/photo"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	client := jsonplaceholder.NewSecondClient()
	photoHandler := photo.NewHandler(client)

	r.GET("/photos", photoHandler.GetPhotos)
	r.Run(":8080")

}
