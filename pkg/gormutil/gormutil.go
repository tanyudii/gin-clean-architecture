package gormutil

import (
	"fmt"
	"github.com/vodeacloud/hr-api/pkg/common"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Scope func(db *gorm.DB) *gorm.DB

func Paginate(limit int32, page int32) Scope {
	var offset int32
	if page > 0 {
		offset = limit * (page - 1)
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(int(offset)).Limit(int(limit))
	}
}

func SearchLikeRight(qb *gorm.DB, search string, columns []string) *gorm.DB {
	if search == "" || len(columns) == 0 {
		return qb
	}
	for i := range columns {
		column := columns[i]
		qb = qb.Or(column+" LIKE ?", search+"%")
	}
	return qb
}

type MapSortableColumn map[int32]string

func Sort(sort int32, mapSortable MapSortableColumn) Scope {
	return func(db *gorm.DB) *gorm.DB {
		if sort == 0 {
			return db
		}
		column, ok := mapSortable[common.AbsInt32(sort)]
		if !ok {
			return db
		}
		sortType := "ASC"
		if sort < 0 {
			sortType = "DESC"
		}
		return db.Order(fmt.Sprintf("%s %s", column, sortType))
	}
}

func Count(db *gorm.DB, model schema.Tabler) (total int64, err error) {
	err = db.Model(&model).Count(&total).Error
	return
}
