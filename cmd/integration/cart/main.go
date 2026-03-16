package main

import (
	"golang-exercises/internal/integration"
	"golang-exercises/internal/integration/cart"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	dummyClient := integration.NewDummyClient()
	cartHandler := &cart.CartHandler{Client: dummyClient}

	r.GET("/api/top-spenders", cartHandler.GetTopSpenders)
	r.GET("/api/users/:userId/spending", cartHandler.GetUserSpendingDetail)

	r.Run(":8080")
}
