package gocrud_gorm

import (
	"fmt"

	"github.com/spf13/cast"
	"gorm.io/gorm"
)

func asGormDB(db interface{}) (*gorm.DB, bool) {
	if db == nil {
		return nil, false
	}
	tx, ok := db.(*gorm.DB)
	return tx, ok
}

func EQ(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	tx, ok := asGormDB(db)
	if !ok {
		return db
	}
	return tx.Where(fmt.Sprintf("%s = ?", field), value)
}

func NEQ(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	tx, ok := asGormDB(db)
	if !ok {
		return db
	}
	return tx.Where(fmt.Sprintf("%s <> ?", field), value)
}

func LT(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	tx, ok := asGormDB(db)
	if !ok {
		return db
	}
	return tx.Where(fmt.Sprintf("%s < ?", field), value)
}

func LE(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	tx, ok := asGormDB(db)
	if !ok {
		return db
	}
	return tx.Where(fmt.Sprintf("%s <= ?", field), value)
}

func GT(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	tx, ok := asGormDB(db)
	if !ok {
		return db
	}
	return tx.Where(fmt.Sprintf("%s > ?", field), value)
}

func GE(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	tx, ok := asGormDB(db)
	if !ok {
		return db
	}
	return tx.Where(fmt.Sprintf("%s >= ?", field), value)
}

func IN(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	tx, ok := asGormDB(db)
	if !ok {
		return db
	}
	return tx.Where(fmt.Sprintf("%s IN ?", field), value)
}

func NOTIN(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	tx, ok := asGormDB(db)
	if !ok {
		return db
	}
	return tx.Where(fmt.Sprintf("%s NOT IN ?", field), value)
}

func LIKE(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	tx, ok := asGormDB(db)
	if !ok {
		return db
	}
	s := cast.ToString(value)
	return tx.Where(fmt.Sprintf("%s LIKE ?", field), "%"+s+"%")
}

func NOTLIKE(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	tx, ok := asGormDB(db)
	if !ok {
		return db
	}
	s := cast.ToString(value)
	return tx.Where(fmt.Sprintf("%s NOT LIKE ?", field), "%"+s+"%")
}

func LIKELEFT(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	tx, ok := asGormDB(db)
	if !ok {
		return db
	}
	s := cast.ToString(value)
	return tx.Where(fmt.Sprintf("%s LIKE ?", field), s+"%")
}

func LIKERIGHT(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	tx, ok := asGormDB(db)
	if !ok {
		return db
	}
	s := cast.ToString(value)
	return tx.Where(fmt.Sprintf("%s LIKE ?", field), "%"+s)
}

func BETWEEN(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	tx, ok := asGormDB(db)
	if !ok {
		return db
	}
	return tx.Where(fmt.Sprintf("%s BETWEEN ? AND ?", field), value, value2[0])
}

func NOTBETWEEN(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	tx, ok := asGormDB(db)
	if !ok {
		return db
	}
	return tx.Where(fmt.Sprintf("%s NOT BETWEEN ? AND ?", field), value, value2[0])
}

func ISNULL(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	tx, ok := asGormDB(db)
	if !ok {
		return db
	}
	return tx.Where(fmt.Sprintf("%s IS NULL", field))
}

func ISNOTNULL(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	tx, ok := asGormDB(db)
	if !ok {
		return db
	}
	return tx.Where(fmt.Sprintf("%s IS NOT NULL", field))
}

// -------------- operator

func SETVAL(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	tx, ok := asGormDB(db)
	if !ok {
		return gorm.ErrInvalidDB
	}
	return tx.Update(field, value).Error
}

func UPDATE(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	tx, ok := asGormDB(db)
	if !ok {
		return gorm.ErrInvalidDB
	}
	return tx.Update(field, value).Error
}

func UPDATES(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	tx, ok := asGormDB(db)
	if !ok {
		return gorm.ErrInvalidDB
	}
	if len(value2) == 0 {
		return tx.Updates(value).Error
	}
	return tx.Model(value).Updates(value2[0]).Error
}

func SAVE(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	tx, ok := asGormDB(db)
	if !ok {
		return gorm.ErrInvalidDB
	}
	return tx.Save(value).Error
}

func CREATE(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	tx, ok := asGormDB(db)
	if !ok {
		return gorm.ErrInvalidDB
	}
	return tx.Create(value).Error
}

func DELETE(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	tx, ok := asGormDB(db)
	if !ok {
		return gorm.ErrInvalidDB
	}
	return tx.Delete(value).Error
}

func PAGE(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	tx, ok := asGormDB(db)
	if !ok {
		return db
	}
	offset := cast.ToInt(value)
	pageSize := cast.ToInt(value2[0])
	return tx.Offset(offset).Limit(pageSize)
}

// -------------- sort

func ASC(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	tx, ok := asGormDB(db)
	if !ok {
		return db
	}
	return tx.Order(fmt.Sprintf("%s ASC", field))
}

func DESC(db interface{}, field string, value interface{}, value2 ...interface{}) interface{} {
	tx, ok := asGormDB(db)
	if !ok {
		return db
	}
	return tx.Order(fmt.Sprintf("%s DESC", field))
}
