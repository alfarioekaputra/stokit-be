package helper

import "gorm.io/gorm"

func Preloads(preloadFields ...string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, field := range preloadFields {
			db = db.Preload(field)
		}
		return db
	}
}
