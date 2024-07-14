package gocrud_gorm

import (
	"fmt"
	"gorm.io/gorm"
)

func EQ(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	return db.(*gorm.DB).Where(fmt.Sprintf("%s = ?", field), value)
}

func NEQ(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	return db.(*gorm.DB).Where(fmt.Sprintf("%s <> ?", field), value)
}

func LT(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	return db.(*gorm.DB).Where(fmt.Sprintf("%s < ?", field), value)
}

func LE(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	return db.(*gorm.DB).Where(fmt.Sprintf("%s <= ?", field), value)
}

func GT(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	return db.(*gorm.DB).Where(fmt.Sprintf("%s > ?", field), value)
}

func GE(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	return db.(*gorm.DB).Where(fmt.Sprintf("%s >= ?", field), value)
}

func IN(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	return db.(*gorm.DB).Where(fmt.Sprintf("%s IN ?", field), value)
}

func NOTIN(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	return db.(*gorm.DB).Where(fmt.Sprintf("%s NOT IN ?", field), value)
}

func LIKE(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	s := value.(string)
	return db.(*gorm.DB).Where(fmt.Sprintf("%s LIKE ?", field), "%"+s+"%")
}

func NOTLIKE(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	s := value.(string)
	return db.(*gorm.DB).Where(fmt.Sprintf("%s NOT LIKE %%?%%", field), "%"+s+"%")
}

func LIKELEFT(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	s := value.(string)
	return db.(*gorm.DB).Where(fmt.Sprintf("%s NOT LIKE ?%%", field), s+"%")
}

func LIKERIGHT(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	s := value.(string)
	return db.(*gorm.DB).Where(fmt.Sprintf("%s NOT LIKE %%?", field), "%"+s)
}

func BETWEEN(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	return db.(*gorm.DB).Where(fmt.Sprintf("%s BETWEEN ? AND ?", field), value, value2[0])
}

func NOTBETWEEN(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	return db.(*gorm.DB).Where(fmt.Sprintf("%s NOT BETWEEN ? AND ?", field), value, value2[0])
}

func ISNULL(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	return db.(*gorm.DB).Where(fmt.Sprintf("%s IS NULL", field))
}

func ISNOTNULL(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	return db.(*gorm.DB).Where(fmt.Sprintf("%s IS NOT NULL", field))
}
