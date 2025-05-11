package entity

type User struct {
	ID        string `gorm:"column:id;primaryKey"`
	Email     string `gorm:"column:email;index:idx_email,unique"`
	Username  string `gorm:"column:username"`
	Password  string `gorm:"column:password"`
	Token     string `gorm:"column:token"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (u *User) TableName() string {
	return "users"
}
