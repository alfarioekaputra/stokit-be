package entity

type Supplier struct {
	ID        string    `gorm:"column:id;primaryKey"`
	Name      string    `gorm:"column:name"`
	Address   string    `gorm:"column:address"`
	Phone     string    `gorm:"column:phone"`
	Email     string    `gorm:"column:email"`
	Products  []Product `gorm:"foreignKey:SupplierID"`
	CreatedAt int64     `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64     `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}
