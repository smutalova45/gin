package models

type Product struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Price         int    `json:"price"`
	OriginalPrice int    `json:"original_price"`
	Quantity      int    `json:"quantity"`
	CategoryID    string `json:"category_id"`
}

type CreateProduct struct {
	Name          string `json:"name"`
	Price         int    `json:"price"`
	OriginalPrice int    `json:"original_price"`
	Quantity      int    `json:"quantity"`
	CategoryID    string `json:"category_id"`
}

type UpdateProduct struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Price         int    `json:"price"`
	OriginalPrice int    `json:"original_price"`
	Quantity      int    `json:"quantity"`
	CategoryID    string `json:"category_id"`
}

type ProductSell struct {
	Name string `json:"name"`
}
type ProductSellRequest struct {
	Name     string    `json:"name"`
	Price    int       `json:"-"`
	Quantity int       `json:"quantity"`
	// Product  []Product `json:"product"`
}
type SellRequest struct {
	Products map[string]int `json:"products"`
	BasketID string         `json:"basket_id"`
}
type ProductResponse struct {
	Product []Product `json:"product"`
	Count   int       `json:"count"`
}
