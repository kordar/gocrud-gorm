package gocrud_gorm

import (
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

func SETVAL(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	return db.(*gorm.DB).Update(field, value).Error
}

func UPDATE(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	return db.(*gorm.DB).Update(field, value).Error
}

func UPDATES(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	if len(value2) == 0 {
		return db.(*gorm.DB).Updates(value).Error
	}
	return db.(*gorm.DB).Model(value).Updates(value2[0]).Error
}

func SAVE(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	return db.(*gorm.DB).Save(value).Error
}

func CREATE(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	return db.(*gorm.DB).Create(value).Error
}

func DELETE(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	return db.(*gorm.DB).Delete(value).Error
}

func PAGE(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	offset := cast.ToInt(value)
	pageSize := cast.ToInt(value2[0])
	return db.(*gorm.DB).Offset(offset).Limit(pageSize).Error
}
