package entity

type Category struct {
	ID       string `gorm:"column:id;primaryKey"`
	Name     string `gorm:"column:name"`
	ParentID *string
	Parent   *Category   `gorm:"foreignKey:ParentID;PRELOAD:true" json:"parent,omitempty"`
	Children []*Category `gorm:"-"`
}
