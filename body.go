package gocrud_gorm

import (
	"github.com/kordar/gocrud"
	"gorm.io/gorm"
)

type GormSearchBody struct {
	*gocrud.SearchBody
}

func NewGormSearchBody(body gocrud.SearchBody) GormSearchBody {
	return GormSearchBody{SearchBody: &body}
}

func (search *GormSearchBody) GormQuery(db interface{}, parallel map[string]string) *gorm.DB {
	return search.Query(db, parallel).(*gorm.DB)
}

func (search *GormSearchBody) GormPaginate(db interface{}, parallel map[string]string) (*gorm.DB, error) {
	tx, err := search.Paginate(db, parallel)
	if tx == nil {
		return nil, err
	} else {
		return tx.(*gorm.DB), err
	}
}

func (search *GormSearchBody) GormQueryCustom(f func(search *gocrud.SearchBody) interface{}) *gorm.DB {
	return search.QueryCustom(f).(*gorm.DB)
}

// ===================================================

type GormEditorBody struct {
	*gocrud.EditorBody
}

func NewGormEditorBody(body gocrud.EditorBody) GormEditorBody {
	return GormEditorBody{EditorBody: &body}
}

// GormQuery 条件查询
func (form *GormEditorBody) GormQuery(db *gorm.DB, parallel map[string]string) *gorm.DB {
	return form.Query(db, parallel).(*gorm.DB)
}

// GormQuerySafe 防止空条件更新
func (form *GormEditorBody) GormQuerySafe(db *gorm.DB, parallel map[string]string) (*gorm.DB, error) {
	tx, err := form.QuerySafe(db, parallel)
	if tx == nil {
		return nil, err
	} else {
		return tx.(*gorm.DB), err
	}
}

// GormQueryCustom 自定义条件查询
func (form *GormEditorBody) GormQueryCustom(f func(body *gocrud.EditorBody) interface{}) *gorm.DB {
	return form.QueryCustom(f).(*gorm.DB)
}

// =========================================

type GormRemoveBody struct {
	*gocrud.RemoveBody
}

func NewGormRemoveBody(body gocrud.RemoveBody) GormRemoveBody {
	return GormRemoveBody{RemoveBody: &body}
}

func (remove *GormRemoveBody) GormQuerySafe(db *gorm.DB, parallel map[string]string) (*gorm.DB, error) {
	tx, err := remove.QuerySafe(db, parallel)
	if tx == nil {
		return nil, err
	} else {
		return tx.(*gorm.DB), err
	}
}

func (remove *GormRemoveBody) GormQuery(db *gorm.DB, parallel map[string]string) *gorm.DB {
	return remove.Query(db, parallel).(*gorm.DB)
}

func (remove *GormRemoveBody) GormQueryCustom(f func(body *gocrud.RemoveBody) interface{}) *gorm.DB {
	return remove.QueryCustom(f).(*gorm.DB)
}

// ===================================

type GormFormBody struct {
	*gocrud.FormBody
}

func NewGormFormBody(body gocrud.FormBody) GormFormBody {
	return GormFormBody{FormBody: &body}
}

func (form *GormFormBody) GormQuery(db *gorm.DB, parallel map[string]string) *gorm.DB {
	return form.Query(db, parallel).(*gorm.DB)
}

func (form *GormFormBody) GormQuerySafe(db *gorm.DB, parallel map[string]string) (*gorm.DB, error) {
	tx, err := form.QuerySafe(db, parallel)
	if tx == nil {
		return nil, err
	} else {
		return tx.(*gorm.DB), err
	}
}

func (form *GormFormBody) GormQueryCustom(f func(body *gocrud.FormBody) interface{}) *gorm.DB {
	return form.QueryCustom(f).(*gorm.DB)
}
