package cart

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	Client *DummyClient
}

func (h *CartHandler) fetchCarts() ([]Cart, error) {
	var data CartResponse
	err := h.Client.FetchGeneric("/carts", &data)
	if err != nil {
		return nil, err
	}
	return data.Carts, nil
}

func (h *CartHandler) GetTopSpenders(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "5")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 5
	}

	if limit > 20 {
		limit = 20
	}

	carts, err := h.fetchCarts()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	userMap := make(map[int]float64)

	for _, crt := range carts {
		userMap[crt.UserId] += crt.Total
	}

	var spenders []TopSpender
	for id, total := range userMap {
		spenders = append(spenders, TopSpender{UserId: id, TotalSpent: total})
	}

	sort.Slice(spenders, func(i, j int) bool {
		return spenders[i].TotalSpent > spenders[j].TotalSpent
	})

	if len(spenders) > limit {
		spenders = spenders[:limit]
	}

	c.JSON(http.StatusOK, spenders)

}

func (h *CartHandler) GetUserSpendingDetail(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, err := strconv.Atoi(userIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
	}

	carts, err := h.fetchCarts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "External API failure"})
	}

	var totalSpent float64
	var count int
	found := false
	for _, crt := range carts {
		if crt.UserId == userId {
			totalSpent += crt.Total
			count++
			found = true
		}
	}

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found or has no cart "})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userId":     userId,
		"totalCart":  count,
		"totalSpent": totalSpent,
	})
}
