package entity

type Category struct {
	ID       string `gorm:"primaryKey"`
	Name     string
	ParentID string
	Parent   *Category
}
