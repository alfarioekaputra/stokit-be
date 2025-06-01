package model

type ProductResponse struct {
	ID               string  `json:"id,omitempty"`
	Name             string  `json:"name,omitempty"`
	Barcode          string  `json:"barcode,omitempty"`
	ImageURL         string  `json:"image_url,omitempty"`
	SKU              string  `json:"sku,omitempty"`
	CostPrice        float64 `json:"cost_price,omitempty"`
	SellingPrice     float64 `json:"selling_price,omitempty"`
	Stock            int     `json:"stock,omitempty"`
	LowStockTreshold int     `json:"low_stock_treshold,omitempty"`
	Category         string  `json:"category,omitempty"`
	Supplier         string  `json:"supplier,omitempty"`
	CreatedAt        int64   `json:"created_at,omitempty"`
	UpdatedAt        int64   `json:"updated_at,omitempty"`
}
