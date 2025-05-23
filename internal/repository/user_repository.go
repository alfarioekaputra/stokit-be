package repository

import (
	"stokit/internal/entity"
	"stokit/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) *UserRepository {
	return &UserRepository{
		Log: log,
	}
}

func (r *UserRepository) FindByEmail(db *gorm.DB, user *entity.User, email string) error {
	return db.Where("email = ?", email).First(user).Error
}

func (r *UserRepository) FindByToken(db *gorm.DB, user *entity.User, token string) error {
	return db.Where("token = ?", token).First(user).Error
}

func (r *UserRepository) CountByEmail(db *gorm.DB, email string) (int64, error) {
	var total int64

	err := db.Model(new(entity.User)).Where("email = ?", email).Count(&total).Error

	return total, err
}

func ApplyUserFilter(db *gorm.DB, filter *model.UserFilter) *gorm.DB {
	if filter.Email != nil {
		db = db.Where("email = ?", *filter.Email)
	}
	if filter.Username != nil {
		db = db.Where("username LIKE ?", "%"+*filter.Username+"%")
	}
	return db
}
