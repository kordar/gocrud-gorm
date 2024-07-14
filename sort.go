package gocrud_gorm

import (
	"fmt"
	"gorm.io/gorm"
)

func ASC(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	return db.(*gorm.DB).Order(fmt.Sprintf("%s ASC", field))
}

func DESC(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	return db.(*gorm.DB).Order(fmt.Sprintf("%s DESC", field))
}
