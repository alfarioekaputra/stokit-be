package entity

type Category struct {
	ID       string `gorm:"column:id;primaryKey"`
	Name     string `gorm:"column:name"`
	ParentID *string
	Children []*Category `gorm:"-"`
}
