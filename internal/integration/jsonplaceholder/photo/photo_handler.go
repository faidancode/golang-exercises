package photo

import (
	"golang-exercises/internal/integration/jsonplaceholder"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	client *jsonplaceholder.SecondClient
}

func NewHandler(client *jsonplaceholder.SecondClient) *Handler {
	return &Handler{
		client: client,
	}
}

func (h *Handler) GetPhotos(c *gin.Context) {
	var photos []Photo

	err := h.client.Get("/photos", &photos)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, photos)
}
