package post

import (
	"golang-exercises/internal/integration/jsonplaceholder"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struct ini digunakan sebagai handler controller untuk endpoint yang berhubungan dengan Post.
// Di sini Handler menyimpan dependency:
// Artinya handler ini membutuhkan JSONPlaceholder client untuk memanggil external API.
// Kenapa pakai pointer *?
// - menghindari copy struct
// - mengikuti dependency injection pattern
// - bisa reuse client yang sama
type Handler struct {
	client *jsonplaceholder.Client
}

// Kenapa pakai constructor? Agar dependency bisa di-inject dari luar.
// Kenapa return pointer? tidak copy struct, konsisten dengan pattern Go service/handler
func NewHandler(client *jsonplaceholder.Client) *Handler {

	// Kenapa pakai &Handler? & berarti mengambil alamat memory struct Handler.
	// Jadi yang dikembalikan adalah: pointer ke Handler
	return &Handler{
		client: client,
	}
}

// (h *Handler) disebut receiver. Artinya function ini milik struct Handler.
// Kenapa receiver pointer? menghindari copy struct, bisa mengakses dependency seperti client
// gin.Context adalah object dari framework Gin yang berisi: request, response etc.
func (h *Handler) GetPosts(c *gin.Context) {
	// instance dari DTO dalam bentuk slice.
	var posts []Post

	// Kenapa pakai &posts? Karena function Get membutuhkan pointer untuk mengisi data.
	// &posts = pointer ke instance dto post.
	err := h.client.Get("/posts", &posts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, posts)
}
