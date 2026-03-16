package main

import (
	"golang-exercises/internal/integration/jsonplaceholder"
	"golang-exercises/internal/integration/jsonplaceholder/post"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	client := jsonplaceholder.NewClient()
	postHandler := post.NewHandler(client)

	r.GET("/posts", postHandler.GetPosts)
	r.Run(":8080")
}
