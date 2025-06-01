package entity

func (u *Product) TableName() string {
	return "products"
}

type Product struct {
	ID               string `gorm:"primaryKey"`
	CategoryID       string
	SupplierID       string
	Name             string
	Barcode          string
	ImageURL         string
	SKU              string
	CostPrice        float64
	SellingPrice     float64
	Stock            int
	LowStockTreshold int
	Category         Category `gorm:"foreignKey:CategoryID"`
	Supplier         Supplier `gorm:"foreignKey:SupplierID"`
	CreatedAt        int64    `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt        int64    `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}
