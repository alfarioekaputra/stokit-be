package entity

func (u *Product) TableName() string {
	return "products"
}

type Product struct {
	ID           string `gorm:"primaryKey"`
	Name         string
	Barcode      string
	Brand        string
	CategoryID   string
	ImageURL     string
	SKU          string
	CostPrice    float64
	SellingPrice float64
	Stock        int
	Category     Category `gorm:"foreignKey:CategoryID"`
	CreatedAt    int64    `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt    int64    `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}
