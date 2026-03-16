package cart

type Cart struct {
	ID     int     `json:"id"`
	UserId int     `json:"userId"`
	Total  float64 `json:"total"`
}

type CartResponse struct {
	Carts []Cart `json:"carts"`
}

type TopSpender struct {
	UserId     int     `json:"userId"`
	TotalSpent float64 `json:"totalSpent"`
}
