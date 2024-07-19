package gocrud_gorm

import (
	"fmt"
	"github.com/spf13/cast"
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

// -------------- operator

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
	return db.(*gorm.DB).Offset(offset).Limit(pageSize)
}

// -------------- sort

func ASC(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	return db.(*gorm.DB).Order(fmt.Sprintf("%s ASC", field))
}

func DESC(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	return db.(*gorm.DB).Order(fmt.Sprintf("%s DESC", field))
}
