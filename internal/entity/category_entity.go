package entity

type Category struct {
	ID       string `gorm:"column:id;primaryKey"`
	Name     string
	ParentID string
	Parent   *Category
}
